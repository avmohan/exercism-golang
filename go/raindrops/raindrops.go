// Package raindrops implements the problem
// as mentioned in the accompanying README.
package raindrops

import (
	"strconv"
	"strings"
)

const testVersion = 2

// translation table for each factor
var rainSpeak = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// Convert converts a given a number into raindrops speak,
// as mentioned in the accompanying README.
func Convert(x int) string {
	values := make([]string, 0)
	for _, v := range []int{3, 5, 7} {
		if x%v == 0 {
			values = append(values, rainSpeak[v])
		}
	}
	if len(values) == 0 {
		return strconv.Itoa(x)
	}
	return strings.Join(values, "")
}
