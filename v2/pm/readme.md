# String Pattern Matching

- Brute Force
- Boyer-Moore
- Knuth-Morris-Pratt

Sei $S$ ein String der Länge $m$.
- Substring $S[i...j]$ von $S$ ist eine Subsequenz inklusiv $i$ und $j$.
- Präfix von $P$ ist $S[0..i]$.
- Suffix von $P$ ist $S[i...m - 1]$.
- $T$ Text
- $P$ Pattern, Muster

Pattern-Matching erfodert das effiziente Finden eines Subtrings $S$ aus $T$ der mit $P$ übereinstimmt.

## Boyer-Moore

Depends on two heuristics
- Looking-Glass:
    - Start from at the end of the pattern
- Character-Jump

Kommt das erste verglichene Symbol im Muster nicht vor, so wird das Muster um m verschoben. Falls bei eine Ungleichheit, dass falsche Zeichen an einenanderen Stelle im Muster vorkommt, dann wird das Muster nur so weit geschoben, bis dieses Vorkommt. Bei einer negativen Verschiebung wird um eins verschoben.

### Last Occurence
Analysiert das Pattern $P$ und das Alphabet $\sum{}$, um die last-occurence Funktion $L$ aufzubauen.
Bildet $\sum{}$ auf Ganzzahlen ab, wobei $L(c)$ wie folgt definiert ist:
- $L : \sum{} -> IN_0 \cup {-1}$
- $L : c -> L(c) = max_i \{i | P[i] = c\}$
Diese Funktion $L(c)$ lässt sich als ein Array darstellen, dessen Indices durch numerische Werte des Alphabets gegeben sind.
Die Funktion lässt sich in $O(m+s)$ berechnen mit $m$ Länge von $P$ und $s$ Anzahl Zeichen in Alphabet $\sum{}$.

