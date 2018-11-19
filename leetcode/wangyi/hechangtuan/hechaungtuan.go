package main

import (
	"fmt"
)

func main() {
	var count int
	fmt.Scanf("%d", &count)
	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		var n int
		fmt.Scanf("%d", &n)
		numbers = append(numbers, n)
	}
	var k, d int
	fmt.Scanf("%d", &k)
	fmt.Scanf("%d", &d)
	fmt.Println(Max(numbers, k, d))
}
func Max(numbers []int, k int, d int) int {
	max := -100000000

	maxC := make([][]int, k)
	for i := 0; i < k; i++ {
		subArray := make([]int, len(numbers))
		maxC[i] = subArray
	}

	minC := make([][]int, k)
	for i := 0; i < k; i++ {
		subArray := make([]int, len(numbers))
		minC[i] = subArray
	}

	for i := 0; i < len(numbers); i++ {
		maxC[0][i] = numbers[i]
		minC[0][i] = numbers[i]
	}

	for i := 1; i < k; i++ {
		for j := 0; j < len(numbers); j++ {

			for l := j - 1; l > j-d-1; l-- {
				if l < k {
					break
				}
				if maxC[i-1][l]*numbers[j] > minC[i-1][l]*numbers[j] {
					maxC[i][l] = maxC[i-1][l] * numbers[j]
					minC[i][l] = minC[i-1][l] * numbers[j]
				} else {
					maxC[i][l] = minC[i-1][l] * numbers[j]
					minC[i][l] = maxC[i-1][l] * numbers[j]
				}
			}
		}
	}
	for i := k - 1; i < len(numbers); i++ {
		if maxC[k-1][i] > max {
			max = maxC[k-1][i]
		}
	}
	return max
}
