package main

const INT_MAX = 10000000

func coinChange(coins []int, amount int) int {
	dp := make([][]int, amount+1)
	for i := 0; i < amount; i++ {
		tmp := make([]int, len(coins))
		dp[i] = tmp
	}

	// 全部初始化为最大值
	for i := 0; i <= amount; i++ {
		for j := 0; j <= len(coins); j++ {
			dp[i][j] = INT_MAX
		}
	}

	// 0 元需要 0 张纸币
	for j := 0; j <= len(coins); j++ {
		dp[0][j] = 0
	}

	// 当前纸币大小的mount
	for i := 0; i < len(coins); i++ {
		if coins[i] <= amount {
			dp[coins[i]][i] = 1
		}
	}

	for i := 0; i <= amount; i++ {
		for j := 0; j <= len(counts); j++ {

		}
	}
}
