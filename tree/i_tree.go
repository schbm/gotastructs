package tree

import "github.com/schbm/gotastructs"

type Tree interface {
	gotastructs.Element
	gotastructs.Slicer
	//Iterable
	Parent() Tree
	IsRoot() bool
	Value() gotastructs.Element
	Children() []Tree
	AddChild(Tree)
	RemoveChild(Tree) error
	ChildrenCount() int
	//IsEmpty() bool
}
