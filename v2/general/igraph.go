package general

type Graph[V any, E any] interface {
	NumVertices() int
	Vertices() []V
	NumEdges() int
	Edges() []E
}
