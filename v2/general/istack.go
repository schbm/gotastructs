package general

type Stack[V any] interface {
	Push(V)
	Pop() V
	Top() V
	Size() int
}
