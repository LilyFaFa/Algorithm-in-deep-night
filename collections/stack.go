package collections

import (
	"errors"
)

type Stack struct {
	Items []interface{}
}

func (s *Stack) Pop() (interface{}, error) {
	//current stack is empty
	if len(s.Items) == 0 {
		return nil, errors.New("stack is empty")
	}

	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[0 : len(s.Items)-1]
	return item, nil
}

func (s *Stack) Push(item interface{}) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Size() int {
	return len(s.Items)
}

func (s *Stack) IsEmpty() bool {
	return len(s.Items) == 0
}
