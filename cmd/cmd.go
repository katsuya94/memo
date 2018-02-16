package cmd

import (
	"github.com/katsuya94/memo/core"
	"github.com/spf13/cobra"
)

var (
	Home        string
	ProfileName string
	Profile     core.Profile
)

func NewDefaultCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "memo",
		Short: "memo is a flexible utility for taking structured, searchable notes",
		Long: `memo is a flexible utility for taking structured, searchable notes.

Find more information at https://github.com/katsuya94/memo.
`,
		Run:              CmdOpen,
		PersistentPreRun: Setup,
	}

	c.PersistentFlags().StringVarP(&ProfileName, "profile", "p", "default", "name of the profile to use")

	c.AddCommand(NewCommandOpen())

	return c
}
