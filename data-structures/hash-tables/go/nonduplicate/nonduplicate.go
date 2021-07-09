package nonduplicate

// Nonduplicate takes a string and finds the first non-duplicate character.
// Returns a rune of non-duplicate character that appears first in the input
// string. If there are no duplicates it returns an empty rune.
func Nonduplicate(s string) rune {
	counts := make(map[rune]int)
	for _, c := range s {
		counts[c]++
	}

	for _, c := range s {
		if counts[c] == 1 {
			return c
		}
	}

	return ' '
}
