package booking

import (
	"fmt"
	"time"
)

var layout = "1/2/2006 15:04:05"
// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {

	
	parsedTime, _ := time.Parse(layout, date)


	// Check if the parsed date is valid
	if parsedTime.Year() < 1000 || parsedTime.Year() > 9999 {
		return time.Time{}
	}

	return parsedTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	
	layout = "January 2, 2006 15:04:05"
	t := Schedule(date)


	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout = "Monday, January 2, 2006 15:04:05"
	t := Schedule(date)

	return t.Hour()>=12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout = "1/2/2006 15:04:05"
	t := Schedule(date)
	ans := fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %s.",t.Weekday(),t.Month(), t.Day(), t.Year(), t.Format("15:04"))
	return ans
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	
	t := time.Date(time.Now().Year(),time.September,15,0,0,0,0,time.UTC)
	return t
}
