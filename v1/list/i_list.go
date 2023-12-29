package list

import (
	"github.com/schbm/gotastructs/general"
)

type ListError interface {
	error
}

type List interface {
	general.Iterable
	general.Slicer
	Append(general.Element)
	Insert(general.Element, int) ListError
	Remove(int) ListError
	RemoveElement(general.Element) ListError
	IndexOf(general.Element) (int, ListError)
	Contains(general.Element) bool
	Get(int) (general.Element, ListError)
	IsEmpty() bool
	Size() int
}
