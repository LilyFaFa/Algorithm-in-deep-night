package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	var b []int
	b = make([]int, len(a))
	copy(b, a)
	fmt.Println(a)
	fmt.Println(b)
	c := a
	c[0] = 2
	fmt.Println(a)
	fmt.Println(b)

	test := [...]int{7, 2, 3, 4}
	test2 := []int{7, 2, 3, 4}

	fmt.Println(test)
	fmt.Println(test2)
	modify(test, test2)

	fmt.Println(test)
	fmt.Println(test2)

	m := make([]int, 4, 4)
	fmt.Println(m)
	m = append(m, 1)
	fmt.Println(m)
	fmt.Println(len(m), cap(m))
	m = m[0:2]
	fmt.Println(len(m), cap(m))
	n := m
	fmt.Println(len(n), cap(n))
	n[1] = 3
	fmt.Println(n)
	fmt.Println(m)

}
func modify(test [4]int, test2 []int) {
	test[1] = 1
	test2[1] = 1
}
