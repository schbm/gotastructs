package set

type Set[V comparable] interface {
	Insert(V) error
	Remove(V) error
	Contains(V) bool
	Get() []V
}
