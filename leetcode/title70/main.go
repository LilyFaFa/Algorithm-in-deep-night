package main

import (
	"fmt"
)

func main() {
	c := climbStairs(8)
	fmt.Println(c)

	// just for test
	m := make(map[int]int)
	testmap(m, 2)
	fmt.Println(m[2])
}

func testmap(m map[int]int, i int) {
	m[i] = i
}
