package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mccurdyc/crain/pkg/project"
	"github.com/peterbourgon/ff/v3/ffcli"
	"github.com/rs/zerolog"
)

func main() {
	l := zerolog.New(os.Stdout)
	logger := log.New(l, "crain - ", log.Ldate|log.Ltime|log.LUTC)

	os.Exit(Run(context.Background(), os.Args[1:], logger))
}

func Run(ctx context.Context, args []string, l *log.Logger) int {
	var (
		rootFlagSet = flag.NewFlagSet("crain", flag.ExitOnError)
	)

	root := &ffcli.Command{
		Name:       "crain",
		ShortUsage: "crain <subcommand> [flags]",
		FlagSet:    rootFlagSet,
		Exec: func(context.Context, []string) error {
			return flag.ErrHelp
		},
	}

	if len(args) >= 2 {
		root.Subcommands = []*ffcli.Command{
			project.NewCommand(args[2:], l),
		}
	}

	if len(args) < 2 {
		fmt.Println(ffcli.DefaultUsageFunc(root))
		return 1
	}

	err := root.ParseAndRun(context.Background(), args)
	if err != nil {
		l.Println(err)
		return 1
	}

	return 0
}
