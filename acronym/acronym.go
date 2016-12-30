// Package acronym provides acronym utils
package acronym

import (
	"strings"
	"unicode"
)

const testVersion = 2

// Abbreviate creates an acronym from a phrase
func Abbreviate(phrase string) string {
	var dashBefore, spaceBefore, capitalBefore bool
	var acronym string

	for _, r := range phrase {
		c := string(r)

		// Break loop if capital letter was followed by ":"
		if capitalBefore && r == ':' {
			break
		}

		if dashBefore || spaceBefore || unicode.IsUpper(r) {
			acronym += c
		}

		capitalBefore = unicode.IsUpper(r)
		spaceBefore = unicode.IsSpace(r)
		dashBefore = r == '-'
	}

	return strings.ToUpper(acronym)
}
