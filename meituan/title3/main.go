package main

import (
	"fmt"
	"io"
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

func (s *Set) Size() int {
	return len(s.m)
}

func main() {
	for {
		var point int
		_, err := fmt.Scanf("%d", &point)
		if err == io.EOF {
			break
		} else {
			pointSets := []*Set{}
			firSet := NewSet()
			for i := 0; i < point-1; i++ {
				var x int
				var y int
				fmt.Scanf("%d", &x)
				fmt.Scanf("%d", &y)

				if x == 1 {
					firSet.Add(y)
					flag := false
					for _, pointSet := range pointSets {
						if pointSet.Has(y) {
							flag = true
							break
						}
					}
					if !flag {
						p := NewSet()
						p.Add(y)
						pointSets = append(pointSets, p)
					}
				} else if y == 1 {
					firSet.Add(x)
					flag := false
					for _, pointSet := range pointSets {
						if pointSet.Has(x) {
							flag = true
							break
						}
					}
					if !flag {
						p := NewSet()
						p.Add(x)
						pointSets = append(pointSets, p)
					}
				} else {

					flag := false
					for _, pointSet := range pointSets {
						if pointSet.Has(x) {
							pointSet.Add(y)
							flag = true
							break
						} else if pointSet.Has(y) {
							pointSet.Add(x)
							flag = true
							break
						}
					}
					if !flag {
						p := NewSet()
						p.Add(x)
						p.Add(y)
						pointSets = append(pointSets, p)
						break
					}
				}
			}

			max := 0
			result := 0
			for _, pointset := range pointSets {
				if pointset.Size() > max {
					max = pointset.Size()
				}
				result = result + pointset.Size()*2
			}

			result = result - max
			fmt.Println(result)

		}
	}

}
