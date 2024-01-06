package tree

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// Assert interface implementation
var _ Tree[int, string] = &BinarySearchTree[int, string]{}

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
	tree, err := BinaryTreeRemove[K, V](key, bst.Root)
	if err != nil {
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

// --------------------------------------------------------------------------------

// Is a simple recursive implementation of a tree structure. It saves key values in its internal nodes.
// A node is represented by the current tree whereas its children are also subsequent trees.
// A tree (node) without children is a node.
// This is a binary tree since it has a maximum of two children.
type BinaryTree[K constraints.Ordered, V comparable] struct {
	Key   K
	Value V
	Left  *BinaryTree[K, V]
	Right *BinaryTree[K, V]
}

// If the current tree/node has no children it is a leaf
func IsLeaf[K constraints.Ordered, V comparable](tree *BinaryTree[K, V]) bool {
	return (tree.Left == nil && tree.Right == nil)
}

// BinaryTreeSearch performs a binary search on the given binary tree to find the value associated with the specified key.
// It returns the value if found, otherwise it returns an error.
// To implement the search we start at a specific key node in the tree.
// Which node we travel to next is dictated by the result of the comparison of the keys.
// If we hit a leaf, which is a node without children the key has not been found.
func BinaryTreeSearch[K constraints.Ordered, V comparable](key K, tree *BinaryTree[K, V]) (V, error) {
	// if the current tree is nil. The value of the key cannot be found
	if tree == nil {
		var zero V
		return zero, errors.New("")
	}
	// search in the left or in the right subtree.
	// passes nil to the search funtion if there is no child
	if key < tree.Key {
		return BinaryTreeSearch[K, V](key, tree.Left)
	}
	if key > tree.Key {
		return BinaryTreeSearch[K, V](key, tree.Right)
	}
	// if the key is not smaller or bigger than the current key we found our value
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
		// set the new node to the left or right subtree
		// and return with the new current tree
	} else if key <= tree.Key {
		tree.Left = BinaryTreeInsert[K, V](key, value, tree.Left)
	} else {
		tree.Right = BinaryTreeInsert[K, V](key, value, tree.Right)
	}
	// return the current tree/subtree
	return tree
}

// BinaryTreeRemove removes a node with the specified key from the binary search tree.
// It returns the modified binary search tree after the removal operation.
// If the key is not found in the tree, it returns nil.
// There are three cases:
// 1)
// node/tree/subtree has no children
// 2)
// node has one child
// 3)
// node has two childs
// find the internal node which follows inorder traverse
func BinaryTreeRemove[K constraints.Ordered, V comparable](key K, tree *BinaryTree[K, V]) (*BinaryTree[K, V], error) {
	// if tree is nil return error
	if tree == nil {
		return nil, errors.New("tree is null")
	}
	// compare the key and remove on the corresponding child
	if key < tree.Key {
		newTreeLeft, err := BinaryTreeRemove[K, V](key, tree.Left)
		if err != nil {
			return nil, err
		}
		tree.Left = newTreeLeft

	} else if key > tree.Key {
		newTreeRight, err := BinaryTreeRemove[K, V](key, tree.Right)
		if err != nil {
			return nil, err
		}
		tree.Right = newTreeRight
	} else {
		// if the key matches
		// check if a child misses
		// if yes return the existing child which replaces the current tree
		if tree.Left == nil {
			return tree.Right, nil
		}
		if tree.Right == nil {
			return tree.Left, nil
		}

		// if two childs exist
		// fetch the parent of the next inorder child
		parentInorderChild, err := ParentInorderChild[K, V](tree)
		// an error occurs when either the given tree is null
		// or the right child of the tree has no left child
		if err != nil {
			// we set the left child (grandchild) of the right child which is nil
			// to our left child and replace the current tree with the right child
			tree.Right.Left = tree.Left
			return tree.Right, nil
		}
		// if a parent of our inorder child has been found
		// we set its left child which is null to our left child
		// we set the left child of the parent to the right child of the inorder node
		// we set the right child of the inorder node to our right child
		// finally we swap by returning the inorder node
		inorderChild := parentInorderChild.Left
		inorderChild.Left = tree.Left
		parentInorderChild.Left = inorderChild.Right
		inorderChild.Right = tree.Right
		return inorderChild, nil
	}
	return tree, nil
}

// if error is thrown no subtree of the given tree is the parent of a inorder child
func ParentInorderChild[K constraints.Ordered, V comparable](tree *BinaryTree[K, V]) (*BinaryTree[K, V], error) {
	var parent *BinaryTree[K, V]

	// check right node
	// if empty there is no next higher in order child
	if tree == nil || tree.Right == nil {
		return nil, errors.New("tree has no right child")
	}

	if tree.Right.Left != nil {
		return nil, errors.New("right child has no children")
	}

	// set to right child
	parent = tree.Right
	// for each new left child set it as inorder child
	// Traverse down the leftmost path of the right subtree to find the leftmost child's parent
	for parent.Left != nil && parent.Left.Left != nil {
		parent = parent.Left
	}

	return parent, nil
}
