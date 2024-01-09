package general

import (
	"golang.org/x/exp/constraints"
)

type Map[K constraints.Ordered, V comparable] interface {
	Get(K) (V, error)
	Insert(K, V)
	Remove(K) error
}

type MultiMap[K constraints.Ordered, V comparable] interface {
	Map[K, V]
	GetAll(K) ([]V, error)
	GetSpecific(K, V) (V, error)
	RemoveAll(K) error
	RemoveSpecific(K, V) error
}

type OrderedMultiMap[K constraints.Ordered, V comparable] interface {
	MultiMap[K, V]
	First() (V, error)
	Last() (V, error)
	Successors(K) ([]V, error)
	Predecessors(K) ([]V, error)
}
