package main

import (
	"fmt"
)

func splitArray(nums []int, m int) int {
	dp := [][]int{}
	for i := 0; i < m; i++ {
		d := make([]int, len(nums))
		dp = append(dp, d)
	}

	dp[0][0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[0][i] = dp[0][i-1] + nums[i]
	}

	for i := 1; i < m; i++ {
		for j := i; j < len(nums); j++ {
			min := nums[j]
			if dp[i-1][j-1] > nums[j] {
				min = dp[i-1][j-1]
			}
			sum := nums[j]
			for k := j - 2; k >= i-1; k-- {
				sum += nums[k+1]
				if dp[i-1][k] < sum {
					if min > sum {
						min = sum
					}
				} else {
					if min > dp[i-1][k] {
						min = dp[i-1][k]
					}
				}
			}
			dp[i][j] = min
		}
		fmt.Println(dp[i])
	}
	return dp[m-1][len(nums)-1]
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 4, 7, 8, 9}
	r := splitArray(nums, 5)
	fmt.Println(r)
}
