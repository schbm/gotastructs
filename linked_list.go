package gotastructs

import "errors"

type LinkedListElement struct {
	value Element            //concrete value interface
	next  *LinkedListElement //pointer to next element
}

func (e *LinkedListElement) Value() Element {
	return e.value //return concrete value
}

// next function
func (e *LinkedListElement) Next() *LinkedListElement {
	return e.next //return next element pointer
}

type LinkedList struct {
	head *LinkedListElement //pointer to first element
	size int                //size of list
	tail *LinkedListElement //pointer to last element
}

// NewLinkedList returns a new LinkedList
func NewLinkedList() *LinkedList {
	return &LinkedList{nil, 0, nil}
}

// Append adds a new element to the end of the list
func (l *LinkedList) Append(value Element) {
	if l.IsEmpty() {
		l.head = &LinkedListElement{value, nil}
		l.tail = l.Head()
		l.size++
		return
	}

	l.tail.next = &LinkedListElement{value, nil}
	l.tail = l.tail.Next()
	l.size++
	return
}

func (l *LinkedList) Insert(value Element, index int) ListError {
	if index < 0 || index >= l.Size() {
		return errors.New("index out of bounds")
	}

	if index == 0 {
		l.head = &LinkedListElement{value, l.Head()}
		l.size++
		return nil
	}

	if index == l.Size()-1 { // if index is the last element, just append
		l.Append(value)
		return nil
	}

	//otherwise iterate until index -1 and insert next
	current := l.iterateUntil(index - 1)
	current.next = &LinkedListElement{value, current.Next()}
	l.size++
	return nil
}

func (l *LinkedList) Remove(index int) ListError {
	if index < 0 || index >= l.Size() {
		return errors.New("index out of bounds")
	}

	if index == 0 { // if index is the first element, just remove
		l.head = l.Head().Next()
		l.size--
		if l.IsEmpty() { //if list is now empty set tail nil
			l.tail = nil
		}
		return nil
	}

	//if last element
	if index == l.Size()-1 {
		l.tail = l.iterateUntil(index - 1) //iterate to previous element
		l.tail.next = nil
		l.size--
		return nil
	}

	//otherwise iterate until index -1 and remove next
	//[0,1,2,3,4,5,6,7,8,9] remove 8
	//iterate until index 7
	current := l.iterateUntil(index - 1)
	current.next = current.Next().Next()
	l.size--
	return nil
}

func (l *LinkedList) RemoveElement(value Element) ListError {
	if l.IsEmpty() {
		return errors.New("list is empty")
	}

	if l.Head().Value().Equals(value) {
		l.head = l.Head().Next()
		l.size--
		if l.IsEmpty() {
			l.tail = nil
		}
		return nil
	}

	if l.tail.Value().Equals(value) {
		l.tail = l.iterateUntil(l.Size() - 2)
		l.tail.next = nil
		l.size--
		return nil
	}

	current := l.Head()
	for current.Next() != nil {
		if current.Next().Value().Equals(value) {
			current.next = current.Next().Next()
			l.size--
			return nil
		}
		current = current.Next()
	}
	return errors.New("element not found")
}

func (l *LinkedList) Head() *LinkedListElement {
	return l.head
}

func (l *LinkedList) iterateUntil(index int) *LinkedListElement {
	if index < 0 || index >= l.Size() {
		panic("index out of bounds") //TODO: return error
	}
	current := l.Head()
	for i := 0; i < index; i++ {
		current = current.Next()
	}
	return current
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) IsEmpty() bool {
	return l.Head() == nil
}

func (l *LinkedList) Clear() {
	l.head = nil
	l.size = 0
}

func (l *LinkedList) Contains(value Element) bool {
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

func (l *LinkedList) IndexOf(value Element) (int, ListError) {
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

func (l *LinkedList) Get(index int) (Element, ListError) {
	if index < 0 || index >= l.Size() {
		return nil, errors.New("index out of bounds")
	}

	if index == l.Size()-1 {
		return l.tail.Value(), nil
	}

	current := l.iterateUntil(index)
	return current.Value(), nil
}
