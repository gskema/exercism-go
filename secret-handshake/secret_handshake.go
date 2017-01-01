// Package secret converts integers to secret handshakes
package secret

import "fmt"

const testVersion = 1

var actions = [...]string{
	"wink",
	"double blink",
	"close your eyes",
	"jump",
}

func Handshake(code uint) []string {
	if code < 1 {
		return nil
	}

	bin := fmt.Sprintf("%032b", code)
	seq := []string{}

	for i := 31; i >= 28; i-- {
		if '0' == bin[i] {
			continue
		}
		if '1' == bin[27] {
			seq = append([]string{actions[31-i]}, seq...)
		} else {
			seq = append(seq, actions[31-i])
		}
	}

	return seq
}
