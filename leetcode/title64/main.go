package main

import (
	"fmt"
)

func main() {
	a := [][]int{{1, 2}, {5, 6}, {1, 1}}

	result := minPathSum(a)
	fmt.Println(result)
}
