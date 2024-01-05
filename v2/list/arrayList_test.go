package list

import (
	"testing"
	"time"
)

func TestArrayList(t *testing.T) {
	t.Log("Testing ArrayList")
	list := NewArrayList[int]()
	if list.Size() != 0 {
		t.Error("list should be empty")
	}
	if !list.IsEmpty() {
		t.Error("list should be empty")
	}

	// test append 3 values
	list.Append(1)
	list.Append(2)
	list.Append(3)
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
	if f != 1 {
		t.Error("first element should be 1")
	}

	// get middle element
	m, err := list.Get(1)
	if err != nil {
		t.Error(err)
	}
	if m != 2 {
		t.Error("middle element should be 2")
	}

	// get last element
	l, err := list.Get(2)
	if err != nil {
		t.Error(err)
	}

	if l != 3 {
		t.Error("last element should be 3")
	}

	// test insert first middle last
	err = list.Insert(0, 0)
	if err != nil {
		t.Error(err)
	}
	err = list.Insert(4, 1)
	if err != nil {
		t.Error(err)
	}
	err = list.Insert(5, 2)
	if err != nil {
		t.Error(err)
	}

	v, err := list.Get(1)
	if err != nil {
		t.Error(err)
	}
	if v != 4 {
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
	err = list.Insert(0, 4)
	if err == nil {
		t.Error("should have gotten an error")
	}
}

func TestArrayListTime(t *testing.T) {
	t.Log("Testing ArrayList Time")
	currT := time.Now()
	list := NewArrayList[int]()
	for i := 0; i < 10000000; i++ {
		list.Append(i)
	}
	t.Log("appended 10000000 elements in: ", time.Since(currT))

	//get the last element
	currT = time.Now()
	_, _ = list.Get(list.Size() - 1)
	t.Log("ArrayList Get the last element: ", time.Since(currT))

	//get middle
	currT = time.Now()
	_, _ = list.Get(list.Size() / 2)
	t.Log("ArrayList Get the middle element: ", time.Since(currT))

	//insert first
	currT = time.Now()
	_ = list.Insert(1, 0)
	t.Log("ArrayList Insert first element: ", time.Since(currT))

	//insert last
	currT = time.Now()
	_ = list.Insert(1, list.Size()-1)
	t.Log("ArrayList Insert last element: ", time.Since(currT))

	//insert middle
	currT = time.Now()
	_ = list.Insert(1, list.Size()/2)
	t.Log("ArrayList Insert middle element: ", time.Since(currT))

	//remove first
	currT = time.Now()
	_ = list.Remove(0)
	t.Log("ArrayList Remove first element: ", time.Since(currT))

	//remove last
	currT = time.Now()
	_ = list.Remove(list.Size() - 1)
	t.Log("ArrayList Remove last element: ", time.Since(currT))

	//remove middle
	currT = time.Now()
	_ = list.Remove(list.Size() / 2)
	t.Log("ArrayList Remove middle element: ", time.Since(currT))

	//remove high int value
	currT = time.Now()
	_ = list.Remove(list.Size() - 100)
	t.Log("ArrayList Remove high int value element: ", time.Since(currT))

	//remove high int value
	currT = time.Now()
	_ = list.Remove(list.Size() - 200)
	t.Log("ArrayList Remove high int value element: ", time.Since(currT))

	//remove high int value
	currT = time.Now()
	_ = list.Remove(list.Size() - 300)
	t.Log("ArrayList Remove high int value element: ", time.Since(currT))
}
