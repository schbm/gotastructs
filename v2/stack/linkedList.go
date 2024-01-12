package stack

import (
	"github.com/schbm/gotastructs/v2/general"
	"github.com/schbm/gotastructs/v2/list"
)

var _ general.Stack[string] = &LinkedListStack[string]{}

type LinkedListStack[V comparable] struct {
	List *list.LinkedList[V]
}

func NewLinkedListStack[V comparable]() *LinkedListStack[V] {
	return &LinkedListStack[V]{
		List: list.NewLinkedList[V](false),
	}

}

func (l *LinkedListStack[V]) Push(value V) {
	l.List.InsertLast(value)
}

func (l *LinkedListStack[V]) Pop() (V, error) {
	value, err := l.List.GetLast()
	if err != nil {
		var zeroV V
		return zeroV, err
	}
	err = l.List.RemoveFrom(l.List.Size() - 1)
	if err != nil {
		var zeroV V
		return zeroV, err
	}
	return value, nil
}

func (l *LinkedListStack[V]) Size() int {
	return l.Size()
}

func (l *LinkedListStack[V]) Top() (V, error) {
	return l.List.GetLast()
}
