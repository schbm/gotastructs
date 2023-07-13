package gotastructs

// Slicer is an interface for types that can be converted to a slice of Elements.
type Slicer interface {
	ToSlice() []Element
}
