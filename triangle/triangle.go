package triangle

import (
	"math"
	"sort"
)

const testVersion = 3

func KindFromSides(a, b, c float64) Kind {
	d := sort.Float64Slice{a, b, c}
	for _, v := range d {
		if v == 0 || math.IsNaN(v) || math.IsInf(v, 0) {
			return NaT
		}
	}
	sort.Sort(d)
	switch {
	case d[0]+d[1] < d[2]:
		return NaT
	case d[0] == d[1] && d[0] == d[2] && d[1] == d[2]:
		return Equ
	case d[0] == d[1] || d[0] == d[2] || d[1] == d[2]:
		return Iso
	default:
		return Sca
	}
}

type Kind string

const (
	NaT = "Not a triangle"
	Equ = "Equilateral" // equilateral
	Iso = "Isosceles"   // isosceles
	Sca = "Scalene"     // scalene

)
