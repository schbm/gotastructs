package graph

import "github.com/schbm/gotastructs"

type Node interface {
	Id() string
	Value() gotastructs.Element
	Neighbors() []Node
	NeighborCount() int
}
