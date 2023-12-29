package tree

import (
	"github.com/schbm/gotastructs/general"
	"github.com/schbm/gotastructs/list"
)

type GeneralTree struct {
	children list.List
	parent   *GeneralTree
	value    general.Element
}

func NewTree(list list.List, parent *GeneralTree, value general.Element) *GeneralTree {
	return &GeneralTree{list, parent, value}
}

func (t *GeneralTree) ChildrenCount() int {
	return t.children.Size()
}

func (t *GeneralTree) Value() general.Element {
	return t.value
}

func (t *GeneralTree) Equals(other any) bool {
	v, ok := other.(*GeneralTree)
	if !ok {
		return false
	}
	return t == v //does this work?
}

func (t *GeneralTree) String() string {
	return t.value.String()
}

func (t *GeneralTree) Parent() Tree {
	return t.parent
}

func (t *GeneralTree) IsRoot() bool {
	return t.parent == nil
}

func (t *GeneralTree) AddChild(child Tree) {
	t.children.Append(child)
}

func (t *GeneralTree) Children() []Tree {
	iterator := t.children.Iterator()

	children := make([]Tree, 0, t.children.Size())
	for iterator.HasNext() {
		children = append(children, iterator.Next().(Tree))
	}

	return children
}

func (t *GeneralTree) RemoveChild(child Tree) error {
	err := t.children.RemoveElement(child)
	if err != nil {
		return err
	}
	return nil
}

func (t *GeneralTree) Compare(other any) int8 {
	if t.Equals(other) {
		return 0
	}

	v, ok := other.(*GeneralTree)
	if !ok {
		return -1
	}

	if t.ChildrenCount() > v.ChildrenCount() {
		return 1
	}
	return 0
}

func (t *GeneralTree) ToSlice() []general.Element {
	childs := t.Children()
	result := make([]general.Element, 0, t.ChildrenCount()+1)
	result = append(result, t.Value())
	for _, el := range childs {
		result = append(result, el.ToSlice()...)
	}
	return result
}
