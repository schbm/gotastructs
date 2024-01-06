package pm

// Depends on two heuristics
// - Looking-Glass:
// Start from at the end of the pattern
// Character-Jump

/*
a pattern matching algorithm
rithm
  rithm
       rithm
            rithm
			     rithm
				      rithm
					   rithm

*/

// Kommt das erste verglichene Symbol im Muster nicht vor,
// so wird das Muster um m verschoben.

// Falls bei eine Ungleichheit, dass falsche Zeichen an eine
// anderen Stelle im Muster vorkommt,
// dann wird das Muster nur so weit geschoben, bis dieses Vorkommt.

// Bei einer negativen Verschiebung wird um eins verschoben.
// Laut buch O(m+s)
// Hier aber O(m)???
func BoyerMooreLastOccurence(p string) map[string]int {
	lastOccurrences := make(map[string]int)

	for i := 0; i < len(p); i++ {
		lastOccurrence := lastOccurrences[string(p[i])]
		if i > lastOccurrence {
			lastOccurrences[string(p[i])] = i
		}
	}
	return lastOccurrences
}

// O(n*m+s)
func BoyerMooreMatch(t, p string) int {
	var lastOccurences map[string]int = BoyerMooreLastOccurence(p)
	var m int = len(p)
	var n int = len(t)
	var i int = m - 1 //actual T-index
	var j int = m - 1 //actual P-index

	for {
		if t[i] == p[j] {
			if j == 0 {
				return i //match at i
			} else {
				i = i - 1
				j = j - 1
			}
		} else {
			//character jump
			l := lastOccurences[string(t[i])]
			i = i + m - min(j, 1+l)
			j = m - 1
		}
		if i > n-1 {
			break
		}
	}

	return -1
}
