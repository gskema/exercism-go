// Package triangle provides calculating triangle types from side lengths
package triangle

import "math"

const testVersion = 3

type Kind struct {
	name string
}

// Pick values for the following identifiers used by the test program.
var NaT = Kind{"NaT"} // not a triangle
var Equ = Kind{"Equ"} // equilateral
var Iso = Kind{"Iso"} // isosceles
var Sca = Kind{"Sca"} // scalene
var Deg = Kind{"Deg"} // degenerate

// IsLength checks if given float64 can be a valid triangle side length
func IsLength(length float64) bool {
	return length >= 0 && !math.IsNaN(length) && !math.IsInf(length, 0)
}

// IsTriangle checks if side lengths a,b,c can form a triangle
func IsTriangle(a, b, c float64) bool {
	hasValidSides := IsLength(a) && IsLength(b) && IsLength(c)
	hasAtLeastOneSide := a+b+c > 0
	sidesCanConnect := a+b >= c && a+c >= b && b+c >= a

	return hasValidSides && hasAtLeastOneSide && sidesCanConnect
}

// IsDegenerate calculates if a triangle looks like a line (all 3 points on the same line)
func IsDegenerate(a, b, c float64) bool {
	if !IsTriangle(a, b, c) {
		return false
	}

	// Ones side is of 0 length
	if a == 0 || b == 0 || c == 0 {
		return true
	}

	// Length of longest side is equal to the sum of the two other sides
	if b+c == a || a+c == b || a+b == c {
		return true
	}

	return false
}

// IsIsosceles calculates if a triangle is an equilateral triangle (with three equal sides)
func IsEquilateral(a, b, c float64) bool {
	if !IsTriangle(a, b, c) {
		return false
	}
	return a == b && b == c
}

// IsIsosceles calculates if a triangle is an Isoceles triangle (with two equal sides)
func IsIsosceles(a, b, c float64) bool {
	if !IsTriangle(a, b, c) {
		return false
	}
	return a == b || a == c || b == c
}

// IsScalene calculates if a triangle has unequal (all 3) sides
func IsScalene(a, b, c float64) bool {
	if !IsTriangle(a, b, c) {
		return false
	}
	return a != b && a != c && b != c
}

// Code this function.
func KindFromSides(a, b, c float64) Kind {
	switch true {
	case !IsTriangle(a, b, c):
		return NaT
	//case IsDegenerate(a, b, c):
	//	return Deg
	case IsEquilateral(a, b, c):
		return Equ
	case IsIsosceles(a, b, c):
		return Iso
	case IsScalene(a, b, c):
		return Sca
	}

	return NaT
}
