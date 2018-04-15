package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/katsuya94/memo/cmd"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	command := cmd.NewDefaultCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
