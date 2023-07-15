package list

import "github.com/schbm/gotastructs"

type ArrayListIterator struct {
	list  []gotastructs.Element
	index int
}

func (l *ArrayListIterator) Next() gotastructs.Element {
	if l.index >= len(l.list) {
		return nil
	}
	el := l.list[l.index]
	l.index++
	return el
}

func (l *ArrayListIterator) HasNext() bool {
	return l.index < len(l.list)
}

func (l *ArrayList) Iterator() gotastructs.Iterator {
	return &ArrayListIterator{
		list:  l.elements,
		index: 0,
	}
}
