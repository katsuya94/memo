package util

import (
	"testing"
	"time"

	assert "github.com/stretchr/testify/assert"
)

func TestToday(t *testing.T) {
	ti := time.Now()
	d := Today()
	assert.Equal(t, ti.Year(), d.Year())
	assert.Equal(t, int(ti.Month()), d.Month())
	assert.Equal(t, ti.Day(), d.Day())
}

func TestNewDate(t *testing.T) {
	d := NewDate(1993, 8, 5)
	assert.Equal(t, 1993, d.Year())
	assert.Equal(t, 8, d.Month())
	assert.Equal(t, 5, d.Day())
}

func TestNewDateFromTime(t *testing.T) {
	ti := time.Now()
	d := NewDateFromTime(ti)
	assert.Equal(t, ti.Year(), d.Year())
	assert.Equal(t, int(ti.Month()), d.Month())
	assert.Equal(t, ti.Day(), d.Day())
}

func TestNewDateFromString(t *testing.T) {
	d, err := NewDateFromString("1994-09-24")
	assert.Nil(t, err)
	assert.Equal(t, 1994, d.Year())
	assert.Equal(t, 9, d.Month())
	assert.Equal(t, 24, d.Day())
}

func TestNewDateFromString_InvalidFormat(t *testing.T) {
	_, err := NewDateFromString("1994-9-24")
	assert.EqualError(t, err, "parsing time \"1994-9-24\": month out of range")
}

func TestDate_Year(t *testing.T) {
	d := NewDate(1993, 8, 5)
	assert.Equal(t, 1993, d.Year())
}

func TestDate_Month(t *testing.T) {
	d := NewDate(1993, 8, 5)
	assert.Equal(t, 8, d.Month())
}

func TestDate_Day(t *testing.T) {
	d := NewDate(1993, 8, 5)
	assert.Equal(t, 5, d.Day())
}

func TestDate_Before_Before(t *testing.T) {
	d1 := NewDate(1993, 8, 5)
	d2 := NewDate(1994, 9, 24)
	assert.True(t, d1.Before(d2))
}

func TestDate_Before_After(t *testing.T) {
	d1 := NewDate(1993, 8, 5)
	d2 := NewDate(1994, 9, 24)
	assert.False(t, d2.Before(d1))
}

func TestDate_Before_Same(t *testing.T) {
	d := NewDate(1993, 8, 5)
	assert.False(t, d.Before(d))
}

func TestDate_String(t *testing.T) {
	d := NewDate(1993, 8, 5)
	assert.Equal(t, "1993-08-05", d.String())
}
