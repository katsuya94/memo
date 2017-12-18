package main

import (
	"fmt"
	"strings"
	"time"
)

type openCmd struct{}

var dateLayouts = []string{
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

func (cmd *openCmd) run(args ...string) error {
	var (
		d date
		t = time.Now()
	)

	if len(args) == 0 {
		d = newDate(t.Year(), t.Month(), t.Day())
	} else {
		var (
			dateString = strings.Join(args, " ")
			given      time.Time
			err        error
		)

		for _, layout := range dateLayouts {
			given, err = time.Parse(layout, dateString)
			if err == nil {
				break
			}
		}

		if err != nil {
			return fmt.Errorf("Invalid date format: %v", dateString)
		}

		if given.Year() == 0 {
			d = newDate(t.Year(), given.Month(), given.Day())
		} else {
			d = newDate(given.Year(), given.Month(), given.Day())
		}
	}

	return open(d)
}

func (*openCmd) usage() string {
	return `Usage: memo open [date]

   Opens the entry for a given date, but opens read-only if the day is in the past. The default day is today.
`
}

func (*openCmd) description() string {
	return "open (default)   Opens the entry for a given date"
}
