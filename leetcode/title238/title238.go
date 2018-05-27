package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	var resultFron []int
	var resultTail []int
	for i := 0; i < len(nums); i++ {
		resultFron = append(resultFron, 1)
		resultTail = append(resultTail, 1)
	}

	for i := 1; i < len(nums); i++ {
		resultFron[i] = resultFron[i-1] * nums[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		resultTail[i] = resultTail[i+1] * nums[i+1]
	}
	for i := 0; i < len(resultFron); i++ {
		resultFron[i] = resultFron[i] * resultTail[i]
	}
	return resultFron
}
func main() {
	input := []int{0, 0}
	result := productExceptSelf(input)
	fmt.Println(result)
}
