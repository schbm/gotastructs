package pm

// Vergleicht Pattern p mit dem Text t für jede mögliche Position von p relativ zu t,
// bis eine Übereinstimmung oder alle mögliche Platzierungen ausprobiert wurden.
// Returns starting index of a substring of t for which p matches or -1
// O(nm)
func BruteForceMatch(t, p string) int {
	var n int = len(t)
	var m int = len(p)
	var diff int = n - m
	for i := 0; i <= diff; i++ { // Modified the condition to include the last character
		// test i-te Verschiebung des Patterns
		j := 0
		for j < m && t[i+j] == p[j] {
			j++
			if j == m {
				return i
			}
		}
	}
	return -1
}
