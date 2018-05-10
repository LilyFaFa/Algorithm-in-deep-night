package main

import (
	"fmt"
)

func main() {
	var m int
	var n int
	fmt.Scanf("%d", &m)
	fmt.Scanf("%d", &n)
	result := uniquePaths(m, n)
	fmt.Println(result)
}
