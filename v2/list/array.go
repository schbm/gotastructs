package list

import (
	"errors"

	"github.com/schbm/gotastructs/v2/general"
)

var _ general.List[int] = &ArrayList[int]{}

const (
	arrLInitSize        = 100 // initial size of the array
	arrLMultiplier      = 2   // multiplier for the array when it needs to grow
	arrLGrowthThreshold = 5   // threshold of free elements for when the array needs to grow
)

// ArrayList is a generic list implementation that stores elements of type V.
type ArrayList[V comparable] struct {
	elements []V
}

// NewArrayList creates a new ArrayList.
func NewArrayList[V comparable]() *ArrayList[V] {
	return &ArrayList[V]{
		elements: make([]V, 0, arrLInitSize),
	}
}

// grow increases the capacity of the ArrayList by creating a new slice with a larger capacity
// and copying the existing elements into it.
func (l *ArrayList[V]) grow() {
	newElements := make([]V, len(l.elements), cap(l.elements)*arrLMultiplier)
	copy(newElements, l.elements)
	l.elements = newElements
}

// Insert adds an element to the ArrayList.
// If the capacity of the ArrayList is not sufficient to accommodate the new element,
// the ArrayList will be grown to ensure enough capacity.
func (l *ArrayList[V]) Insert(element V) {
	if cap(l.elements)-len(l.elements) < arrLGrowthThreshold {
		l.grow()
	}

	l.elements = append(l.elements, element)
}

// InsertTo inserts an element at the specified index in the ArrayList.
// If the index is out of bounds, it returns an error.
// If the capacity of the ArrayList is not sufficient, it automatically grows the underlying array.
// The existing elements from the specified index onwards are shifted to the right to accommodate the new element.
// The element is then inserted at the specified index.
func (l *ArrayList[V]) InsertTo(element V, index int) error {
	if index < 0 || index > len(l.elements)-1 {
		return errors.New("index out of bounds")
	}

	if cap(l.elements)-len(l.elements) < arrLGrowthThreshold {
		l.grow()
	}

	l.elements = append(l.elements[:index+1], l.elements[index:]...)
	l.elements[index] = element

	return nil
}

// RemoveFrom removes an element from the ArrayList at the specified index.
// It returns an error if the index is out of bounds.
func (l *ArrayList[V]) RemoveFrom(index int) error {
	if index < 0 || index > len(l.elements)-1 {
		return errors.New("index out of bounds")
	}
	l.elements = append(l.elements[:index], l.elements[index+1:]...)
	return nil
}

// IndexOf returns the index of the first occurrence of the specified element in the ArrayList.
// If the element is found, the function returns the index of the element.
// If the element is not found, the function returns -1 and an error indicating that the element was not found.
func (l *ArrayList[V]) IndexOf(element V) (int, error) {
	for i, v := range l.elements {
		if v == element {
			return i, nil
		}
	}

	return -1, errors.New("element not found")
}

// Remove removes the first occurrence of the specified element from the ArrayList.
// It returns an error if the element is not found in the list.
func (l *ArrayList[V]) Remove(element V) error {
	index, err := l.IndexOf(element)
	if err != nil {
		return err
	}

	return l.RemoveFrom(index)
}

func (l *ArrayList[V]) Contains(element V) bool {
	for _, v := range l.elements {
		if v == element {
			return true
		}
	}
	return false
}

// Get retrieves the first occurrence of the specified element in the ArrayList.
// It returns the element and nil error if found, otherwise it returns the zero value of type V and the error.
func (l *ArrayList[V]) Get(element V) (V, error) {
	n, err := l.IndexOf(element)
	if err != nil {
		var zeroV V
		return zeroV, err
	}
	return l.elements[n], nil
}

// IsEmpty checks if the ArrayList is empty.
// It returns true if the ArrayList is empty, otherwise false.
func (l *ArrayList[V]) IsEmpty() bool {
	return len(l.elements) == 0
}

// Size returns the number of elements in the ArrayList.
func (l *ArrayList[V]) Size() int {
	return len(l.elements)
}

// GetFrom returns the element at the specified index in the ArrayList.
// It returns an error if the index is out of bounds.
func (l *ArrayList[V]) GetFrom(index int) (V, error) {
	if index < 0 || index >= len(l.elements) {
		var zeroV V
		return zeroV, errors.New("index out of bounds")
	}
	return l.elements[index], nil
}
