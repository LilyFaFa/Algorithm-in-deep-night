package main

import (
	"fmt"
)

func main() {
	var m int
	var numbers []int
	fmt.Scanf("%d", &m)
	for i := 0; i < m; i++ {
		var in int
		fmt.Scanf("%d", &in)
		numbers = append(numbers, in)
	}
	result := numFactoredBinaryTrees(numbers)
	fmt.Println(result)
	fmt.Println(numbers)
}
