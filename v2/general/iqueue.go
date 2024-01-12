package general

type Queue[V any] interface {
	Enqueue(V)
	Dequeue() V
	First() V
	Size() int
}
