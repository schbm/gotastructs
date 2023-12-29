package graph

type Graph[V comparable, E comparable] interface {
	NumVertices() int
	NumEdges() int
	Vertices() []V
	Edges() []E
}
