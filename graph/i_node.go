package graph

import (
	"github.com/schbm/gotastructs/general"
)

type Node interface {
	Id() string
	Value() general.Element
	Neighbors() []Node
	NeighborCount() int
}
