package gotastructs

type Queue interface {
	Iterable
	Append(Element)
	Remove() Element
	Peek() Element
	Size() int
	IsEmpty() bool
}
