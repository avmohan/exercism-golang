package hamming

import "errors"

const testVersion = 5

// Distance returns the hamming distance between 2 strings
// of equal lengths. If the input strings have unequal length,
// an error is returned.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Unequal lengths")
	}
	diff := 0
	for i := range a {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff, nil
}
