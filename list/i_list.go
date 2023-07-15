package list

import "github.com/schbm/gotastructs"

type ListError interface {
	error
}

type List interface {
	gotastructs.Iterable
	gotastructs.Slicer
	Append(gotastructs.Element)
	Insert(gotastructs.Element, int) ListError
	Remove(int) ListError
	RemoveElement(gotastructs.Element) ListError
	IndexOf(gotastructs.Element) (int, ListError)
	Contains(gotastructs.Element) bool
	Get(int) (gotastructs.Element, ListError)
	IsEmpty() bool
	Size() int
}
