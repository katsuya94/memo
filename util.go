package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type date time.Time

var zeroDate = date{}

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

func parseDate(s string) (date, error) {
	var year, month, day int

	_, err := fmt.Sscanf(s, "%d-%d-%d", &year, &month, &day)
	if err != nil {
		return date{}, err
	}

	return newDate(year, time.Month(month), day), nil
}

type command interface {
	run(...string) error
	usage() string
	description() string
}

type errUsage struct {
	cmd command
}

func (err errUsage) Error() string {
	return err.cmd.usage()
}

type errNotFound struct{}

func (errNotFound) Error() string {
	return "memo not found"
}

type errMalformedSectionHeader struct {
	line string
}

func (err errMalformedSectionHeader) Error() string {
	return fmt.Sprintf("malformed section header: %v", err.line)
}

func confirm(recommended bool, message string) (bool, error) {
	if recommended {
		fmt.Printf("%v (Y/n) ", message)
	} else {
		fmt.Printf("%v (y/N) ", message)
	}

	raw, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		return false, err
	}

	input := strings.ToLower(strings.TrimSpace(string(raw)))

	switch input {
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
