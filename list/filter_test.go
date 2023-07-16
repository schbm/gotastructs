package list

import (
	"github.com/schbm/gotastructs/element"
	"github.com/schbm/gotastructs/general"
	"testing"
)

func TestFilterList(t *testing.T) {
	list := NewArrayList()
	list.Append(element.NewInt(1))
	list.Append(element.NewInt(2))
	list.Append(element.NewInt(3))
	list.Append(element.NewInt(4))
	list.Append(element.NewInt(5))
	list.Append(element.NewInt(6))
	list.Append(element.NewInt(7))
	list.Append(element.NewInt(8))
	list.Append(element.NewInt(9))
	list.Append(element.NewInt(10))
	FilterList(func(el general.Element) bool {
		v, ok := el.(*element.WrappedInt)
		if !ok {
			return true
		}
		if v.Value()%2 == 0 {
			return true
		}

		return false
	}, list)
	if list.Size() != 5 {
		t.Error("FilterList failed")
	}
	iter := list.Iterator()
	for iter.HasNext() {
		t.Log(iter.Next().(*element.WrappedInt).String())
	}
}
