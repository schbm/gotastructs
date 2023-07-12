package gotastructs

type WrappedString string

func (w *WrappedString) Equals(other Comparable) bool {
	v, ok := other.(*WrappedString)
	if !ok {
		return false
	}
	return *w == *v
}

func (w *WrappedString) String() string {
	return string(*w)
}

func (w *WrappedString) Compare(other Comparable) int8 {
	return 0
}

func NewString(value string) *WrappedString {
	v := WrappedString(value)
	return &v //is this allowed? im not sure
}
