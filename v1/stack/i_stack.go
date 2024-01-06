package stack

import (
	"github.com/schbm/gotastructs/v1/general"
)

type Stack interface {
	general.Iterable
	general.Slicer
	Push(general.Element)
	Pop() (general.Element, error)
	Size() int
	Peek() (general.Element, error)
	IsEmpty() bool
}
