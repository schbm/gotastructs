package gotastructs

import "testing"

func TestGeneralTree(t *testing.T) {
	t.Log("Testing GeneralTree")
	rootTree := NewGeneralTree(NewLinkedList(), nil, &WrappedInt{1})
	rootTree.AddChild(NewGeneralTree(NewLinkedList(), nil, &WrappedInt{2}))
	rootTree.AddChild(NewGeneralTree(NewLinkedList(), nil, &WrappedInt{3}))
	t2 := NewGeneralTree(NewLinkedList(), nil, &WrappedInt{4})
	t2.AddChild(NewGeneralTree(NewLinkedList(), nil, &WrappedInt{10}))
	rootTree.AddChild(t2)

	sl := rootTree.ToSlice()
	for _, v := range sl {
		t.Log(v.String())
	}
}
