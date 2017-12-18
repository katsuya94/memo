package main

import (
	"fmt"
	"strings"
	"time"
)

type OpenCmd struct{}

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

func (cmd *OpenCmd) Run(args ...string) error {
	var (
		date time.Time
		t    time.Time = time.Now()
	)

	if len(args) == 0 {
		date = normalizedDate(t.Year(), t.Month(), t.Day())
	} else {
		var (
			err        error
			dateString string = strings.Join(args, " ")
		)

		for _, layout := range dateLayouts {
			date, err = time.Parse(layout, dateString)
			if err == nil {
				break
			}
		}

		if err != nil {
			return fmt.Errorf("Invalid date format: %v", dateString)
		}

		if date.Year() == 0 {
			date = normalizedDate(t.Year(), date.Month(), date.Day())
		}
	}

	fmt.Printf("%v\n", date)

	return nil
}

func normalizedDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, &time.Location{})
}

func (*OpenCmd) Usage() string {
	return `Usage: memo open [date]

   Opens the entry for a given date, but opens read-only if the day is in the past. The default day is today.
`
}

func (*OpenCmd) Description() string {
	return "open (default)   Opens the entry for a given date"
}
