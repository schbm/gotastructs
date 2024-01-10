package list

import (
	"testing"

	"github.com/schbm/gotastructs/v2/general"
)

func TestListImplementation(t *testing.T) {
	size := 100

	tests := []struct {
		Name string
		List general.List[int]
		Size int
	}{
		{
			Name: "array",
			List: NewArrayList[int](),
			Size: size,
		},
		{
			Name: "linked",
			List: NewLinkedList[int](true),
			Size: size,
		},
		{
			Name: "doubly linked",
			List: NewDoublyLinkedList[int](),
			Size: size,
		},
	}

	for _, test := range tests {
		t.Log("starting with", test.Name)
		list := test.List

		//test zero value on empty list
		if list.Size() != 0 {
			t.Error("initial list size is not zero")
		}
		if list.Contains(0) {
			t.Error("empty list contains zero")
		}
		if _, err := list.Get(0); err == nil {
			t.Error("could get something even if list is empty")
		}
		if err := list.Remove(0); err == nil {
			t.Error("could remove something even if list is empty")
		}
		//-------------------------------

		for i := 0; i < test.Size; i++ {
			list.Insert(i + 1)
		}

		if list.Size() != test.Size {
			t.Error("error inserting into list, size does not match")
		}

		for i := 0; i < test.Size; i++ {
			_, err := list.Get(i + 1)
			if err != nil {
				t.Error("error getting item", i+1)
				continue
			}

			val, err := list.GetFrom(i)
			if err != nil {
				t.Error("error getting itemby index", i)
				continue
			}

			if val != i+1 {
				t.Error("item does not match expected value", i, val)
			}

			if !list.Contains(i + 1) {
				t.Error("list does not contain item")
			}
		}

		// remove highest
		if list.Remove(test.Size) != nil {
			t.Error("could not remove highest item")
		}
		if list.Contains(test.Size) {
			t.Error("highest item still exists")
		}
		if list.Size() != test.Size-1 {
			t.Error("size was not decreased")
		}

		// remove lowest
		err := list.Remove(0)
		if err == nil {
			t.Error("could remove item 0 even if not in list")
		}
		err = list.Remove(1)
		if err != nil {
			t.Error("cannot remove first element")
		}
		// index errors
		if _, err := list.GetFrom(-1); err == nil {
			t.Error("could get from -1 index")
		}
		if err := list.RemoveFrom(-1); err == nil {
			t.Error("could remove from -1 index")
		}
		if err := list.InsertTo(100, -1); err == nil {
			t.Error("could insert into -1 index")
		}
		if _, err := list.GetFrom(list.Size()); err == nil {
			t.Error("could get from list.Size() index")
		}
		if err := list.RemoveFrom(list.Size()); err == nil {
			t.Error("could remove from list.Size() index")
		}
		if err := list.InsertTo(100, list.Size()); err == nil {
			t.Error("could insert into list.Size() index")
		}
	}
}
