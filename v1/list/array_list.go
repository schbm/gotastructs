package list

import (
	"errors"
	"github.com/schbm/gotastructs/general"
)

const (
	arrLInitSize        = 100 // initial size of the array
	arrLMultiplier      = 2   // multiplier for the array when it needs to grow
	arrLGrowthThreshold = 5   // threshold of free elements for when the array needs to grow
)

/*
ArrayList is a list implementation that uses an array as its underlying data structure.
*/
type ArrayList struct {
	elements []general.Element
}

// NewArrayList creates a new ArrayList.
func NewArrayList() *ArrayList {
	return &ArrayList{
		elements: make([]general.Element, 0, arrLInitSize),
	}
}

func (l *ArrayList) grow() {
	newElements := make([]general.Element, len(l.elements), cap(l.elements)*arrLMultiplier)
	copy(newElements, l.elements)
	l.elements = newElements
}

func (l *ArrayList) Append(element general.Element) {
	if cap(l.elements)-len(l.elements) < arrLGrowthThreshold {
		l.grow()
	}

	l.elements = append(l.elements, element)
}

func (l *ArrayList) Insert(element general.Element, index int) ListError {
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

func (l *ArrayList) Remove(index int) ListError {
	if index < 0 || index > len(l.elements)-1 {
		return errors.New("index out of bounds")
	}

	l.elements = append(l.elements[:index], l.elements[index+1:]...)

	return nil
}

func (l *ArrayList) IndexOf(element general.Element) (int, ListError) {
	for i, v := range l.elements {
		if v.Equals(element) {
			return i, nil
		}
	}

	return -1, errors.New("element not found")
}

func (l *ArrayList) RemoveElement(element general.Element) ListError {
	index, err := l.IndexOf(element)
	if err != nil {
		return err
	}

	return l.Remove(index)
}

func (l *ArrayList) Contains(element general.Element) bool {
	for _, v := range l.elements {
		if v.Equals(element) {
			return true
		}
	}
	return false
}

func (l *ArrayList) Get(index int) (general.Element, ListError) {
	if index < 0 || index > len(l.elements)-1 {
		return nil, errors.New("index out of bounds")
	}

	return l.elements[index], nil
}

func (l *ArrayList) IsEmpty() bool {
	return len(l.elements) == 0
}

func (l *ArrayList) Size() int {
	return len(l.elements)
}

type ArrayListIterator struct {
	list  []general.Element
	index int
}

func (l *ArrayListIterator) Next() general.Element {
	if l.HasNext() {
		return nil
	}
	el := l.list[l.index]
	l.index++
	return el
}

func (l *ArrayListIterator) HasNext() bool {
	return l.index < len(l.list)
}

func (l *ArrayList) Iterator() general.Iterator {
	return &ArrayListIterator{
		list:  l.elements,
		index: 0,
	}
}

func (l *ArrayList) ToSlice() []general.Element {
	return l.elements
}
