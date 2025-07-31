package booking

import (
	"fmt"
	"time"
)

var FORMATS = []string{
	"1/2/2006 15:04:05",
	"January 2, 2006 15:04:05",
	"Monday, January 2, 2006 15:04:05",
}

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	for _, dateFormat := range FORMATS {
		appointment, err := time.Parse(dateFormat, date)
		if err == nil {
			return appointment
		}
	}
	panic("Invalid date format")
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	appointment := Schedule(date)
	return appointment.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	appointment := Schedule(date)
	return appointment.Hour() >= 12 && appointment.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	appointment := Schedule(date)
	return fmt.Sprintf("You have an appointment on %s.", appointment.Format("Monday, January 2, 2006, at 15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	today := time.Now()
	anniversary := Schedule(fmt.Sprintf("%d/%d/%d 00:00:00", 9, 15, today.Year()))
	return anniversary
}
