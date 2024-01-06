package general

import (
	"golang.org/x/exp/constraints"
)

type Map[K constraints.Ordered, V comparable] interface {
	Get(K) V
	Insert(K, V)
	Remove(K)
}

type MultiMap[K constraints.Ordered, V comparable] interface {
	Map[K, V]
	FindAll(K) []V
}

type OrderedMultiMap[K constraints.Ordered, V comparable] interface {
	MultiMap[K, V]
	First() V
	Last() V
	Successors(K) []V
	Predecessors(K) []V
}
