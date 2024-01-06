package list

import (
	"strconv"
	"testing"
	"time"

	"github.com/schbm/gotastructs/v1/element"
)

func TestArrayList(t *testing.T) {
	t.Log("Testing ArrayList")
	var list List = NewArrayList()
	if list.Size() != 0 {
		t.Error("list should be empty")
	}
	if !list.IsEmpty() {
		t.Error("list should be empty")
	}

	// test append 3 values
	list.Append(element.NewInt(1))
	list.Append(element.NewInt(2))
	list.Append(element.NewInt(3))
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
	if !f.Equals(element.NewInt(1)) {
		t.Error("first element should be 1")
	}

	// get middle element
	m, err := list.Get(1)
	if err != nil {
		t.Error(err)
	}
	if !m.Equals(element.NewInt(2)) {
		t.Error("middle element should be 2")
	}

	// get last element
	l, err := list.Get(2)
	if err != nil {
		t.Error(err)
	}

	if !l.Equals(element.NewInt(3)) {
		t.Error("last element should be 3")
	}

	// test insert first middle last
	err = list.Insert(element.NewInt(0), 0)
	if err != nil {
		t.Error(err)
	}
	err = list.Insert(element.NewInt(4), 1)
	if err != nil {
		t.Error(err)
	}
	err = list.Insert(element.NewInt(5), 2)
	if err != nil {
		t.Error(err)
	}

	v, err := list.Get(1)
	if err != nil {
		t.Error(err)
	}
	if !v.Equals(element.NewInt(4)) {
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
	err = list.Insert(element.NewInt(0), 4)
	if err == nil {
		t.Error("should have gotten an error")
	}
}

func TestArrayListTime(t *testing.T) {
	t.Log("Testing ArrayList Time")
	currT := time.Now()
	var list List = NewArrayList()
	for i := 0; i < 10000000; i++ {
		list.Append(element.NewInt(i))
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
	_ = list.Insert(element.NewInt(1), 0)
	t.Log("ArrayList Insert first element: ", time.Since(currT))

	//insert last
	currT = time.Now()
	_ = list.Insert(element.NewInt(1), list.Size()-1)
	t.Log("ArrayList Insert last element: ", time.Since(currT))

	//insert middle
	currT = time.Now()
	_ = list.Insert(element.NewInt(1), list.Size()/2)
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

	//test iterator
	iter := list.Iterator()
	currT = time.Now()
	for iter.HasNext() {
		iter.Next()
	}
	t.Log("ArrayList iteration with iterator: ", time.Since(currT))

	//test iteration with slice
	sl := list.ToSlice()
	currT = time.Now()
	for _, _ = range sl {
	}
	t.Log("ArrayList iteration with slice: ", time.Since(currT))
}

func TestNormalSlice(t *testing.T) {
	t.Log("Testing normal array Time")
	list := make([]int, 0, 200)

	currT := time.Now()
	for i := 0; i < 10000000; i++ {
		list = append(list, i)
	}
	t.Log("appended 10000000 elements in: ", time.Since(currT))

	currT = time.Now()
	for _, _ = range list {
	}
	t.Log("array iteration with iterator: ", time.Since(currT))
}

func BenchmarkALAppendInt(b *testing.B) {
	list := NewArrayList()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Append(element.NewInt(i))
	}
}

func BenchmarkALAppendString(b *testing.B) {
	list := NewArrayList()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Append(element.NewString(strconv.Itoa(i)))
	}
}

func BenchmarkALInsert(b *testing.B) {
	list := NewArrayList()

	// Fill the list with some initial elements
	for i := 0; i < b.N; i++ {
		list.Append(element.NewInt(i))
	} //BenchmarkInsert-8         130466            125376 ns/op

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Insert(element.NewInt(i), i%list.Size())
	}
}
