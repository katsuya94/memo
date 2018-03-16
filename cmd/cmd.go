package cmd

import (
	"github.com/katsuya94/memo/core"
	"github.com/katsuya94/memo/editor"
	"github.com/spf13/cobra"
)

var (
	Home        string
	Profile     core.Profile
	ProfileName string
	Editor      editor.Editor
)

func NewDefaultCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "memo [date]",
		Short: "memo is a flexible utility for taking structured, searchable notes",
		Long: `memo is a flexible utility for taking structured, searchable notes.

Find more information at https://github.com/katsuya94/memo.
`,
		Args:              cobra.ArbitraryArgs,
		RunE:              CmdOpen,
		PersistentPreRunE: Setup,
	}

	c.PersistentFlags().StringVarP(&ProfileName, "profile", "p", "", "name of the profile to use")
	c.AddCommand(NewCommandOpen())

	return c
}
