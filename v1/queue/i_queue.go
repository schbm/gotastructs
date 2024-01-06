package queue

import (
	"github.com/schbm/gotastructs/v1/general"
)

type Queue interface {
	general.Iterable
	Append(general.Element)
	Remove() general.Element
	Peek() general.Element
	Size() int
	IsEmpty() bool
}
