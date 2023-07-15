package list

import "github.com/schbm/gotastructs"

type LinkedListIterator struct {
	current *LinkedListElement
}

func (i *LinkedListIterator) HasNext() bool {
	return i.current != nil
}

func (i *LinkedListIterator) Next() gotastructs.Element {

	if !i.HasNext() {
		return nil
	}

	value := i.current.Value()
	i.current = i.current.Next()
	return value
}

func (l *LinkedList) Iterator() gotastructs.Iterator {
	return &LinkedListIterator{l.Head()}
}
