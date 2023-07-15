package tree

import (
	"github.com/schbm/gotastructs/element"
	"github.com/schbm/gotastructs/list"
	"testing"
)

func TestGeneralTree(t *testing.T) {
	t.Log("Testing GeneralTree")
	rootTree := NewGeneralTree(list.NewLinkedList(), nil, &element.WrappedInt{1})
	rootTree.AddChild(NewGeneralTree(list.NewLinkedList(), nil, &element.WrappedInt{2}))
	rootTree.AddChild(NewGeneralTree(list.NewLinkedList(), nil, &element.WrappedInt{3}))
	t2 := NewGeneralTree(list.NewLinkedList(), nil, &element.WrappedInt{4})
	t2.AddChild(NewGeneralTree(list.NewLinkedList(), nil, &element.WrappedInt{10}))
	rootTree.AddChild(t2)

	sl := rootTree.ToSlice()
	for _, v := range sl {
		t.Log(v.String())
	}
}
