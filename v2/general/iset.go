package general

// Set is basically a list but without duplicates
type Set[V comparable] interface {
	// Insert an element
	Insert(V)
	// Removes an element
	Remove(V) error
	// Checks if the element is in the list
	Contains(V) bool
	// Returns the size of the list
	Size() int
}
