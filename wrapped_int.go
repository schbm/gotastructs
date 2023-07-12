package gotastructs

import "strconv"

type WrappedInt struct {
	value int
}

// implements IComparator
func (w *WrappedInt) Equals(other Comparable) bool {
	v, ok := other.(*WrappedInt)
	if !ok {
		return false
	}
	return w.value == v.value
}

func (w *WrappedInt) String() string {
	return strconv.Itoa(w.value)
}

func (w *WrappedInt) Compare(other Comparable) int8 {
	if w.Equals(other) {
		return 0
	}

	v, ok := other.(*WrappedInt)
	if !ok {
		return -1
	}

	if w.value > v.value {
		return 1
	}
	return -1
}

func NewInt(value int) *WrappedInt {
	return &WrappedInt{
		value: value,
	}
}
