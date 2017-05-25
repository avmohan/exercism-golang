package acronym

import (
	"strings"
	"unicode"
)

const testVersion = 2

// Abbreviate abbreviates a string as explained in the readme
func Abbreviate(instring string) string {
	outrunes := make([]string, 0)
	prev := '-'
	for _, v := range instring {
		if (unicode.IsUpper(v) && !unicode.IsUpper(prev) ||
			unicode.IsPunct(prev) ||
			unicode.IsSpace(prev)) &&
			!unicode.IsSpace(v) {
			outrunes = append(outrunes, string(unicode.ToUpper(v)))
		}
		prev = v
	}
	return strings.Join(outrunes, "")
}
