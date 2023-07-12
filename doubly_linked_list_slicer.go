package gotastructs

func (l *DoublyLinkedList) ToSlice() []Element {
	slice := make([]Element, l.Size())
	current := l.Head()
	for i := 0; i < l.Size(); i++ {
		slice[i] = current.Value()
		current = current.Next()
	}
	return slice
}
