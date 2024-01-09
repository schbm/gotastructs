package genmap

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type SliceMultiMap[K constraints.Ordered, V comparable] struct {
	m map[K][]V
}

/*
Get(K) V
Insert(K, V)
Remove(K)
GetAll(K) []V
*/
func (multiMap *SliceMultiMap[K, V]) Get(key K) (V, error) {
	values, found := multiMap.m[key]
	if !found {
		var zeroV V
		return zeroV, errors.New("not found")
	}
	if len(values) < 1 {
		var zeroV V
		return zeroV, errors.New("not found")
	}

	return values[0], nil
}

// GetAll returns all the values associated with the given key in the SliceMultiMap.
// If the key is not found, it returns an empty slice and an error.
// If the key is found but has no associated values, it returns an empty slice and an error.
// The function returns the slice (address) of the data
func (multiMap *SliceMultiMap[K, V]) GetAll(key K) ([]V, error) {
	values, found := multiMap.m[key]
	if !found {
		return nil, errors.New("not found")
	}
	if len(values) < 1 {
		return nil, errors.New("not found")
	}

	return values, nil
}

func (multiMap *SliceMultiMap[K, V]) Insert(key K, value V) {
	multiMap.m[key] = append(multiMap.m[key], value)
}

func (multiMap *SliceMultiMap[K, V]) RemoveAll(key K) error {
	_, found := multiMap.m[key]
	if !found {
		return errors.New("not found")
	}

	delete(multiMap.m, key)
	return nil
}
