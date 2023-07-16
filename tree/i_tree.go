package tree

import (
	"github.com/schbm/gotastructs/general"
)

type Tree interface {
	general.Element
	general.Slicer
	//Iterable
	Parent() Tree
	IsRoot() bool
	Value() general.Element
	Children() []Tree
	AddChild(Tree)
	RemoveChild(Tree) error
	ChildrenCount() int
	//IsEmpty() bool
}
