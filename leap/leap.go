// Package leap implements check for leap year
package leap

const testVersion = 3

// IsLeapYear checks that a given Gregorian calendar
// year is a leap year.
func IsLeapYear(x int) bool {
	return x%400 == 0 || x%4 == 0 && x%100 != 0
}
