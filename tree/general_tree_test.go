package tree

import (
	"github.com/schbm/gotastructs/element"
	"github.com/schbm/gotastructs/list"
	"testing"
)

func TestGeneralTree(t *testing.T) {
	t.Log("Testing GeneralTree")
	rootTree := NewTree(list.NewLinkedList(false), nil, element.NewInt(1))
	rootTree.AddChild(NewTree(list.NewLinkedList(false), nil, element.NewInt(2)))
	rootTree.AddChild(NewTree(list.NewLinkedList(false), nil, element.NewInt(3)))
	t2 := NewTree(list.NewLinkedList(false), nil, element.NewInt(4))
	t2.AddChild(NewTree(list.NewLinkedList(false), nil, element.NewInt(10)))
	rootTree.AddChild(t2)

	sl := rootTree.ToSlice()
	for _, v := range sl {
		t.Log(v.String())
	}
}
