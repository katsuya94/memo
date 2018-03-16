package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCommandOpen() *cobra.Command {
	c := &cobra.Command{
		Use:   "open [date]",
		Short: "Opens the memo for a given day",
		Long: `Opens the memo for a given day.
`,
		Args: cobra.MaximumNArgs(1),
		Run:  CmdOpen,
	}

	return c
}

func CmdOpen(cmd *cobra.Command, args []string) {
	fmt.Println("open")
}
