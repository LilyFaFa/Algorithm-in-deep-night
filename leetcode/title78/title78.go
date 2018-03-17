package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	result := [][]int{}
	member := []int{}
	result = append(result, []int{})
	backtrack(&result, member, nums, 0)
	return result
}

func backtrack(result *([][]int), member []int, nums []int, current int) {

	for i := current; i < len(nums); i++ {
		insert := []int{}
		for _, m := range member {
			insert = append(insert, m)
		}
		insert = append(insert, nums[i])
		if len(insert) != 0 {
			(*result) = append((*result), insert)
		}
		backtrack(result, insert, nums, i+1)
	}
}

func main() {
	test := []int{9, 0, 3, 5, 8, 9, 10, 12}
	result := subsets(test)
	fmt.Println(result)
}
