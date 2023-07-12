package gotastructs

type GeneralTree struct {
	children List
	parent   *GeneralTree
	value    Element
}

func NewGeneralTree(list List, parent *GeneralTree, value Element) *GeneralTree {
	return &GeneralTree{list, parent, value}
}

func (t *GeneralTree) ChildrenCount() int {
	return t.children.Size()
}

func (t *GeneralTree) Value() Element {
	return t.value
}

func (t *GeneralTree) Equals(other Comparable) bool {
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

func (t *GeneralTree) RemoveChild(child Tree) {
	t.children.RemoveElement(child)
}

func (t *GeneralTree) Compare(other Comparable) int8 {
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

func (t *GeneralTree) ToSlice() []Element {
	childs := t.Children()
	result := make([]Element, 0, t.ChildrenCount()+1)
	result = append(result, t.Value())
	for _, el := range childs {
		result = append(result, el.ToSlice()...)
	}
	return result
}
