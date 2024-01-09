package general

type List[V comparable] interface {
	Get(int) (V, error)
	IndexOf(V) (int, error)
	Insert(V)
	Remove(V) error
	Contains(V) bool
	Size() int
}
