package cmd

import (
	"fmt"
	"time"

	"github.com/katsuya94/memo/util"
	"github.com/spf13/cobra"
)

func NewCommandOpen() *cobra.Command {
	c := &cobra.Command{
		Use:   "open [date]",
		Short: "Opens the memo for a given day",
		Long: `Opens the memo for a given day.
`,
		Args: cobra.ArbitraryArgs,
		RunE: CmdOpen,
	}

	return c
}

func CmdOpen(cmd *cobra.Command, args []string) error {
	var (
		d   util.Date
		err error
	)

	if len(args) == 0 {
		d = util.Today()
	} else {
		d, err = parseDate(args[0])
	}

	if err != nil {
		return err
	}

	memo, err := Profile.Get(d)
	if err != nil {
		return err
	}

	contents, err := Editor.Edit(memo.Contents())
	if err != nil {
		return err
	}

	err = memo.SetContents(contents)
	if err != nil {
		return err
	}

	return Profile.Put(d, memo)
}

var dateFormats = []string{
	"2006-1-2",
	"1/2/2006",
	"1.2.2006",
	"Jan 2 2006",
	"January 2 2006",
	"06-1-2",
	"1/2/06",
	"1.2.06",
	"Jan 2 06",
	"January 2 06",
	"1-2",
	"1/2",
	"1.2",
	"Jan 2",
	"January 2",
}

func parseDate(s string) (util.Date, error) {
	for _, format := range dateFormats {
		t, err := time.Parse(format, s)
		if err == nil {
			return util.NewDateFromTime(t), nil
		}
	}
	return util.Date{}, fmt.Errorf("invalid date format: %v", s)
}
