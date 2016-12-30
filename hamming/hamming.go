// Package hamming provides utils to calculate Hamming distance
// between two strings
package hamming

import "errors"

const testVersion = 5

// Distance calculates Hamming distance between two strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("String length must be equal")
	}

	var distance int = 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
