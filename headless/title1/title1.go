package main

import (
	"fmt"
)

/*
输入一个m*n的矩阵
*/
func main() {
	var m, n int
	fmt.Scanf("%d", &m)
	fmt.Scanf("%d", &n)
	var input [][]int
	for i := 0; i < m; i++ {
		var innum []int
		for j := 0; j < n; j++ {
			var num int
			fmt.Scanf("%d", &num)
			innum = append(innum, num)
		}
		input = append(input, innum)
	}
	fmt.Println(input)
}

type Stack struct {
	Items []interface{}
}

func (s *Stack) Pop() (interface{}, error) {
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

func (s *Stack) TopValue(item interface{}) {
	if len(s.Items) == 0 {
		return nil, errors.New("stack is empty")
	}

	item := s.Items[len(s.Items)-1]
	return item, nil
}
