package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scanf("%s", &input)
	fmt.Println(jiaocuo(input))
}
func jiaocuo(input string) int {
	count := make([]int, len(input))
	count[0] = 1
	for i := 1; i < len(input); i++ {
		if input[i] != input[i-1] {
			count[i] = count[i-1] + 1
		} else {
			count[i] = 1
		}
	}
	var max int
	for _, i := range count {
		if i > max {
			max = i
		}
	}
	return max
}
