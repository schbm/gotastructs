package gotastructs

import list2 "github.com/schbm/gotastructs/list"

type GeneralStackIterator struct {
	list  list2.List
	index int
}

func (s *GeneralStackIterator) HasNext() bool {
	return s.index >= 0
}

func (s *GeneralStackIterator) Next() Element {
	v, err := s.list.Get(s.index)
	if err != nil {
		return nil
	}
	s.index--
	return v
}

func (s *GeneralStack) Iterator() Iterator {
	return &GeneralStackIterator{index: s.list.Size(), list: s.list}
}
