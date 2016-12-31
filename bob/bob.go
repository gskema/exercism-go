// Package bob provides Bob object which you can talk to
package bob

import (
	"regexp"
	"strings"
)

const testVersion = 2

func IsQuestion(talk string) bool {
	return '?' == talk[len(talk)-1] && !IsYelling(talk)
}

func IsYelling(talk string) bool {
	if strings.ToUpper(talk) != talk {
		return false
	}

	matched, _ := regexp.MatchString("[A-Z]", talk)
	return matched
}

func IsSilence(talk string) bool {
	return 0 == len(talk)
}

// Hey allows allows talking to Bob
func Hey(talk string) string {
	talk = strings.TrimSpace(talk)

	switch true {
	case IsSilence(talk):
		return "Fine. Be that way!"
	case IsQuestion(talk):
		return "Sure."
	case IsYelling(talk):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}
