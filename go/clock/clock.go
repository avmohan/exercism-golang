// Package clock implements a clock, that tracks hours & minutes
// and rolls over after 24 hours.
package clock

import "fmt"

const testVersion = 4

// Clock stores number of minutes
type Clock int

const (
	minutesInADay = 60 * 24
)

// New returns a new Clock initialized to the set values.
func New(hour, minute int) Clock {
	return Clock(hour*60 + minute).normalize()
}

// String uses "hh:mm" format for the string representation
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hours(), c.Minutes())
}

// Add given number
func (c Clock) Add(minutes int) Clock {
	return (c + Clock(minutes)).normalize()
}

// Hours return number of hours
func (c Clock) Hours() Clock {
	return c / 60
}

// Minutes return number of minutes
func (c Clock) Minutes() Clock {
	return c % 60
}

func (c Clock) normalize() Clock {
	c %= minutesInADay
	if c < 0 {
		c += minutesInADay
	}
	return c
}
