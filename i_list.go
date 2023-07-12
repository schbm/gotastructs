package gotastructs

type ListError interface {
	error
}

type List interface {
	Iterable
	Slicer
	Append(Element)
	Insert(Element, int) ListError
	Remove(int) ListError
	RemoveElement(Element) ListError
	IndexOf(Element) (int, ListError)
	Contains(Element) bool
	Get(int) (Element, ListError)
	IsEmpty() bool
	Size() int
}
