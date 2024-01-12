package general

import "golang.org/x/exp/constraints"

type Tree[K constraints.Ordered, V comparable] interface {
	Map[K, V]
}

type SearchTree[K constraints.Ordered, V comparable] interface {
	//OrderedMap
}
