package element

import "strings"

type WrappedString string

func (w *WrappedString) Equals(other any) bool {
	v, ok := other.(*WrappedString)
	if !ok {
		return false
	}
	return *w == *v
}

func (w *WrappedString) String() string {
	return string(*w)
}

func (w *WrappedString) Compare(other any) int8 {
	if w.Equals(other) {
		return 0
	}

	v, ok := other.(*WrappedString)
	if !ok {
		return 0
	}

	return int8(strings.Compare(string(*w), string(*v)))
}

func NewString(value string) *WrappedString {
	v := WrappedString(value)
	return &v
}
