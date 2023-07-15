package list

import (
	"github.com/schbm/gotastructs"
	"testing"
	"time"
)

func TestWrappedString(t *testing.T) {
	t.Log("Testing WrappedString")
	ws1 := gotastructs.WrappedString("test")
	ws2 := gotastructs.WrappedString("test")
	if !ws1.Equals(&ws2) {
		t.Error("should be equal")
	}
	ws3 := gotastructs.WrappedString("test2")
	if ws1.Equals(&ws3) {
		t.Error("should not be equal")
	}
	t.Log(ws3)
	list := NewLinkedList()
	list.Append(&ws1)
	list.Append(&ws2)
	list.Append(&ws3)
	if list.Size() != 3 {
		t.Error("list should have 3 elements")
	}
	if list.IsEmpty() {
		t.Error("list should not be empty")
	}

	// test get first element
	f, err := list.Get(0)
	if err != nil {
		t.Error(err)
	}
	if !f.Equals(&ws1) {
		t.Error("first element should be ws1")
	}
}

func TestDoublyLinkedList(t *testing.T) {
	t.Log("Testing DoublyLinkedList")
	var list List = NewDoublyLinkedList()
	if list.Size() != 0 {
		t.Error("list should be empty")
	}
	if !list.IsEmpty() {
		t.Error("list should be empty")
	}

	// test append 3 values
	list.Append(&gotastructs.WrappedInt{1})
	list.Append(&gotastructs.WrappedInt{2})
	list.Append(&gotastructs.WrappedInt{3})
	if list.Size() != 3 {
		t.Error("list should have 3 elements")
	}
	if list.IsEmpty() {
		t.Error("list should not be empty")
	}

	// test get first element
	f, err := list.Get(0)
	if err != nil {
		t.Error(err)
	}
	if !f.Equals(&gotastructs.WrappedInt{1}) {
		t.Error("first element should be 1")
	}

	// get middle element
	m, err := list.Get(1)
	if err != nil {
		t.Error(err)
	}
	if !m.Equals(&gotastructs.WrappedInt{2}) {
		t.Error("middle element should be 2")
	}

	// get last element
	l, err := list.Get(2)
	if err != nil {
		t.Error(err)
	}

	if !l.Equals(&gotastructs.WrappedInt{3}) {
		t.Error("last element should be 3")
	}

	// test insert first middle last
	list.Insert(&gotastructs.WrappedInt{0}, 0)
	list.Insert(&gotastructs.WrappedInt{4}, 1)
	list.Insert(&gotastructs.WrappedInt{5}, 2)

	v, err := list.Get(1)
	if err != nil {
		t.Error(err)
	}
	if !v.Equals(&gotastructs.WrappedInt{4}) {
		t.Error("middle element should be 4")
	}

	//test remove first middle last
	err = list.Remove(0)
	if err != nil {
		t.Error(err)
	}
	err = list.Remove(1)
	if err != nil {
		t.Error(err)
	}
	err = list.Remove(2)
	if err != nil {
		t.Error(err)
	}

	if list.Size() != 3 {
		t.Error("list should have 3 elements")
	}

	//test out of bounds
	err = list.Remove(3)
	if err == nil {
		t.Error("should have gotten an error")
	}

	// test out of bound insert
	err = list.Insert(&gotastructs.WrappedInt{0}, 4)
	if err == nil {
		t.Error("should have gotten an error")
	}

}

func TestDoublyLinkedListTime(t *testing.T) {
	t.Log("Testing DoublyLinkedList Time")
	currT := time.Now()
	var list List = NewDoublyLinkedList()
	for i := 0; i < 10000000; i++ {
		list.Append(&gotastructs.WrappedInt{i})
	}
	t.Log("appended 10000000 elements in", time.Since(currT))
	//get the last element
	currT = time.Now()
	list.Get(list.Size() - 1)
	t.Log("DoublyLinkedList Get the last element: ", time.Since(currT))

	//get middle
	currT = time.Now()
	list.Get(list.Size() / 2)
	t.Log("DoublyLinkedList Get the middle element: ", time.Since(currT))

	//insert first
	currT = time.Now()
	list.Insert(&gotastructs.WrappedInt{1}, 0)
	t.Log("DoublyLinkedList Insert first element: ", time.Since(currT))

	//insert last
	currT = time.Now()
	list.Insert(&gotastructs.WrappedInt{1}, list.Size()-1)
	t.Log("DoublyLinkedList Insert last element: ", time.Since(currT))

	//insert middle
	currT = time.Now()
	list.Insert(&gotastructs.WrappedInt{1}, list.Size()/2)
	t.Log("DoublyLinkedList Insert middle element: ", time.Since(currT))

	//remove first
	currT = time.Now()
	list.Remove(0)
	t.Log("DoublyLinkedList Remove first element: ", time.Since(currT))

	//remove last
	currT = time.Now()
	list.Remove(list.Size() - 1)
	t.Log("DoublyLinkedList Remove last element: ", time.Since(currT))

	//remove middle
	currT = time.Now()
	list.Remove(list.Size() / 2)
	t.Log("DoublyLinkedList Remove middle element: ", time.Since(currT))

	//remove high int value
	currT = time.Now()
	list.Remove(list.Size() - 100)
	t.Log("DoublyLinkedList Remove high int value element: ", time.Since(currT))

	//remove high int value
	currT = time.Now()
	list.Remove(list.Size() - 200)
	t.Log("DoublyLinkedList Remove high int value element: ", time.Since(currT))

	//remove high int value
	currT = time.Now()
	list.Remove(list.Size() - 300)
	t.Log("DoublyLinkedList Remove high int value element: ", time.Since(currT))

	//test iterator
	iter := list.Iterator()
	currT = time.Now()
	for iter.HasNext() {
		iter.Next()
	}
	t.Log("DoublyLinkedList iteration with iterator: ", time.Since(currT))

	//test iteration with slice
	sl := list.ToSlice()
	currT = time.Now()
	for _, _ = range sl {
	}
	t.Log("DoublyLinkedList iteration with slice: ", time.Since(currT))
}
