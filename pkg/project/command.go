package project

import (
	"context"
	"flag"
	"log"
	"strings"

	"github.com/peterbourgon/ff"
	"github.com/peterbourgon/ff/v3/ffcli"
)

func NewCommand(args []string, l *log.Logger) *ffcli.Command {
	var (
		fs        = flag.NewFlagSet("project", flag.ExitOnError)
		tmplsPath = fs.String("templates-path", "./_tmpls", "Path to templates.")
		tmpls     = fs.String("templates", "base", "A comma-separate list of template names.")
		dest      = fs.String("destination", ".", "The destination directory where the project should be created.")
		config    = fs.String("config", "config.toml", "A config file defining values for the required variables for all templates used.")
	)

	err := ff.Parse(fs, args)
	if err != nil {
		l.Printf("failed to parse project flags: %+v", err)
		return nil
	}

	return &ffcli.Command{
		Name:       "project",
		ShortUsage: "project <subcommand>",
		FlagSet:    fs,
		Subcommands: []*ffcli.Command{
			{
				Name:       "generate",
				ShortUsage: "generate [flags]",
				ShortHelp:  "Generates a new project.",
				FlagSet:    fs,
				Exec: generate(
					generateConfig{
						templatesPath:      *tmplsPath,
						templates:          strings.Split(*tmpls, ","),
						templateConfigFile: *config,
						dest:               *dest,
					},
					l),
			},
		},
		Exec: func(context.Context, []string) error {
			return flag.ErrHelp
		},
	}
}
