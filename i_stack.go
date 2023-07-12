package gotastructs

type Stack interface {
	Iterable
	Push(Element)
	Pop() (Element, error)
	Size() int
	Peek() (Element, error)
	IsEmpty() bool
}
