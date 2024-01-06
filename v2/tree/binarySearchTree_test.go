package tree

import (
	"testing"
)

func TestBST(t *testing.T) {
	// Create a new Tree instance
	var tree BinarySearchTree[int, string]
	var _ Tree[int, string] = &BinarySearchTree[int, string]{}

	// Insert some key-value pairs
	tree.Insert(1, "one")
	tree.Insert(2, "two")
	tree.Insert(3, "three")

	// Find existing keys
	value, err := tree.Find(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if value != "one" {
		t.Errorf("Expected value 'one', got %v", value)
	}

	// Find non-existing key
	_, err = tree.Find(4)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Remove a key
	err = tree.Remove(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Find removed key
	_, err = tree.Find(2)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
