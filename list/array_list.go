package list

import (
	"errors"
	"github.com/schbm/gotastructs"
)

const (
	ARRAY_LIST_INITIAL_SIZE    = 100
	ARRAY_LIST_MULTIPLIER      = 2
	ARRAY_LIST_GROTH_THRESHOLD = 5
)

type ArrayList struct {
	elements []gotastructs.Element
}

func NewArrayList() *ArrayList {
	return &ArrayList{
		elements: make([]gotastructs.Element, 0, ARRAY_LIST_INITIAL_SIZE),
	}
}

func (l *ArrayList) grow() {
	newElements := make([]gotastructs.Element, len(l.elements), cap(l.elements)*ARRAY_LIST_MULTIPLIER)
	copy(newElements, l.elements)
	l.elements = newElements
}

func (l *ArrayList) Append(element gotastructs.Element) {
	if cap(l.elements)-len(l.elements) < ARRAY_LIST_GROTH_THRESHOLD {
		l.grow()
	}

	l.elements = append(l.elements, element)
}

func (l *ArrayList) Insert(element gotastructs.Element, index int) ListError {
	if index < 0 || index > len(l.elements)-1 {
		return errors.New("index out of bounds")
	}

	if cap(l.elements)-len(l.elements) < ARRAY_LIST_GROTH_THRESHOLD {
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

func (l *ArrayList) IndexOf(element gotastructs.Element) (int, ListError) {
	for i, v := range l.elements {
		if v.Equals(element) {
			return i, nil
		}
	}

	return -1, errors.New("element not found")
}

func (l *ArrayList) RemoveElement(element gotastructs.Element) ListError {
	index, err := l.IndexOf(element)
	if err != nil {
		return err
	}

	return l.Remove(index)
}

func (l *ArrayList) Contains(element gotastructs.Element) bool {
	for _, v := range l.elements {
		if v.Equals(element) {
			return true
		}
	}
	return false
}

func (l *ArrayList) Get(index int) (gotastructs.Element, ListError) {
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
