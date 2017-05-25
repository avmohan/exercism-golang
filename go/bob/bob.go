package bob

import (
	"strings"
)

const testVersion = 2

func Hey(s string) string {
	s = strings.TrimSpace(s)
	switch {
	case s == "":
		return "Fine. Be that way!"
	case strings.ToUpper(s) == s && strings.ContainsAny(s, "ABCDEFGHIJKLMNOPQRSTUVWXYZ"):
		return "Whoa, chill out!"
	case string(s[len(s)-1]) == "?":
		return "Sure."
	default:
		return "Whatever."
	}
}
