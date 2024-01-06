package set

import "github.com/schbm/gotastructs/v1/general"

type Set interface { //t
	general.Iterable
	general.Slicer
	Add(general.Element)
	Remove(general.Element)
	Contains(general.Element) bool
	Size() int
	IsEmpty() bool
}
