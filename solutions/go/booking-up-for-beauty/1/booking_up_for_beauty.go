package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)

    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
    t, _ := time.Parse(layout, date)
    now := time.Now()

    return t.Before(now)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
    t, _ := time.Parse(layout, date)
    isAfterNoon := false

    if t.Hour() >= 12 && t.Hour() < 18 {
        isAfterNoon = true
    }

    return isAfterNoon
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
    t, _ := time.Parse(layout, date)
    f := t.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")

    return f
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	d := time.Now()
	t := time.Date(d.Year(), 9, 15, 0, 0, 0, 0, d.Location())
    return t
}
