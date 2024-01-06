# Trees

Weitere Informationen in den jeweiligen Implementationen.

- Binäre Bäume
    - AVL
    - Splay
    - ..

Multimaps sind ungeordnete Abbildungen eines Keys (k) zu einem Wert (v).
Keys können dublikate enthalten.

Geordnete Multimaps folgen einer vollständigen Ordnungsrelation 
$k_m <= k_n$
Sie liefern zusätzliche Operationen:
- First()
- Last()
- Successor()
- Predecessor()

## Binäre Suche
Kann auf einer Multimap mittels array-basierter Sequenz, sortiert nach k, realisiert werden.
- Terminiert nach $O(log_n)$ Schritten. (Halbiert Anzahl der Kandidaten)

## Suchtabelle
Ist eine Multimap, welche mithilfe einer sortierten Sequenz implementiert wird.
Die Einträge der Multimap werden in einer Array-basierten Sequenz abgespeichert, sortiert nach Schlüssel.
- find() benötigt $O(log_n)$ falls Binärsuche eingesetzt wird.
- insert() im schlimmsten Fall $O(n)$, da alle Einträge um einen Platz verschoben werden müssen.
- remove() $O(n)$ aus dem selben Grund wie insert().
Suchtabellen nur dann effektiv, wenn die Multimap klein ist und vor allem Such-operationen ausgeführt werden.

