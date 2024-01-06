package queue

import (
	"github.com/schbm/gotastructs/v1/general"
	"github.com/schbm/gotastructs/v1/list"
)

type GeneralQueue struct {
	list list.List
}

func NewQueue(list list.List) *GeneralQueue {
	return &GeneralQueue{
		list: list, // Assuming DoublyLinkedList implements List interface
	}
}

func (q *GeneralQueue) Append(element general.Element) {
	q.list.Append(element)
}

func (q *GeneralQueue) Remove() general.Element {
	if q.IsEmpty() {
		return nil
	}
	element, _ := q.list.Get(0)
	err := q.list.Remove(0)
	if err != nil {
		panic("underlying list returned error")
	}
	return element
}

func (q *GeneralQueue) Peek() general.Element {
	if q.IsEmpty() {
		return nil
	}
	element, _ := q.list.Get(0)
	return element
}

func (q *GeneralQueue) Size() int {
	return q.list.Size()
}

func (q *GeneralQueue) IsEmpty() bool {
	return q.list.IsEmpty()
}

// Remaining methods to satisfy the general.Iterable interface
func (q *GeneralQueue) Iterator() general.Iterator {
	return q.list.Iterator()
}

func (q *GeneralQueue) ToSlice() []general.Element {
	return q.list.ToSlice()
}
