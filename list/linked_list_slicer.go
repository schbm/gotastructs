package list

import "github.com/schbm/gotastructs"

func (l *LinkedList) ToSlice() []gotastructs.Element {
	slice := make([]gotastructs.Element, l.Size())
	current := l.Head()
	for i := 0; i < l.Size(); i++ {
		slice[i] = current.Value()
		current = current.Next()
	}
	return slice
}
