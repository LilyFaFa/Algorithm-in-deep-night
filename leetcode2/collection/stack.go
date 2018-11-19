package collection

import (
	"container/list"
)

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	l := new(list.List)
	return &Stack{l}
}

func (s *Stack) Pop() interface{} {
	if elem := s.list.Back(); elem != nil {
		s.list.Remove(elem)
		return elem.Value
	}
	return nil
}

func (s *Stack) Push(value interface{}) {
	s.list.PushBack(value)
}
