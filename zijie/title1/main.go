package main

import (
	"fmt"
)

type Set struct {
	m map[int]bool
}

func NewSet() *Set {
	return &Set{make(map[int]bool)}
}

func (s *Set) Add(items ...int) {
	for _, item := range items {
		s.m[item] = true
	}
}

func (s *Set) Has(item int) bool {
	if _, has := s.m[item]; has {
		return true
	}
	return false
}

func (s *Set) Merge(t *Set) {
	for item := range t.m {
		s.m[item] = true
	}
}

func (s *Set) ShouldMerge(t *Set) bool {

	for item := range t.m {
		for item2 := range s.m {
			if item == item2 {
				return true
			}
		}
	}

	return false
}

func main() {
	n := 0
	fmt.Scan(&n)
	inputs := [][]int{}
	friendSets := []*Set{}

	for i := 0; i < n; i++ {
		input := []int{}
		x := 0
		for {
			fmt.Scan(&x)
			if x != 0 {
				input = append(input, x)
			} else {
				break
			}
		}
		inputs = append(inputs, input)
	}

	for index, input := range inputs {
		flag := false
		for _, friendSet := range friendSets {
			if friendSet.Has(index + 1) {
				friendSet.Add(input...)
				flag = true
				break
			}
			for _, p := range input {
				if friendSet.Has(p) {
					friendSet.Add(input...)
					friendSet.Add(index + 1)
					flag = true
					break
				}
			}
			if flag {
				break
			}

		}
		if !flag {
			f := NewSet()
			f.Add(input...)
			f.Add(index + 1)
			friendSets = append(friendSets, f)
		}
	}

	sum := len(friendSets)
	for index, friendSet1 := range friendSets {
		for i := 0; i < index; i++ {
			friendSet2 := friendSets[i]
			if friendSet2.ShouldMerge(friendSet1) {
				friendSet2.Merge(friendSet1)
				sum--
				break
			}

		}

	}
	fmt.Println(sum)

}
