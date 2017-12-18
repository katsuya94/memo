package main

import (
	"fmt"
	"os"
)

type Command interface {
	Run(...string) error
	Usage() string
	Description() string
}

type UsageError struct {
	cmd Command
}

func (err *UsageError) Error() string {
	return err.cmd.Usage()
}

func main() {
	var (
		cmd  Command
		args []string
	)

	if len(os.Args) < 2 {
		cmd = &OpenCmd{}
	} else {
		cmd = resolveCmd(os.Args[1])
		if cmd == nil {
			cmd = &OpenCmd{}
			args = os.Args[1:]
		} else {
			args = os.Args[2:]
		}
	}

	err := cmd.Run(args...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func resolveCmd(name string) Command {
	switch name {
	case "open":
		return &OpenCmd{}
	case "help":
		return &HelpCmd{}
	}

	return nil
}
