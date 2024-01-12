package general

type Slicer[V any] interface {
	Slice() []V
}
