package general

type Iterator interface {
	HasNext() bool
	Next() Element
}

type Iterable interface {
	Iterator() Iterator
}
