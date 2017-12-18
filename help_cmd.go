package main

import (
	"fmt"
	"strings"
)

type helpCmd struct{}

var commands = []command{
	&openCmd{},
	&helpCmd{},
}

func (cmd *helpCmd) run(args ...string) error {
	switch len(args) {
	case 0:
		subcommandDescriptions := make([]string, len(commands))

		for i, cmd := range commands {
			subcommandDescriptions[i] = cmd.description()
		}

		subcommandsDescription := strings.Join(
			subcommandDescriptions,
			"\n   ",
		)

		fmt.Printf(`Usage: memo [subcommand] ...

Subcommands:
   %v
`, subcommandsDescription)
	case 1:
		cmd := resolveCmd(args[0])
		if cmd == nil {
			return fmt.Errorf("Unknown subcommand: %v", args[0])
		}

		fmt.Println(cmd.usage())
	default:
		return &usageError{cmd}
	}

	return nil
}

func (*helpCmd) usage() string {
	return `Usage: memo help [subcommand]

   Shows help message, including all subcommands. Shows help message for the given subcommand, if specified.`
}

func (*helpCmd) description() string {
	return "help             Shows this help message or help for subcommands"
}
