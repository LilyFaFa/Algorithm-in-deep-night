package main

import (
	"fmt"
)

func main() {
	a := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	b := [][]int{{1, 0}}
	result := uniquePathsWithObstacles(a)
	fmt.Println(result)

}
