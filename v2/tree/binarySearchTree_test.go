package tree

import (
	"testing"
)

// TODO better tests
func TestBST(t *testing.T) {
	// Create a new Tree instance
	var tree BinarySearchTree[int, string]
	var _ Tree[int, string] = &BinarySearchTree[int, string]{}

	// Insert some key-value pairs
	tree.Insert(1, "one")
	tree.Insert(2, "two")
	tree.Insert(3, "three")
	tree.Insert(3, "second three")
	tree.Insert(4, "four")
	tree.Insert(5, "five")
	tree.Insert(6, "six")
	tree.Insert(7, "seven")

	// Find existing keys
	value, err := tree.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if value != "one" {
		t.Errorf("Expected value 'one', got %v", value)
	}

	// Find non-existing key
	_, err = tree.Get(8)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Remove a key
	err = tree.Remove(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Find removed key
	_, err = tree.Get(2)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	result := make([]string, 0)
	for _, output := range TreeToInorderSlice[int, string](result, tree.Root) {
		t.Log(output)
	}
}
