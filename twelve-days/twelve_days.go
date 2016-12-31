// Package twelve sing a Christmas song
package twelve

import (
	"fmt"
	"strings"
)

const testVersion = 1

var days = [...]string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var gifts = [...]string{
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

// Song returns a verse for given day
func Verse(day int) string {
	day = day - 1
	verse := fmt.Sprintf("On the %v day of Christmas my true love gave to me", days[day])
	for i := day; i >= 0; i-- {
		if i == 0 && day != 0 {
			verse += ", and " + gifts[i]
		} else {
			verse += ", " + gifts[i]
		}
	}
	return verse + "."
}

// Song returns a Christmas song
func Song() string {
	lines := []string{}
	for day := range days {
		lines = append(lines, Verse(day+1))
	}
	return strings.Join(lines, "\n") + "\n"
}
