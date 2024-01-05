package tree

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// Binary Search Tree Implementation
// Wraps Binary Tree
// Given nodes u, v and w. u is in the left subtree of v and w is in the right so that key(u) <= key(v) <= key(w)
// Runtimes:
// find: O(log n) binary search
// insert: O(n)
// remove: O(n)
// In the worst case, old entries have to be moved to make space for a new entry
type BinarySearchTree[K constraints.Ordered, V comparable] struct {
	Root *BinaryTree[K, V]
}

func (bst *BinarySearchTree[K, V]) Clear() {
	bst.Root = nil
}

func (bst *BinarySearchTree[K, V]) IsEmpty() bool {
	return bst.Root == nil
}

func (bst *BinarySearchTree[K, V]) Find(key K) (V, error) {
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
	tree := BinaryTreeRemove[K, V](key, bst.Root)
	if tree == nil {
		return errors.New("key not found in tree")
	}
	bst.Root = tree
	return nil
}

func (bst *BinarySearchTree[K, V]) Insert(key K, value V) {
	if bst.IsEmpty() {
		bst.Root = &BinaryTree[K, V]{
			Key:   key,
			Value: value,
		}
	}
	bst.Root = BinaryTreeInsert[K, V](key, value, bst.Root)
}

//--------------------------------------------------------------------------------

type BinaryTree[K constraints.Ordered, V comparable] struct {
	Key   K
	Value V
	Left  *BinaryTree[K, V]
	Right *BinaryTree[K, V]
}

func IsLeaf[K constraints.Ordered, V comparable](tree *BinaryTree[K, V]) bool {
	return (tree.Left == nil && tree.Right == nil)
}

// BinaryTreeSearch performs a binary search on the given binary tree to find the value associated with the specified key.
// It returns the value if found, otherwise it returns an error.
func BinaryTreeSearch[K constraints.Ordered, V comparable](key K, tree *BinaryTree[K, V]) (V, error) {
	if tree == nil {
		var zero V
		return zero, errors.New("")
	}
	if key < tree.Key {
		return BinaryTreeSearch[K, V](key, tree.Left)
	}
	if key > tree.Key {
		return BinaryTreeSearch[K, V](key, tree.Right)
	}
	return tree.Value, nil
}

// BinaryTreeInsert inserts a key-value pair into a binary tree.
// It returns the updated binary tree.
// The key must be of type K, which is required to be ordered.
// The value must be of type V, which is required to be comparable.
// If the tree is nil, a new binary tree node is created with the given key and value.
// If the key is less than or equal to the current node's key, the insertion is performed on the left subtree.
// Otherwise, the insertion is performed on the right subtree.
// The updated binary tree is returned.
func BinaryTreeInsert[K constraints.Ordered, V comparable](key K, value V, tree *BinaryTree[K, V]) *BinaryTree[K, V] {
	if tree == nil {
		return &BinaryTree[K, V]{
			Key:   key,
			Value: value,
		}
	} else if key <= tree.Key {
		tree.Left = BinaryTreeInsert[K, V](key, value, tree.Left)
	} else {
		tree.Right = BinaryTreeInsert[K, V](key, value, tree.Right)
	}
	return tree
}

// BinaryTreeRemove removes a node with the specified key from the binary search tree.
// It returns the modified binary search tree after the removal operation.
// If the key is not found in the tree, it returns nil.
func BinaryTreeRemove[K constraints.Ordered, V comparable](key K, tree *BinaryTree[K, V]) *BinaryTree[K, V] {
	if key == tree.Key {
		return nil
	}

	if key < tree.Key {
		tree.Left = BinaryTreeRemove[K, V](key, tree.Left)
	} else if key > tree.Key {
		tree.Right = BinaryTreeRemove[K, V](key, tree.Right)
	}

	return tree
}
