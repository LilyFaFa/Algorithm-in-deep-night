package main

import (
	"fmt"
	"sort"
)

func main() {

	var num int
	fmt.Scanf("%d", &num)
	m := make(map[int]int)
	s := []int{}
	for i := 0; i < num; i++ {
		var a int
		var b int
		fmt.Scanf("%d", &a)
		fmt.Scanf("%d", &b)
		if _, ok := m[a]; ok {
			m[a] = m[a] + b
		} else {
			s = append(s, a)
			m[a] = b
		}
	}
	sort.Ints(s)
	for _, i := range s {
		fmt.Print(i)
		fmt.Print(" ")
		fmt.Println(m[i])
	}
}
