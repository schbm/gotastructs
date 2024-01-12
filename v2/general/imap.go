package general

import (
	"golang.org/x/exp/constraints"
)

// Map represents a collection of key-value pairs.
// Keys can not contain duplicates.
type Map[K constraints.Ordered, V any] interface {
	Get(K) (V, error)
	Insert(K, V)
	Remove(K) error
	Size() int
}

// With duplicates
type MultiMap[K constraints.Ordered, V comparable] interface {
	Map[K, V]
	GetAll(K) ([]V, error)
	GetSpecific(K, V) (V, error)
	RemoveSpecific(K, V) error
}
