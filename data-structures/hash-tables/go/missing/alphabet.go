package alphabet

func toChar(i int) rune {
	return rune('a' - 1 + i)
}

// Alphabet finds a letter from the alphabet that does not appear in the input
// string.
// Returns missing letter as `string`.
func Alphabet(s string) rune {
	presentChars := make(map[rune]bool)
	for _, c := range s {
		presentChars[c] = true
	}

	for i := 1; i <= 26; i++ {
		letter := toChar(i)
		if presentChars[letter] != true {
			return letter
		}
	}

	return ' '
}
