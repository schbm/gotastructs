package gotastructs

import (
	list2 "github.com/schbm/gotastructs/list"
	"testing"
)

func TestFilterList(t *testing.T) {
	list := list2.NewArrayList()
	list.Append(NewInt(1))
	list.Append(NewInt(2))
	list.Append(NewInt(3))
	list.Append(NewInt(4))
	list.Append(NewInt(5))
	list.Append(NewInt(6))
	list.Append(NewInt(7))
	list.Append(NewInt(8))
	list.Append(NewInt(9))
	list.Append(NewInt(10))
	FilterList(func(el Element) bool {
		v, ok := el.(*WrappedInt)
		if !ok {
			return true
		}
		if v.value%2 == 0 {
			return true
		}

		return false
	}, list)
	if list.Size() != 5 {
		t.Error("FilterList failed")
	}
	iter := list.Iterator()
	for iter.HasNext() {
		t.Log(iter.Next().(*WrappedInt).value)
	}
}
