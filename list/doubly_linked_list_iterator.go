package list

import "github.com/schbm/gotastructs"

type DoublyLinkedListIterator struct {
	current *DoublyLinkedListElement
}

func (i *DoublyLinkedListIterator) HasNext() bool {
	return i.current != nil
}

func (i *DoublyLinkedListIterator) Next() gotastructs.Element {
	if !i.HasNext() {
		return nil
	}
	value := i.current.Value()
	i.current = i.current.Next()
	return value
}

func (l *DoublyLinkedList) Iterator() gotastructs.Iterator {
	return &DoublyLinkedListIterator{l.Head()}
}
