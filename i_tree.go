package gotastructs

type Tree interface {
	Element
	Slicer
	//Iterable
	Parent() Tree
	IsRoot() bool
	Value() Element
	Children() []Tree
	AddChild(Tree)
	RemoveChild(Tree) error
	ChildrenCount() int
	//IsEmpty() bool
}
