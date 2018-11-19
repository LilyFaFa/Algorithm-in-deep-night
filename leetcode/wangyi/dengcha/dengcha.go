package main

import (
	"fmt"
	"sort"
)

func main() {
	var m int
	fmt.Scanf("%d", &m)
	var input []int
	for i := 0; i < m; i++ {
		var in int
		fmt.Scanf("%d", &in)
		input = append(input, in)
	}
	if dengcha(input) {
		fmt.Println("Possible")
	} else {
		fmt.Println("Impossible")
	}
}

func dengcha(input []int) bool {
	if len(input) < 2 {
		return true
	}
	sort.Ints(input)
	cha := input[1] - input[0]
	for i := 1; i < len((input))-1; i++ {
		if input[i+1]-input[i] != cha {
			return false
		}
	}
	return true
}
