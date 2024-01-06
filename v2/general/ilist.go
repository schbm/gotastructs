package general

type List[V comparable] interface {
	Get(V) (V, error)
	Insert(V)
	Remove(V) error
}
