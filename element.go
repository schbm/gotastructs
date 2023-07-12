package gotastructs

type Comparable interface {
	Equals(Comparable) bool
	Compare(Comparable) int8
}

type Element interface {
	Comparable
	Stringer
}
