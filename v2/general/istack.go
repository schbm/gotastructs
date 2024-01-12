package general

type Stack[V any] interface {
	Push(V)
	Pop() (V, error)
	Top() (V, error)
	Size() int
}
