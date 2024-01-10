package tree

import (
	"errors"

	"github.com/schbm/gotastructs/v2/general"
	"golang.org/x/exp/constraints"
)

// Assert interface implementation
var _ general.Tree[int, string] = &BinarySearchTree[int, string]{}

// Binary Search Tree Implementation
// Wraps Binary Tree to hide certain implementation details
// It implements the tree interface.
// Saves K/V in his internal nodes
// Given nodes u, v and w. u is in the left subtree of v and w is in the right so that key(u) <= key(v) <= key(w)
// Runtimes:
// find: O(log n) binary search
// insert: O(n)
// remove: O(n)
// In the worst case, old entries have to be moved to make space for a new entry
type BinarySearchTree[K constraints.Ordered, V comparable] struct {
	Root  *BinaryTree[K, V]
	Count int
}

func (bst *BinarySearchTree[K, V]) Clear() {
	bst.Root = nil
}

func (bst *BinarySearchTree[K, V]) IsEmpty() bool {
	return bst.Root == nil
}

func (bst *BinarySearchTree[K, V]) Get(key K) (V, error) {
	if bst.IsEmpty() {
		var v V
		return v, errors.New("not found")
	}
	return BinaryTreeSearch[K, V](key, bst.Root)
}

func (bst *BinarySearchTree[K, V]) Remove(key K) error {
	if bst.IsEmpty() {
		return errors.New("cannot remove from an empty tree")
	}
	tree, err := BinaryTreeRemove[K, V](key, bst.Root)
	if err != nil {
		return errors.New("key not found in tree")
	}
	bst.Root = tree
	bst.Count--
	return nil
}

func (bst *BinarySearchTree[K, V]) Insert(key K, value V) {
	defer func() { bst.Count++ }()
	if bst.IsEmpty() {
		bst.Root = &BinaryTree[K, V]{
			Key:   key,
			Value: value,
		}
		return
	}
	bst.Root = BinaryTreeInsert[K, V](key, value, bst.Root)
}

func (bst *BinarySearchTree[K, V]) Size() int {
	return bst.Count
}
