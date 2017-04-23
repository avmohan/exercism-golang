package pangram

import (
	"strings"
)

const testVersion = 1

// IsPangram checks whether the given string is a pangram
// in the alphabet of english letters (case insensitive).
func IsPangram(s string) bool {
	s = strings.ToUpper(s)
	seen := make(map[rune]bool)
	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			seen[ch] = true
		}
	}
	return len(seen) == 26
}
