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

// NewDoublyLinkedList creates a new instance of DoublyLinkedList.
// It initializes the list with nil head and tail nodes and a length of 0.
func NewDoublyLinkedList[V comparable]() *DoublyLinkedList[V] {
	return &DoublyLinkedList[V]{nil, nil, 0}
}

// Insert inserts a new element with the given value at the end of the doubly linked list.
// If the list is empty, the new element becomes both the head and the tail of the list.
// Otherwise, the new element is added after the current tail element.
func (l *DoublyLinkedList[V]) Insert(value V) {
	// TODO: Avoid this check?
	if l.IsEmpty() {
		l.head = &DoublyLinkedListElement[V]{value, nil, nil}
		l.tail = l.Head()
		l.size++
		return
	}

	l.Tail().next = &DoublyLinkedListElement[V]{value, nil, l.Tail()}
	l.tail = l.Tail().Next()
	l.size++
}

// InsertTo inserts a new element with the given value at the specified index in the doubly linked list.
// If the index is out of bounds, it returns an error.
// If the index is the first element, the new element is prepended to the list.
// If the index is the last element, the new element is appended to the list.
// Otherwise, the new element is inserted at the specified index.
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
	current, _ := l.iterateUntil(index)
	current.Previous().next = &DoublyLinkedListElement[V]{value, current, current.Previous()}
	current.previous = current.Previous().next
	l.size++
	return nil
}

// RemoveFrom removes the element at the specified index from the doubly linked list.
// It returns an error if the index is out of bounds.
func (l *DoublyLinkedList[V]) RemoveFrom(index int) error {
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
	current, _ := l.iterateUntil(index)
	current.Previous().next = current.Next()
	current.Next().previous = current.Previous()
	l.size--
	return nil
}

// Remove removes the first occurrence of the specified value from the doubly linked list.
// If the value is found and removed, it returns nil. If the list is empty, it returns an error.
// If the value is the first element, it is removed and the head of the list is updated.
// If the value is the last element, it is removed and the tail of the list is updated.
// If the value is found in the middle of the list, it is removed and the previous and next pointers are updated.
// If the value is not found in the list, it returns an error.
func (l *DoublyLinkedList[V]) Remove(value V) error {
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

// IndexOf returns the index of the first occurrence of the specified value in the doubly linked list.
// It returns -1 and an error if the list is empty or if the element is not found.
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

// Contains checks if the doubly linked list contains the specified value.
// It returns true if the value is found, otherwise it returns false.
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

func (l *DoublyLinkedList[V]) Get(element V) (V, error) {
	if l.IsEmpty() {
		var zeroV V
		return zeroV, errors.New("list is empty")
	}

	current := l.Head()
	for i := 0; current != nil; i++ {
		if current.Value() == element {
			return current.value, nil
		}
		current = current.Next()
	}
	return current.value, errors.New("element not found")
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

func (l *DoublyLinkedList[V]) iterateUntil(n int) (*DoublyLinkedListElement[V], error) {
	if n < 0 || n >= l.Size() {
		return nil, errors.New("index out of bounds")
	}

	current := l.Head()
	//if higher than half
	if n > l.Size()/2 {
		current = l.Tail()
		for i := l.Size() - 1; i > n; i-- {
			current = current.Previous()
		}
		return current, nil
	}

	for i := 0; i < n; i++ {
		current = current.Next()
	}
	return current, nil
}

// Returns the value at the index from the doubly linked list
func (l *DoublyLinkedList[V]) GetFrom(index int) (V, error) {
	if index < 0 || index >= l.Size() {
		var zeroV V
		return zeroV, errors.New("index out of bounds")
	}
	val, err := l.iterateUntil(index)
	if err != nil {
		var zeroV V
		return zeroV, err
	}
	return val.Value(), nil
}
