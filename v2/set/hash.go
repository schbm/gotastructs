package set

import "errors"

var _ Set[string] = NewHashSet[string]()

// Set represents a set data structure.
type HashSet[T comparable] map[T]bool

// NewSet creates a new set.
func NewHashSet[T comparable]() HashSet[T] {
	return make(map[T]bool)
}

// Add adds an element to the set.
func (s HashSet[T]) Insert(elem T) error {
	if s[elem] {
		return errors.New("already exists")
	}
	s[elem] = true
	return nil
}

// Remove removes an element from the set.
func (s HashSet[T]) Remove(elem T) error {
	if !s[elem] {
		return errors.New("does not exist")
	}
	delete(s, elem)
	return nil
}

// Contains checks if the set contains a given element.
func (s HashSet[T]) Contains(elem T) bool {
	return s[elem]
}

// Size returns the size of the set.
func (s HashSet[T]) Size() int {
	return len(s)
}

// Clear removes all elements from the set.
func (s HashSet[T]) Clear() {
	for elem := range s {
		delete(s, elem)
	}
}
