package set

import "github.com/schbm/gotastructs/v1/general"

type HashSet struct {
	data map[general.Element]general.Element
}

func NewHashSet() *HashSet {
	return &HashSet{ //t
		data: make(map[general.Element]general.Element),
	}
}

func (h *HashSet) Add(el general.Element) {
	h.data[el] = el
}

func (h *HashSet) Remove(el general.Element) {
	if h.IsEmpty() {
		return
	}
	if _, ok := h.data[el]; ok {
		delete(h.data, el)
	}
}

func (h *HashSet) Contains(el general.Element) bool {
	_, ok := h.data[el]
	return ok
}

func (h *HashSet) Size() int {
	return len(h.data)
}

func (h *HashSet) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *HashSet) ToSlice() []general.Element {
	if h.IsEmpty() {
		return nil
	}
	v := make([]general.Element, 0, len(h.data))
	for _, value := range h.data {
		v = append(v, value)
	}
	return v
}

type HashSetIterator struct {
	data  []general.Element
	index int
}

func (l *HashSetIterator) Next() general.Element {
	if !l.HasNext() {
		return nil
	}
	el := l.data[l.index]
	l.index++
	return el
}

func (l *HashSetIterator) HasNext() bool {
	return l.index < len(l.data)
}

func (l *HashSet) Iterator() general.Iterator {
	return &HashSetIterator{
		data:  l.ToSlice(),
		index: 0,
	}
}
