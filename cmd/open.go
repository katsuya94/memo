package cmd

import (
	"github.com/spf13/cobra"
)

func NewCommandOpen() *cobra.Command {
	c := &cobra.Command{
		Use:   "open",
		Short: "Opens the memo for a given day",
		Long: `Opens the memo for a given day.
`,
		Run: CmdOpen,
	}

	return c
}

func CmdOpen(cmd *cobra.Command, args []string) {
}
