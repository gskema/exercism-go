// Package house sings nursery rhymes
package house

import "strings"

const testVersion = 1

type SongLine struct {
	action  string
	subject string
}

var songLines = []SongLine{
	{"", "the horse and the hound and the horn"},
	{"belonged to", "the farmer sowing his corn"},
	{"kept", "the rooster that crowed in the morn"},
	{"woke", "the priest all shaven and shorn"},
	{"married", "the man all tattered and torn"},
	{"kissed", "the maiden all forlorn"},
	{"milked", "the cow with the crumpled horn"},
	{"tossed", "the dog"},
	{"worried", "the cat"},
	{"killed", "the rat"},
	{"ate", "the malt"},
	{"lay in", "the house that Jack built"},
}

// Verse return n-th verse
func Verse(num int) string {
	var verseLines []string
	l := len(songLines)
	i := l - num

	verseLines = append(verseLines, "This is "+songLines[i].subject)
	for i++; i < l; i++ {
		verseLines = append(verseLines, "that "+songLines[i].action+" "+songLines[i].subject)
	}

	return strings.Join(verseLines, "\n") + "."
}

// Song returns song
func Song() string {
	var verses []string
	for num := 1; num <= len(songLines); num++ {
		verses = append(verses, Verse(num))
	}
	return strings.Join(verses, "\n\n")
}
