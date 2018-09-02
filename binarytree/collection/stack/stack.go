package stack

import (
	"errors"
)

type Stack struct {
	Item []interface{}
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Pop() (interface{}, error) {
	if len(s.Item) == 0 {
		return nil, errors.New("Empty stack")
	}
	item := s.Item[len(s.Item)-1]
	s.Item = s.Item[:len(s.Item)-1]
	return item, nil
}

func (s *Stack) Push(i interface{}) error {
	s.Item = append(s.Item, i)
	return nil
}
