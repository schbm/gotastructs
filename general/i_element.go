package general

type Equaler interface {
	// Equals compares this Comparables to another Comparables.
	// Returns true if they are equal, false otherwise.
	// Two Comparables are equal if they have the same type and value.
	// For custom types that implement Comparable, this method should be
	// implemented in a way that makes sense for the type.
	// As this probably will be called a lot of times, it should be
	// implemented efficiently.
	Equals(any) bool
}

// Comparer describes a type that can be compared to another type of the
// same specific type.
type Comparer interface {
	// Compare compares this Comparable to another type.
	// Returns 0 if they are equal, a positive number if this Comparable is
	// greater than the other, and a negative number if this
	// Comparable is less than the other Comparable.
	// For custom types that implement Comparable, this method should be
	// implemented in a way that makes sense for the type.
	// As this probably will be called a lot of times, it should be
	// implemented efficiently.
	// If two comparables are not of the same specific type, the return value
	// is undefined. It is recommended to return 0 in this case.
	Compare(any) int8
}

// Element describes a type that can be stored in a data structure.
// It is currently a combination of Comparable and Stringer.
type Element interface {
	Equaler
	Comparer
	Stringer
}
