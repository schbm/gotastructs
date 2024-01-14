package tree

import (
	"testing"

	"github.com/schbm/gotastructs/v2/general"
)

// TODO better tests
func TestBST(t *testing.T) {
	// Create a new Tree instance
	var tree BinarySearchTree[int, string]
	var _ general.Tree[int, string] = &BinarySearchTree[int, string]{}

	// Insert some key-value pairs
	tree.Insert(4, "four")
	tree.Insert(2, "two")
	tree.Insert(3, "three")
	tree.Insert(3, "second three")

	tree.Insert(1, "one")

	tree.Insert(6, "six")
	tree.Insert(5, "five")
	tree.Insert(7, "seven")
	printInorder := tree.Inorder()
	t.Log(printInorder)
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
	printInorder = tree.Inorder()
	t.Log(printInorder)
	// Find removed key
	_, err = tree.Get(2)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	var compareInorder []string = []string{"one", "second three", "three", "four", "five", "six", "seven"}
	result := tree.Inorder()
	for i, output := range result {
		if output != compareInorder[i] {
			t.Errorf("Expected %v, got %v", compareInorder[i], output)
		}
	}

	if tree.Size() != 7 {
		t.Errorf("Expected size 7, got %v", tree.Size())
	}

	if tree.Height() != 3 {
		t.Errorf("Expected height 3, got %v", tree.Height())
	}

}
