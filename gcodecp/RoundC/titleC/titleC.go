package main

import (
	"fmt"
)

func main() {
	a := make(map[int]int)
	a[1] = 2

	fmt.Println(a)
	tmp := []int{1, 2, 4}
	copytt(a, tmp)
}

func copytt(a map[int]int, tmp []int) {
	b := make(map[int]int)
	for k, v := range a {
		b[k] = v
	}
	fmt.Println(b)
	fmt.Println(len(b))
	b[1] = 4

	t := make([]int, len(tmp))
	copy(t, tmp)
	fmt.Println(tmp)
	t = append(t, 4)
	fmt.Println(t)
}
