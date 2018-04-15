package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/katsuya94/memo/core"
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
		memo = append(memo, core.Section{})
	}

	e, err := editor.NewEditor()
	if err != nil {
		return err
	}

	if err = core.WriteMemo(memo, e); err != nil {
		return err
	}
	if err = e.Close(); err != nil {
		return err
	}

	for {
		if err = e.Launch(false); err != nil {
			return err
		}

		if memo, err = core.ReadMemo(e); err == nil {
			break
		}

		if _, ok := err.(core.MemoFormatError); !ok {
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

	return Profile.Put(d, memo)
}

func confirm(recommended bool, message string) (bool, error) {
	if recommended {
		fmt.Printf("%v (Y/n) ", message)
	} else {
		fmt.Printf("%v (y/N) ", message)
	}

	b, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		return recommended, err
	}

	b = bytes.TrimSpace(b)
	b = bytes.ToLower(b)

	switch string(b) {
	case "":
		return recommended, nil
	case "n":
		return false, nil
	case "y":
		return true, nil
	default:
		return confirm(recommended, "Please type y or n.")
	}
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
