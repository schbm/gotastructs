package stack

import (
	"errors"
	"github.com/schbm/gotastructs/general"
	"github.com/schbm/gotastructs/list"
)

// pushes and pulles last item from list
type GeneralStack struct {
	list list.List
}

func NewStack(list list.List) *GeneralStack {
	return &GeneralStack{list: list}
}

func (s *GeneralStack) Push(el general.Element) {
	s.list.Append(el)
}

func (s *GeneralStack) Pop() (general.Element, error) {
	if s.list.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	v, err := s.list.Get(s.list.Size() - 1)
	if err != nil {
		return nil, err
	}
	err = s.list.Remove(s.list.Size() - 1)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (s *GeneralStack) Size() int {
	return s.list.Size()
}

func (s *GeneralStack) Peek() (general.Element, error) {
	if s.list.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	v, err := s.list.Get(s.list.Size() - 1)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (s *GeneralStack) IsEmpty() bool {
	return s.list.IsEmpty()
}

type GeneralStackIterator struct {
	list  list.List
	index int
}

func (s *GeneralStackIterator) HasNext() bool {
	return s.index >= 0
}

func (s *GeneralStackIterator) Next() general.Element {
	v, err := s.list.Get(s.index)
	if err != nil {
		return nil
	}
	s.index--
	return v
}

func (s *GeneralStack) Iterator() general.Iterator {
	return &GeneralStackIterator{index: s.list.Size(), list: s.list}
}

func (s *GeneralStack) ToSlice() []general.Element {
	return s.list.ToSlice()
}
