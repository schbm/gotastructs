package tree

import (
	"golang.org/x/exp/constraints"
)

//var _ general.SearchTree[int, string] = &AVLTree[int, string]{}

type AVLTree[K constraints.Ordered, V comparable] struct {
	bst BinarySearchTree[K, V]
}
