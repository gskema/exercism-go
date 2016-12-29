/*
Package clock provides Clock type
*/
package clock

import "fmt"

const testVersion = 4

type Clock struct {
	hours, minutes int
}

// Creates and returns a new Clock
func New(hours, minutes int) Clock {
	const minutesInDay = 24 * 60
	var visibleHours, visibleMinutes, totalMinutes int

	// Convert everything to total minutes
	totalMinutes = (hours * 60) + minutes

	// Make sure our minutes fit within the day range
	// This also takes care of 24:00 -> 00:00
	totalMinutes %= minutesInDay

	// Minutes may be negative, normalize them
	if totalMinutes < 0 {
		totalMinutes += minutesInDay
	}

	visibleMinutes = totalMinutes % 60
	visibleHours = (totalMinutes - visibleMinutes) / 60

	return Clock{visibleHours, visibleMinutes}
}

// Outputs Clock in String form
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}

// Creates a new clock by adding X minutes
func (c Clock) Add(minutes int) Clock {
	return New(c.hours, c.minutes+minutes)
}
