package gigasecond

import "time"

const (
	testVersion = 4
	giga        = 1000000000
	gigaToNano  = giga * giga
)

// AddGigasecond returns the time obtained by
// adding 1 giga second to the given time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigaToNano)
}
