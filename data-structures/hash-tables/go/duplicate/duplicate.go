package duplicate

// Duplicate takes an array of strings and returns the first duplicate value
// it finds.
// Assumes that there is always at least one duplicate value in the input array.
// Returns the first duplicate string.
func Duplicate(s []string) string {
	table := make(map[string]bool)

	for _, str := range s {
		if table[str] {
			return str
		} else {
			table[str] = true
		}
	}

	return ""
}
