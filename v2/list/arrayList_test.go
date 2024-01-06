package list

import (
	"testing"
	"time"
)

func TestArrayList(t *testing.T) {
	t.Log("Testing ArrayList")

	tests := []struct {
		name string
		size int
	}{
		{
			name: "100",
			size: 100,
		},
		{
			name: "1000",
			size: 1000,
		},
		{
			name: "10000",
			size: 1000,
		},
	}

	for _, test := range tests {
		data := make([]int, test.size)

		list := NewArrayList[int]()

		for i := 0; i < test.size; i++ {
			data[i] = i
		}

		for _, number := range data {
			list.Append(number)
		}

		if list.Size() != test.size {
			t.Error("incorrect size")
		}

		val, err := list.Get(test.size - 1)
		if err != nil || val != test.size-1 {
			t.Errorf("max number cannot be found")
		}
	}
}

func MeasureTime(name string, t *testing.T) func() {
	start := time.Now()
	return func() {
		t.Logf("Time taken by %s is %v\n", name, time.Since(start))
	}
}
