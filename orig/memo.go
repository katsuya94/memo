package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	var (
		cmd  command
		args []string
	)

	if len(os.Args) < 2 {
		cmd = &openCmd{}
	} else {
		cmd = resolveCmd(os.Args[1])
		if cmd == nil {
			cmd = &openCmd{}
			args = os.Args[1:]
		} else {
			args = os.Args[2:]
		}
	}

	err := cmd.run(args...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func resolveCmd(name string) command {
	switch name {
	case "open":
		return &openCmd{}
	case "help":
		return &helpCmd{}
	}

	return nil
}
