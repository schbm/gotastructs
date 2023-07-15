package stack

import "github.com/schbm/gotastructs"

type Stack interface {
	gotastructs.Iterable
	Push(gotastructs.Element)
	Pop() (gotastructs.Element, error)
	Size() int
	Peek() (gotastructs.Element, error)
	IsEmpty() bool
}
