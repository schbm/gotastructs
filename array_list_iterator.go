package gotastructs

type ArrayListIterator struct {
	list  []Element
	index int
}

func (l *ArrayListIterator) Next() Element {
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

func (l *ArrayList) Iterator() Iterator {
	return &ArrayListIterator{
		list:  l.elements,
		index: 0,
	}
}
