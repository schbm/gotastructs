package genmap

import (
	"testing"

	"github.com/schbm/gotastructs/v2/general"
)

func TestMultiMap(t *testing.T) {
	// Create a new MultiMap instance
	var mm general.MultiMap[int, string] = NewSliceMultiMap[int, string]()

	// Insert key-value pairs
	mm.Insert(1, "one")
	mm.Insert(1, "two")
	mm.Insert(2, "three")

	// Get all values for a key
	values, err := mm.GetAll(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedValues := []string{"one", "two"}
	if !equalSlice(values, expectedValues) {
		t.Errorf("Expected values %v, got %v", expectedValues, values)
	}

	// Get a specific value for a key
	value, err := mm.GetSpecific(1, "two")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if value != "two" {
		t.Errorf("Expected value 'two', got %v", value)
	}

	// Remove a specific value for a key
	err = mm.RemoveSpecific(1, "two")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify that the value is removed
	_, err = mm.GetSpecific(1, "two")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Remove a specific value for a key
	err = mm.RemoveSpecific(1, "one")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func equalSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
