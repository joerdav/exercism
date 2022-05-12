package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	t, err := time.Parse("1/2/2006 15:04:05", date)
	if err == nil {
		return t
	}
	t, err = time.Parse("January 2, 2006 15:04:05", date)
	if err == nil {
		return t
	}
	t, err = time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err == nil {
		return t
	}
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	return Schedule(date).Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	d := Schedule(date).Hour()
	return d >= 12 && d < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	return Schedule(date).Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
