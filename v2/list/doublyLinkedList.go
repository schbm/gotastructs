package list

import (
	"errors"

	"github.com/schbm/gotastructs/v2/general"
)

var _ general.List[int] = NewDoublyLinkedList[int]()

// DoublyLinkedListElement is a node with pointers to the next and previous elements
type DoublyLinkedListElement[V comparable] struct {
	value    V
	next     *DoublyLinkedListElement[V]
	previous *DoublyLinkedListElement[V]
}

func (e *DoublyLinkedListElement[V]) Value() V {
	return e.value
}

func (e *DoublyLinkedListElement[V]) Next() *DoublyLinkedListElement[V] {
	return e.next
}

func (e *DoublyLinkedListElement[V]) Previous() *DoublyLinkedListElement[V] {
	return e.previous
}

// DoublyLinkedList is a list of DoublyLinkedListElements
type DoublyLinkedList[V comparable] struct {
	head *DoublyLinkedListElement[V] // first element
	tail *DoublyLinkedListElement[V] // last element
	size int
}

func NewDoublyLinkedList[V comparable]() *DoublyLinkedList[V] {
	return &DoublyLinkedList[V]{nil, nil, 0}
}

func (l *DoublyLinkedList[V]) Insert(value V) {
	if l.IsEmpty() {
		l.head = &DoublyLinkedListElement[V]{value, nil, nil}
		l.tail = l.Head()
		l.size++
		return
	}

	l.Tail().next = &DoublyLinkedListElement[V]{value, nil, l.Tail()}
	l.tail = l.Tail().Next()
	l.size++
	return
}

func (l *DoublyLinkedList[V]) InsertTo(value V, index int) error {
	if index < 0 || index >= l.Size() {
		return errors.New("index out of bounds")
	}

	if index == 0 { // if index is the first element, just prepend
		l.head = &DoublyLinkedListElement[V]{value, l.Head(), nil}
		l.Head().Next().previous = l.Head()
		l.size++
		return nil
	}

	if index == l.Size()-1 { // if index is the last element, just append
		l.Insert(value)
		return nil
	}

	// otherwise, find the element at the index and insert
	current := l.iterateUntil(index)
	current.Previous().next = &DoublyLinkedListElement[V]{value, current, current.Previous()}
	current.previous = current.Previous().next
	l.size++
	return nil
}

func (l *DoublyLinkedList[V]) Remove(index int) error {
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

func (l *DoublyLinkedList[V]) RemoveElement(value V) error {
	if l.IsEmpty() {
		return errors.New("list is empty")
	}
	if l.Head().value == value { // if value is the first element, just remove
		l.head = l.Head().Next()
		if l.IsEmpty() {
			l.tail = nil
		} else {
			l.Head().previous = nil
		}
		l.size--
		return nil
	}
	if l.Tail().Value() == value { // if value is the last element, just remove
		l.tail = l.Tail().Previous()
		l.Tail().next = nil
		l.size--
		return nil
	}

	// otherwise, find the element and remove
	current := l.Head()

	for current != nil && current.Value() != value {
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

func (l *DoublyLinkedList[V]) IndexOf(value V) (int, error) {
	if l.IsEmpty() {
		return -1, errors.New("list is empty")
	}

	current := l.Head()
	for i := 0; current != nil; i++ {
		if current.Value() == value {
			return i, nil
		}
		current = current.Next()
	}
	return -1, errors.New("element not found")
}

func (l *DoublyLinkedList[V]) Contains(value V) bool {
	if l.IsEmpty() {
		return false
	}

	current := l.Head()
	for current != nil {
		if current.Value() == value {
			return true
		}
		current = current.Next()
	}
	return false
}

func (l *DoublyLinkedList[V]) Get(index int) (V, error) {
	if index < 0 || index >= l.Size() {
		var zeroV V
		return zeroV, errors.New("index out of bounds")
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

func (l *DoublyLinkedList[V]) IsEmpty() bool {
	return l.head == nil
}

func (l *DoublyLinkedList[V]) Size() int {
	return l.size
}

func (l *DoublyLinkedList[V]) Head() *DoublyLinkedListElement[V] {
	return l.head
}

func (l *DoublyLinkedList[V]) Tail() *DoublyLinkedListElement[V] {
	return l.tail
}

func (l *DoublyLinkedList[V]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *DoublyLinkedList[V]) iterateUntil(n int) *DoublyLinkedListElement[V] {
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
