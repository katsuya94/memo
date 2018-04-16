package cmd

import (
	"bytes"
	"fmt"
	"time"

	"github.com/katsuya94/memo/editor"
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

	if len(memo) == 0 {
		memo = append(memo, util.Section{})
	}

	e, err := editor.NewEditor()
	if err != nil {
		return err
	}

	if err = util.WriteMemo(memo, e); err != nil {
		return err
	}
	if err = e.Close(); err != nil {
		return err
	}

	for {
		if err = e.Launch(false); err != nil {
			return err
		}

		if memo, err = util.ReadMemo(e); err == nil {
			break
		}

		if _, ok := err.(util.MemoFormatError); !ok {
			return err
		}

		fmt.Println(err)

		if err = e.Close(); err != nil {
			return err
		}

		if ok, err := confirm(false, "Discard changes?"); err != nil {
			return err
		} else if ok {
			return nil
		}
	}

	if err = e.Close(); err != nil {
		return err
	}

	processedMemo := util.Memo{}
	for _, s := range memo {
		if !(s.Name == "" && len(s.Tags) == 0 && len(bytes.TrimSpace(s.Body)) == 0) {
			processedMemo = append(processedMemo, s)
		}
	}

	return Profile.Put(d, processedMemo)
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
		if err != nil {
			continue
		}

		if t.Year() == 0 {
			t = t.AddDate(time.Now().Year(), 0, 0)
		}

		return util.NewDateFromTime(t), nil
	}
	return util.Date{}, fmt.Errorf("invalid date format: %v", s)
}
