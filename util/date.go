package util

import (
	"fmt"
	"time"
)

type Date time.Time

func Today() Date {
	return NewDateFromTime(time.Now())
}

func NewDateFromTime(t time.Time) Date {
	return NewDate(t.Year(), int(t.Month()), t.Day())
}

func NewDate(year int, month int, day int) Date {
	return Date(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
}

func (d Date) Year() int {
	return time.Time(d).Year()
}

func (d Date) Month() int {
	return int(time.Time(d).Month())
}

func (d Date) Day() int {
	return time.Time(d).Day()
}

func (d Date) String() string {
	return fmt.Sprintf("%v-%v-%v", d.Year(), d.Month(), d.Day())
}
