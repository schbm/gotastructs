package list

import (
	"errors"

	"github.com/schbm/gotastructs/v2/general"
)

var _ general.List[int] = &LinkedList[int]{}

// LinkedListNode is a single element in a linked list
// Arrays have better cache locality compared to linked lists.
type LinkedListNode[V comparable] struct {
	value V                  //concrete value interface
	next  *LinkedListNode[V] //pointer to next element
}

func (e *LinkedListNode[V]) Value() V {
	return e.value //return concrete value
}

// next function
func (e *LinkedListNode[V]) Next() *LinkedListNode[V] {
	return e.next //return next element pointer
}

type LinkedList[V comparable] struct {
	head              *LinkedListNode[V] //pointer to first element
	size              int                //size of list
	tail              *LinkedListNode[V] //pointer to last element
	heuristicApproach bool
}

// NewLinkedList returns a new LinkedList
func NewLinkedList[V comparable](enableHeuristicApproach bool) *LinkedList[V] {
	return &LinkedList[V]{
		head:              nil,
		size:              0,
		tail:              nil,
		heuristicApproach: enableHeuristicApproach,
	}
}

// Append adds a new element to the end of the list
func (l *LinkedList[V]) Insert(value V) {
	if l.IsEmpty() {
		l.head = &LinkedListNode[V]{value, nil}
		l.tail = l.Head()
		l.size++
		return
	}

	l.tail.next = &LinkedListNode[V]{value, nil}
	l.tail = l.tail.Next()
	l.size++
	return
}

func (l *LinkedList[V]) InsertTo(value V, index int) error {
	if index < 0 || index >= l.Size() {
		return errors.New("index out of bounds")
	}

	if index == 0 {
		l.head = &LinkedListNode[V]{value, l.Head()}
		l.size++
		return nil
	}

	if index == l.Size()-1 { // if index is the last element, just append
		l.Insert(value)
		return nil
	}

	//otherwise iterate until index -1 and insert next
	current, err := l.iterateUntil(index - 1)
	if err != nil {
		return err
	}
	current.next = &LinkedListNode[V]{value, current.Next()}
	l.size++
	return nil
}

func (l *LinkedList[V]) RemoveFrom(index int) error {
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
		//iterate to previous element
		el, err := l.iterateUntil(index - 1)
		if err != nil {
			return err
		}
		l.tail = el
		l.tail.next = nil
		l.size--
		return nil
	}

	//otherwise iterate until index -1 and remove next
	//[0,1,2,3,4,5,6,7,8,9] remove 8
	//iterate until index 7
	current, err := l.iterateUntil(index - 1)
	if err != nil {
		return err
	}
	current.next = current.Next().Next()
	l.size--
	return nil
}

func (l *LinkedList[V]) Remove(value V) error {
	if l.IsEmpty() {
		return errors.New("list is empty")
	}

	if l.Head().Value() == value {
		l.head = l.Head().Next()
		l.size--
		if l.IsEmpty() {
			l.tail = nil
		}
		return nil
	}

	if l.tail.Value() == value {
		el, err := l.iterateUntil(l.Size() - 2)
		if err != nil {
			return err
		}
		l.tail = el
		l.tail.next = nil
		l.size--
		return nil
	}

	current := l.Head()
	for current.Next() != nil {
		if current.Next().Value() == value {
			current.next = current.Next().Next()
			l.size--
			return nil
		}
		current = current.Next()
	}
	return errors.New("element not found")
}

func (l *LinkedList[V]) Head() *LinkedListNode[V] {
	return l.head
}

func (l *LinkedList[V]) iterateUntil(index int) (*LinkedListNode[V], error) {
	if index < 0 || index >= l.Size() {
		return nil, errors.New("index out of bounds")
	}
	current := l.Head()
	for i := 0; i < index; i++ {
		current = current.Next()
	}
	return current, nil
}

func (l *LinkedList[V]) Size() int {
	return l.size
}

func (l *LinkedList[V]) IsEmpty() bool {
	return l.Head() == nil
}

func (l *LinkedList[V]) Clear() {
	l.head = nil
	l.size = 0
}

func (l *LinkedList[V]) Contains(value V) bool {
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

func (l *LinkedList[V]) IndexOf(value V) (int, error) {
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

func (l *LinkedList[V]) Get(index int) (V, error) {
	if index < 0 || index >= l.Size() {
		var zeroV V
		return zeroV, errors.New("index out of bounds")
	}

	if index == l.Size()-1 {
		return l.tail.Value(), nil
	}

	if l.heuristicApproach {
		if index != 0 { //head contains no prev
			//swap found element to first position
			prevNode, err := l.iterateUntil(index - 1) //since index != 0 it is possible
			if err != nil {
				var zeroV V
				return zeroV, err
			}
			resultNode := prevNode.Next()
			//remove node from current position
			prevNode.next = resultNode.next //replace current loc with next
			//insert at position 0
			//this is always possible
			//since atleast 2 nodes exist
			resultNode.next = l.head
			l.head = resultNode
			//if it was the last element update the tail
			if index == l.Size()-1 {
				l.tail = prevNode
			}
			return resultNode.Value(), nil
		}
	}

	current, err := l.iterateUntil(index)
	if err != nil {
		var zeroV V
		return zeroV, err
	}
	return current.Value(), nil
}
