package main

import (
	"sort"
)

func maxProduct(nums []int) int {
	max := nums[0]
	maxPro := make([]int, len(nums))
	minPro := make([]int, len(nums))
	maxPro[0] = nums[0]
	minPro[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		a := maxPro[i-1] * nums[i]
		b := minPro[i-1] * nums[i]
		c := nums[i]
		maxPro[i], minPro[i] = findMaxMin(a, b, c)
	}
	for i := 0; i < len(maxPro); i++ {
		if maxPro[i] > max {
			max = maxPro[i]
		}
	}
	return max
}

func findMaxMin(a int, b int, c int) (max int, min int) {
	input := []int{a, b, c}
	sort.Ints(input)
	return input[2], input[0]
}
