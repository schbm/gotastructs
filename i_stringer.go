package gotastructs

// Stringer is an interface for types that can be converted to a string.
type Stringer interface {
	// ToString returns a string representation of the object.
	String() string
}
