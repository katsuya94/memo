package main

import (
	"time"
)

type date time.Time

func newDate(year int, month time.Month, day int) date {
	return date(time.Date(year, month, day, 0, 0, 0, 0, &time.Location{}))
}

func (d date) year() int {
	return time.Time(d).Year()
}

func (d date) month() time.Month {
	return time.Time(d).Month()
}

func (d date) day() int {
	return time.Time(d).Day()
}

func (d date) String() string {
	return time.Time(d).Format("2006-01-02")
}

type command interface {
	run(...string) error
	usage() string
	description() string
}

type usageError struct {
	cmd command
}

func (err usageError) Error() string {
	return err.cmd.usage()
}
