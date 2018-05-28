package main

import (
	"fmt"
)

//根据补码，其最大值二进制表示，首位0，其余1，那么，
//const INT_MAX = int(^uint(0) >> 1)
//不能使用上面补码，因为上面的值加1就是一个很小的值，就会影响整个计算结果
const INT_MAX = 10000000

func coinChange(coins []int, amount int) int {
	dp := make([][]int, amount+1)
	for i := 0; i < amount+1; i++ {
		a := make([]int, len(coins))
		dp[i] = a
	}

	// 全部初始化为最大值
	fmt.Println(len(dp))
	fmt.Println(len(dp[0]))
	for i := 0; i < amount+1; i++ {
		for j := 0; j < len(coins); j++ {
			dp[i][j] = INT_MAX
		}
	}

	// 0元需要0张纸币
	for i := 0; i < len(coins); i++ {
		dp[0][i] = 0
	}

	// 纸币价格金额的只需要一张该纸币即可
	for i := 0; i < len(coins); i++ {
		if coins[i] <= amount {
			dp[coins[i]][i] = 1
		}
	}

	// 动态规划使用
	for j := 0; j < amount+1; j++ {
		for i := 0; i < len(coins); i++ {
			// 防止数组越界
			if i-1 >= 0 {
				if (j-coins[i]) >= 0 && dp[j][i-1] > dp[j-coins[i]][i]+1 {
					dp[j][i] = dp[j-coins[i]][i] + 1
				} else {
					dp[j][i] = dp[j][i-1]
				}
			} else {
				if (j - coins[i]) >= 0 {
					dp[j][i] = dp[j-coins[i]][i] + 1
				}
			}

		}
	}
	min := INT_MAX
	for i := 0; i < len(coins); i++ {
		if dp[amount][i] < min {
			min = dp[amount][i]
		}
	}
	if min == INT_MAX {
		min = -1
	}
	return min
}

func main() {
	fmt.Println(INT_MAX)
	coins := []int{2}
	result := coinChange(coins, 3)
	fmt.Println(result)
}
