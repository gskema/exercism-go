// Package raindrop provides rainspeak utils
package raindrops

import (
	"strconv"
	"strings"
)

const testVersion = 2

// Convert converts number to rainspeak
func Convert(number int) string {
	words := []string{}

	if number%3 == 0 {
		words = append(words, "Pling")
	}

	if number%5 == 0 {
		words = append(words, "Plang")
	}

	if number%7 == 0 {
		words = append(words, "Plong")
	}

	if len(words) > 0 {
		return strings.Join(words, "")
	}

	return strconv.Itoa(number)
}
