package general

// List represents a unordered collection of elements.
// Duplicates are allowed
type List[V comparable] interface {
	// Gets the first entry which matches the element
	Get(V) (V, error)
	// Gets the entry which matches the index
	GetFrom(int) (V, error)
	// Gets the index of the first element that matches
	IndexOf(V) (int, error)
	// Insert an element
	Insert(V)
	InsertTo(V, int) error // If the index is out of bounds, it returns an error.
	// Removes an element
	Remove(V) error
	RemoveFrom(int) error
	// Checks if the element is in the list
	Contains(V) bool
	// Returns the size of the list
	Size() int
}
