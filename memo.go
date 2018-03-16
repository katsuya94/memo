package main

import (
	"os"

	"github.com/katsuya94/memo/cmd"
)

func main() {
	command := cmd.NewDefaultCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
