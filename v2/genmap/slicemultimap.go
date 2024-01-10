package genmap

import (
	"errors"

	"github.com/schbm/gotastructs/v2/general"
)

var _ general.MultiMap[string, int] = &SliceMultiMap[string, int]{}

func NewSliceMultiMap[K comparable, V comparable]() *SliceMultiMap[K, V] {
	newMap := make(map[K][]V)
	return &SliceMultiMap[K, V]{
		M: newMap,
	}
}

type SliceMultiMap[K comparable, V comparable] struct {
	M     map[K][]V
	Count int
}

/*
Get(K) V
Insert(K, V)
Remove(K)
GetAll(K) []V
*/
func (multiMap *SliceMultiMap[K, V]) Get(key K) (V, error) {

	values, found := multiMap.M[key]
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
	if len(multiMap.M) < 1 {
		return nil, errors.New("map is empty")
	}

	values, found := multiMap.M[key]
	if !found {
		return nil, errors.New("not found")
	}
	if len(values) < 1 {
		return nil, errors.New("not found")
	}
	result := make([]V, len(values))
	copy(result, values)
	return result, nil
}

func (multiMap *SliceMultiMap[K, V]) Insert(key K, value V) {
	slice, err := multiMap.GetAll(key)
	if err != nil { // no value, create a new slice
		multiMap.M[key] = []V{value}
	}
	// append
	multiMap.M[key] = append(slice, value)
	multiMap.Count++
}

func (multiMap *SliceMultiMap[K, V]) Remove(key K) error {
	values, found := multiMap.M[key]
	if !found {
		return errors.New("not found")
	}
	delete(multiMap.M, key)
	multiMap.Count -= len(values)
	return nil
}

func (multiMap *SliceMultiMap[K, V]) GetSpecific(key K, value V) (V, error) {
	values, err := multiMap.GetAll(key)
	if err != nil {
		var zeroV V
		return zeroV, err
	}
	for _, v := range values {
		if v == value {
			return v, nil
		}
	}
	var zeroV V
	return zeroV, errors.New("item not found")
}

func (multiMap *SliceMultiMap[K, V]) RemoveSpecific(key K, value V) error {
	values, err := multiMap.GetAll(key)
	if err != nil {
		return err
	}
	// if only one item and item is the same delete
	if len(values) == 1 {
		for _, v := range values {
			if value == v {
				err := multiMap.Remove(key)
				if err != nil {
					return err
				}
				return nil
			}
			return errors.New("value does not match")
		}
	}
	//otherwhise iterate
	for i, v := range values {
		if value == v {
			newValues, err := RemoveAt[K, V](i, values)
			if err != nil {
				return err
			}
			multiMap.M[key] = newValues
			multiMap.Count--
			return nil
		}
	}
	return errors.New("element not matched")
}

func (multiMap *SliceMultiMap[K, V]) Size() int {
	return multiMap.Count
}

// 4 Cases: min => len >= 2
// empty list or sm 2
// first item
// last item
// in the middle
func RemoveAt[K comparable, V comparable](i int, valList []V) ([]V, error) {
	if i < 0 || i > len(valList)-1 {
		return nil, errors.New("index out of bounds")
	}
	if len(valList) < 2 {
		return nil, errors.New("RemoveAt expects slice of len >= 2")
	}
	// start
	if i == 0 {
		return valList[1:], nil
	}
	// send
	if i == len(valList)-1 {
		return valList[:len(valList)-1], nil
	}
	// middle
	return append(valList[:i], valList[i+1:]...), nil
}
