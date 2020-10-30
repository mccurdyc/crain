package {{{.Name}}}

import (
	"context"
	"flag"
	"log"

	"github.com/peterbourgon/ff/v3/ffcli"
)

func NewCommand(logger *log.Logger) *ffcli.Command {
	var (
		rootFlagSet = flag.NewFlagSet("{{{.Name}}}", flag.ExitOnError)
	)

	return &ffcli.Command{
		Name:       "{{{.Name}}}",
		ShortUsage: "{{{.Name}}} <subcommand>",
		FlagSet:    rootFlagSet,
		Exec: func(context.Context, []string) error {
			return flag.ErrHelp
		},
	}
}
