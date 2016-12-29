// Package gigasecond provides function to calculate time
// when someone has lived for a gigasecond
package gigasecond

import "time"

const testVersion = 4

// Calculates time when someone has lived for a gigasecond
// 10^9 = 1000000000
func AddGigasecond(birthday time.Time) time.Time {
	return birthday.Add(time.Duration(1000000000) * time.Second)
}
