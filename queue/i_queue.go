package queue

import "github.com/schbm/gotastructs"

type Queue interface {
	gotastructs.Iterable
	Append(gotastructs.Element)
	Remove() gotastructs.Element
	Peek() gotastructs.Element
	Size() int
	IsEmpty() bool
}
