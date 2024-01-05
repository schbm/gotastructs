package tree

import "golang.org/x/exp/constraints"

type Tree[K constraints.Ordered, V comparable] interface {
	Find(key K) (V, error)
	Insert(key K, value V)
	Remove(key K) error
}
