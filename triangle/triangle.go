// Package triangle provides calculating triangle types from side lengths
package triangle

import "math"

const testVersion = 3

type Kind int

const (
	NaT Kind = iota // 0 not a triangle
	Deg				// 1 degenerate
	Equ				// 2 equilateral
	Iso				// 3 isosceles
	Sca				// 4 scalene
)

// KindFromSides returns triangle kind by side lengths
func KindFromSides(a, b, c float64) Kind {
	// All sides must be valid number (can be 0)
	for _, side := range []float64{a, b, c} {
		if math.IsNaN(side) || math.IsInf(side, 0) || (side < 0) {
			return NaT;
		}
	}

	// At least one side must have length and all sides must connect (inequality theorem)
	if !(a + b + c > 0 && a+b >= c && a+c >= b && b+c >= a) {
		return NaT;
	}

	switch {
	//case a * b * c == 0 || b+c == a || a+c == b || a+b == c:
	//	return Deg
	case a == b && b == c: // 3 equal sides: equilateral
		return Equ
	case a == b || a == c || b == c: // two equal side: isoceles
		return Iso
	default: // all side unequal: scalene
		return Sca
	}
}
