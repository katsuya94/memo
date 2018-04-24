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

	memo, ok, err := editMemo(memo)
	if err != nil {
		return err
	} else if !ok {
		return nil
	}

	processedMemo := util.Memo{}

	return Profile.Put(d, processedMemo)
}

func editMemo(memo util.Memo) (util.Memo, bool, error) {
	e, err := editor.NewEditor()
	if err != nil {
		return nil, false, err
	}

	if err = util.WriteMemo(memo, e); err != nil {
		return nil, false, err
	}
	if err = e.Close(); err != nil {
		return nil, false, err
	}

	for {
		if err = e.Launch(false); err != nil {
			return nil, false, err
		}

		if memo, err = util.ReadMemo(e); err == nil {
			break
		}

		if _, ok := err.(util.MemoFormatError); !ok {
			return nil, false, err
		}

		fmt.Println(err)

		if err = e.Close(); err != nil {
			return nil, false, err
		}

		if ok, err := confirm(false, "Discard changes?"); err != nil {
			return nil, false, err
		} else if ok {
			return nil, false, nil
		}
	}

	memo = filterEmptySections(memo)

	if err = e.Close(); err != nil {
		return nil, false, err
	}

	return memo, true, nil
}

func filterEmptySections(memo util.Memo) util.Memo {
	var filtered util.Memo
	for _, s := range memo {
		if !(s.Name == "" && len(s.Tags) == 0 && len(bytes.TrimSpace(s.Body)) == 0) {
			filtered = append(filtered, s)
		}
	}
	return filtered
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
