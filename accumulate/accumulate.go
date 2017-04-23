package accumulate

const testVersion = 1

// Accumulate applies function f on each element of inlist and
// returns the resulting slice as outlist
func Accumulate(inlist []string, f func(string) string) []string {
	outlist := make([]string, len(inlist))
	for i := range inlist {
		outlist[i] = f(inlist[i])
	}
	return outlist
}
