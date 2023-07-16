package set

import "github.com/schbm/gotastructs/general"

type Set interface {
	general.Iterable
	general.Slicer
	Add(general.Element)
	Remove(general.Element)
	Contains(general.Element) bool
	Size() int
	IsEmpty() bool
}
