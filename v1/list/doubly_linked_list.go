package list

import (
	"errors"

	"github.com/schbm/gotastructs/v1/general"
)

// DoublyLinkedListElement is a node with pointers to the next and previous elements
type DoublyLinkedListElement struct {
	value    general.Element
	next     *DoublyLinkedListElement
	previous *DoublyLinkedListElement
}

func (e *DoublyLinkedListElement) Value() general.Element {
	return e.value
}

func (e *DoublyLinkedListElement) Next() *DoublyLinkedListElement {
	return e.next
}

func (e *DoublyLinkedListElement) Previous() *DoublyLinkedListElement {
	return e.previous
}

// DoublyLinkedList is a list of DoublyLinkedListElements
type DoublyLinkedList struct {
	head *DoublyLinkedListElement // first element
	tail *DoublyLinkedListElement // last element
	size int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{nil, nil, 0}
}

func (l *DoublyLinkedList) Append(value general.Element) {
	if l.IsEmpty() {
		l.head = &DoublyLinkedListElement{value, nil, nil}
		l.tail = l.Head()
		l.size++
		return
	}

	l.Tail().next = &DoublyLinkedListElement{value, nil, l.Tail()}
	l.tail = l.Tail().Next()
	l.size++
	return
}

func (l *DoublyLinkedList) Insert(value general.Element, index int) ListError {
	if index < 0 || index >= l.Size() {
		return errors.New("index out of bounds")
	}

	if index == 0 { // if index is the first element, just prepend
		l.head = &DoublyLinkedListElement{value, l.Head(), nil}
		l.Head().Next().previous = l.Head()
		l.size++
		return nil
	}

	if index == l.Size()-1 { // if index is the last element, just append
		l.Append(value)
		return nil
	}

	// otherwise, find the element at the index and insert
	current := l.iterateUntil(index)
	current.Previous().next = &DoublyLinkedListElement{value, current, current.Previous()}
	current.previous = current.Previous().next
	l.size++
	return nil
}

func (l *DoublyLinkedList) Remove(index int) ListError {
	if index < 0 || index >= l.Size() {
		return errors.New("index out of bounds")
	}

	if index == 0 { // if index is the first element, just remove
		l.head = l.Head().Next()
		if l.IsEmpty() {
			l.tail = nil
		} else {
			l.Head().previous = nil
		}
		l.size--
		return nil
	}

	if index == l.Size()-1 { // if index is the last element, just remove
		l.tail = l.tail.Previous()
		l.tail.next = nil
		l.size--
		return nil
	}

	// otherwise, find the element at the index and remove
	current := l.iterateUntil(index)
	current.Previous().next = current.Next()
	current.Next().previous = current.Previous()
	l.size--
	return nil
}

func (l *DoublyLinkedList) RemoveElement(value general.Element) ListError {
	if l.IsEmpty() {
		return errors.New("list is empty")
	}
	if l.Head().value.Equals(value) { // if value is the first element, just remove
		l.head = l.Head().Next()
		if l.IsEmpty() {
			l.tail = nil
		} else {
			l.Head().previous = nil
		}
		l.size--
		return nil
	}
	if l.Tail().Value().Equals(value) { // if value is the last element, just remove
		l.tail = l.Tail().Previous()
		l.Tail().next = nil
		l.size--
		return nil
	}

	// otherwise, find the element and remove
	current := l.Head()
	for current != nil && !current.Value().Equals(value) {
		current = current.Next()
	}
	if current == nil {
		return errors.New("element not found")
	}
	current.Previous().next = current.Next()
	current.Next().previous = current.Previous()

	l.size--
	return nil
}

func (l *DoublyLinkedList) IndexOf(value general.Element) (int, ListError) {
	if l.IsEmpty() {
		return -1, errors.New("list is empty")
	}

	current := l.Head()
	for i := 0; current != nil; i++ {
		if current.Value().Equals(value) {
			return i, nil
		}
		current = current.Next()
	}
	return -1, errors.New("element not found")
}

func (l *DoublyLinkedList) Contains(value general.Element) bool {
	if l.IsEmpty() {
		return false
	}

	current := l.Head()
	for current != nil {
		if current.Value().Equals(value) {
			return true
		}
		current = current.Next()
	}
	return false
}

func (l *DoublyLinkedList) Get(index int) (general.Element, ListError) {
	if index < 0 || index >= l.Size() {
		return nil, errors.New("index out of bounds")
	}

	if index == 0 { // if index is the first element, just return
		return l.Head().Value(), nil
	}

	if index == l.Size()-1 { // if index is the last element, just return
		return l.Tail().Value(), nil
	}

	current := l.iterateUntil(index)
	return current.Value(), nil
}

func (l *DoublyLinkedList) IsEmpty() bool {
	return l.head == nil
}

func (l *DoublyLinkedList) Size() int {
	return l.size
}

func (l *DoublyLinkedList) Head() *DoublyLinkedListElement {
	return l.head
}

func (l *DoublyLinkedList) Tail() *DoublyLinkedListElement {
	return l.tail
}

func (l *DoublyLinkedList) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *DoublyLinkedList) iterateUntil(n int) *DoublyLinkedListElement {
	current := l.Head()
	//if higher than half
	if n > l.Size()/2 {
		current = l.Tail()
		for i := l.Size() - 1; i > n; i-- {
			current = current.Previous()
		}
		return current
	}

	for i := 0; i < n; i++ {
		current = current.Next()
	}
	return current
}

type DoublyLinkedListIterator struct {
	current *DoublyLinkedListElement
}

func (i *DoublyLinkedListIterator) HasNext() bool {
	return i.current != nil
}

func (i *DoublyLinkedListIterator) Next() general.Element {
	if !i.HasNext() {
		return nil
	}
	value := i.current.Value()
	i.current = i.current.Next()
	return value
}

func (l *DoublyLinkedList) Iterator() general.Iterator {
	return &DoublyLinkedListIterator{l.Head()}
}

func (l *DoublyLinkedList) ToSlice() []general.Element {
	slice := make([]general.Element, l.Size())
	current := l.Head()
	for i := 0; i < l.Size(); i++ {
		slice[i] = current.Value()
		current = current.Next()
	}
	return slice
}
