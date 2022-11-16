package meetup

import (
	"time"
)

// Define the WeekSchedule type here.
type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Fifth
	Last
	Teenth
)

func nthDay(wShed WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	d := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	count := 0
	for {
		if d.Weekday() != wDay {
			d = d.AddDate(0, 0, 1)
			continue
		}
		if count != int(wShed) {
			count++
			d = d.AddDate(0, 0, 1)
			continue
		}
		return d.Day()
	}
}
func lastDay(wDay time.Weekday, month time.Month, year int) int {
	d := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	var l time.Time
	for {
		if d.Month() != month {
			break
		}
		if d.Weekday() == wDay {
			l = d
		}
		d = d.AddDate(0, 0, 1)
	}
	return l.Day()
}
func teenthDay(wDay time.Weekday, month time.Month, year int) int {
	d := time.Date(year, month, 13, 0, 0, 0, 0, time.UTC)
	for d.Day() <= 19 {
		if d.Weekday() == wDay {
			return d.Day()
		}
		d = d.AddDate(0, 0, 1)
	}
	return -1
}

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	switch {
	case wSched == Last:
		return lastDay(wDay, month, year)
	case wSched == Teenth:
		return teenthDay(wDay, month, year)
	default:
		return nthDay(wSched, wDay, month, year)
	}
}
