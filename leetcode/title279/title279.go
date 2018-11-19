package main

import (
	"fmt"
)

const INT_MAX = 1000000

func numSquares(n int) int {
	var init []int
	i := 0
	for i*i < n {
		init = append(init, i*i)
		i++
	}
	init = append(init, i*i)

	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		a := make([]int, len(init))
		dp[i] = a
	}

	for i := 0; i < len(init); i++ {
		dp[0][i] = INT_MAX
	}

	for i := 0; i < n+1; i++ {
		for j := 0; j < len(init); j++ {
			if init[j] == i {
				dp[i][j] = 1
			} else {
				dp[i][j] = INT_MAX
			}
		}
	}

	for j := 1; j < n+1; j++ {
		for i := 1; i < len(init); i++ {
			if dp[j][i] != INT_MAX {
				continue
			}

			// 防止数组越界
			if i-1 > 0 {
				if (j-init[i]) > 0 && dp[j][i-1] > dp[j-init[i]][i]+1 {
					dp[j][i] = dp[j-init[i]][i] + 1
				} else {
					dp[j][i] = dp[j][i-1]
				}
			} else if i-1 == 0 {
				if (j - init[i]) > 0 {
					dp[j][i] = dp[j-init[i]][i] + 1
				}
			}

		}
	}

	min := INT_MAX
	for i := 1; i < len(init); i++ {
		if dp[n][i] < min {
			min = dp[n][i]
		}
	}
	fmt.Println(dp[n])
	if min == INT_MAX {
		min = -1
	}
	return min
}

func main() {
	s := numSquares(12)
	fmt.Println(s)
}
