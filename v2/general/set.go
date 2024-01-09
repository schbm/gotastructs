package general

type Set[V comparable] interface {
	Insert(V) error
	Remove(V) error
	Contains(V) bool
	Size() int
}
