package gotastructs

type Node interface {
	Id() string
	Value() Element
	Neighbors() []Node
	NeighborCount() int
}

// Graph represents a general graph interface.
type Graph interface {
	// AddNode adds a node to the graph.
	AddNode(node Node)

	// RemoveNode removes a node from the graph.
	RemoveNode(node Node)

	// AddEdge adds an edge between two nodes in the graph.
	AddEdge(node1, node2 Node)

	// RemoveEdge removes an edge between two nodes in the graph.
	RemoveEdge(node1, node2 Node)

	// GetNodes returns all the nodes in the graph.
	GetNodes() []Node

	// HasEdge checks if there is an edge between two nodes in the graph.
	HasEdge(source, target Node) bool
}
