package project

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"text/template"

	"github.com/BurntSushi/toml"
)

const (
	gitDir = ".git"
)

type execFn func(context.Context, []string) error

type generateConfig struct {
	templatesPath      string
	templates          []string
	templateConfigFile string
	dest               string
}

func generate(cfg generateConfig, logger *log.Logger) execFn {
	return func(ctx context.Context, _ []string) error {
		tmplVarDefs, err := readDynamicConfig(cfg.templateConfigFile)
		if err != nil {
			return err
		}

		for _, tmpl := range cfg.templates {
			src := filepath.Join(cfg.templatesPath, tmpl)

			if err := copyDir(cfg.dest, src); err != nil {
				return err
			}

			if err := renderDir(cfg.dest, tmplVarDefs); err != nil {
				return err
			}
		}

		return nil
	}
}

func renderDir(filename string, config map[string]interface{}) error {
	files, err := ioutil.ReadDir(filename)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.Name() == gitDir {
			continue
		}

		if file.IsDir() {
			if err := renderDir(filepath.Join(filename, file.Name()), config); err != nil {
				return err
			}
			continue
		}

		if err := renderFile(filepath.Join(filename, file.Name()), config); err != nil {
			return err
		}
	}

	return nil
}

func renderFile(filename string, config map[string]interface{}) error {
	t, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	tmpl := template.New(filename).Delims("{{{", "}}}")
	tmpl = template.Must(tmpl.Parse(string(t)))

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, config)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, buf.Bytes(), 0644)
}

func readDynamicConfig(f string) (map[string]interface{}, error) {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	dynamicCfg := make(map[string]interface{})

	if err := toml.Unmarshal(b, &dynamicCfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return dynamicCfg, nil
}

// copyDir copies a directory recursively from src to dst.
func copyDir(dst string, src string) error {
	srcStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcStat.Mode()); err != nil {
		return err
	}

	fs, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, f := range fs {
		srcfp := path.Join(src, f.Name())
		dstfp := path.Join(dst, f.Name())

		if f.IsDir() {
			if err := copyDir(dstfp, srcfp); err != nil {
				return err
			}
			continue
		}

		if err := copyFile(dstfp, srcfp); err != nil {
			return err
		}
	}

	return nil
}

// copyFile copies a file from src to dst.
func copyFile(dst string, src string) error {
	from, err := os.Open(src)
	if err != nil {
		return err
	}
	defer from.Close()

	to, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	return err
}
