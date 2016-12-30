// Package pangram provides pangram utils
package pangram

const testVersion = 1

// IsPangram returns true when a sentence is a pangram. ASCII only
func IsPangram(sentence string) bool {
	var letters = make(map[rune]bool)

	// ascii ~ 0-127, a-z ~ 97-122, A-Z 65-90
	for _, r := range sentence {
		if r > 96 && r < 123 {
			letters[r] = true
		} else if r > 64 && r < 91 {
			letters[r+32] = true
		}
	}

	return len(letters) == 26
}
