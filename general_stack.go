package gotastructs

import (
	"errors"
	list2 "github.com/schbm/gotastructs/list"
)

// pushes and pulles last item from list
type GeneralStack struct {
	list list2.List
}

func NewGeneralStack(list list2.List) *GeneralStack {
	return &GeneralStack{list: list}
}

func (s *GeneralStack) Push(el Element) {
	s.list.Append(el)
}

func (s *GeneralStack) Pop() (Element, error) {
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

func (s *GeneralStack) Peek() (Element, error) {
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
