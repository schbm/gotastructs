package list

import (
	"github.com/schbm/gotastructs"
	"testing"
	"time"
)

func TestLinkedList(t *testing.T) {
	t.Log("Testing LinkedList")
	var list List = NewLinkedList()
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
	if f.Equals(&gotastructs.WrappedInt{1}) != true {
		t.Error("first element should be 1")
	}

	// test get last element
	l, err := list.Get(list.Size() - 1)
	if err != nil {
		t.Error(err)
	}
	if l.Equals(&gotastructs.WrappedInt{3}) != true {
		t.Error("last element should be 3")
	}

	//get index out of bounds
	_, err = list.Get(list.Size())
	if err == nil {
		t.Error("should have returned error")
	}

	//inser first
	err = list.Insert(&gotastructs.WrappedInt{0}, 0)
	if err != nil {
		t.Error(err)
	}
	if list.Size() != 4 {
		t.Error("list should have 4 elements")
	}
	f, err = list.Get(0)
	if err != nil {
		t.Error(err)
	}
	if f.Equals(&gotastructs.WrappedInt{0}) != true {
		t.Error("first element should be 0")
	}
	// insert last
	err = list.Insert(&gotastructs.WrappedInt{4}, list.Size()-1)
	if err != nil {
		t.Error(err)
	}
	if list.Size() != 5 {
		t.Error("list should have 5 elements")
	}
	l, err = list.Get(list.Size() - 1)
	if err != nil {
		t.Error(err)
	}
	if l.Equals(&gotastructs.WrappedInt{4}) != true {
		t.Error("last element should be 4")
	}
	// insert out of bounds
	err = list.Insert(&gotastructs.WrappedInt{5}, list.Size())
	if err == nil {
		t.Error("should have returned error")
	}

	// test remove first element
	err = list.Remove(0)
	if err != nil {
		t.Error(err)
	}
	if list.Size() != 4 {
		t.Error("list should have 4 elements")
	}
	f, err = list.Get(0)
	if err != nil {
		t.Error(err)
	}
	if f.Equals(&gotastructs.WrappedInt{1}) != true {
		t.Error("first element should be 1")
	}

	// test remove last element
	err = list.Remove(list.Size() - 1)
	if err != nil {
		t.Error(err)
	}
	if list.Size() != 3 {
		t.Error("list should have 3 elements")
	}
	l, err = list.Get(list.Size() - 1)
	if err != nil {
		t.Error(err)
	}
	if l.Equals(&gotastructs.WrappedInt{3}) != true {
		t.Error("last element should be 3")
	}

	// test remove out of bounds
	err = list.Remove(list.Size())
	if err == nil {
		t.Error("should have returned error")
	}

	// test remove element
	err = list.RemoveElement(&gotastructs.WrappedInt{2})
	if err != nil {
		t.Error(err)
	}
	if list.Size() != 2 {
		t.Error("list should have 2 elements")
	}
	f, err = list.Get(0)
	if err != nil {
		t.Error(err)
	}
	if f.Equals(&gotastructs.WrappedInt{1}) != true {
		t.Error("first element should be 1")
	}
	l, err = list.Get(list.Size() - 1)
	if err != nil {
		t.Error(err)
	}
	if l.Equals(&gotastructs.WrappedInt{3}) != true {
		t.Error("last element should be 3")
	}

	// test remove element not in list
	err = list.RemoveElement(&gotastructs.WrappedInt{2})
	if err == nil {
		t.Error("should have returned error")
	}

	// test remove element from empty list
	list = NewLinkedList()
	err = list.RemoveElement(&gotastructs.WrappedInt{2})
	if err == nil {
		t.Error("should have returned error")
	}

	// insert with index on empty list
	list = NewLinkedList()
	err = list.Insert(&gotastructs.WrappedInt{2}, 0)
	if err == nil {
		t.Error("should have returned error")
	}

}

func TestLinkedListTime(t *testing.T) {
	t.Log("Testing LinkedList Time")
	currT := time.Now()
	var list List = NewLinkedList()
	for i := 0; i < 10000000; i++ {
		list.Append(&gotastructs.WrappedInt{i})
	}
	t.Log("appended 10000000 elements in: ", time.Since(currT))

	//get the last element
	currT = time.Now()
	list.Get(list.Size() - 1)
	t.Log("LinkedList Get the last element: ", time.Since(currT))

	//get middle
	currT = time.Now()
	list.Get(list.Size() / 2)
	t.Log("LinkedList Get the middle element: ", time.Since(currT))

	//insert first
	currT = time.Now()
	list.Insert(&gotastructs.WrappedInt{1}, 0)
	t.Log("LinkedList Insert first element: ", time.Since(currT))

	//insert last
	currT = time.Now()
	list.Insert(&gotastructs.WrappedInt{1}, list.Size()-1)
	t.Log("LinkedList Insert last element: ", time.Since(currT))

	//insert middle
	currT = time.Now()
	list.Insert(&gotastructs.WrappedInt{1}, list.Size()/2)
	t.Log("LinkedList Insert middle element: ", time.Since(currT))

	//remove first
	currT = time.Now()
	list.Remove(0)
	t.Log("LinkedList Remove first element: ", time.Since(currT))

	//remove last
	currT = time.Now()
	list.Remove(list.Size() - 1)
	t.Log("LinkedList Remove last element: ", time.Since(currT))

	//remove middle
	currT = time.Now()
	list.Remove(list.Size() / 2)
	t.Log("LinkedList Remove middle element: ", time.Since(currT))

	//remove high int value
	currT = time.Now()
	list.Remove(list.Size() - 100)
	t.Log("LinkedList Remove high int value element: ", time.Since(currT))

	//remove high int value
	currT = time.Now()
	list.Remove(list.Size() - 200)
	t.Log("LinkedList Remove high int value element: ", time.Since(currT))

	//remove high int value
	currT = time.Now()
	list.Remove(list.Size() - 300)
	t.Log("LinkedList Remove high int value element: ", time.Since(currT))

	//test iterator
	iter := list.Iterator()
	currT = time.Now()
	for iter.HasNext() {
		iter.Next()
	}
	t.Log("LinkedList iteration with iterator: ", time.Since(currT))

	//test iteration with slice
	sl := list.ToSlice()
	currT = time.Now()
	for _, _ = range sl {
	}
	t.Log("LinkedList iteration with slice: ", time.Since(currT))
}
