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

func (l *ArrayList[V]) grow() {
	newElements := make([]V, len(l.elements), cap(l.elements)*arrLMultiplier)
	copy(newElements, l.elements)
	l.elements = newElements
}

func (l *ArrayList[V]) Insert(element V) {
	if cap(l.elements)-len(l.elements) < arrLGrowthThreshold {
		l.grow()
	}

	l.elements = append(l.elements, element)
}

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

func (l *ArrayList[V]) RemoveFrom(index int) error {
	if index < 0 || index > len(l.elements)-1 {
		return errors.New("index out of bounds")
	}

	l.elements = append(l.elements[:index], l.elements[index+1:]...)

	return nil
}

func (l *ArrayList[V]) IndexOf(element V) (int, error) {
	for i, v := range l.elements {
		if v == element {
			return i, nil
		}
	}

	return -1, errors.New("element not found")
}

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

func (l *ArrayList[V]) Get(index int) (V, error) {
	if index < 0 || index > len(l.elements)-1 {
		var v V
		return v, errors.New("index out of bounds")
	}

	return l.elements[index], nil
}

func (l *ArrayList[V]) IsEmpty() bool {
	return len(l.elements) == 0
}

func (l *ArrayList[V]) Size() int {
	return len(l.elements)
}
