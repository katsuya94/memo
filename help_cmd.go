package main

import (
	"fmt"
	"strings"
)

type HelpCmd struct{}

var commands = []Command{
	&OpenCmd{},
	&HelpCmd{},
}

func (cmd *HelpCmd) Run(args ...string) error {
	switch len(args) {
	case 0:
		subcommandDescriptions := make([]string, len(commands))

		for i, cmd := range commands {
			subcommandDescriptions[i] = cmd.Description()
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

		fmt.Print(cmd.Usage())
	default:
		return &UsageError{cmd}
	}

	return nil
}

func (*HelpCmd) Usage() string {
	return `Usage: memo help [subcommand]

   Shows help message, including all subcommands. Shows help message for the given subcommand, if specified.
`
}

func (*HelpCmd) Description() string {
	return "help             Shows this help message or help for subcommands"
}
