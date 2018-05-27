package main

import (
	"fmt"
	//	"sort"
)

//根据补码，其最大值二进制表示，首位0，其余1，那么，
const INT_MAX = int(^uint(0) >> 1)

func coinChange(coins []int, amount int) int {
	dp := make([][]int, amount)
	for i := 0; i < amount; i++ {
		a := make([]int, len(coins))
		dp = append(dp, a)
	}

	// 全部初始化为最大值
	for i := 0; i < len(coins); i++ {
		for j := 0; j < coins; j++ {
			dp[i][j] = INT_MAX
		}
	}

	// 0元需要0张纸币
	for i := 0; i < len(coins); i++ {
		dp[0][i] = 0
	}

	// 纸币价格金额的只需要一张该纸币即可
	for i := 0; i < len(coins); i++ {
		dp[coins[i]][i] = 1
	}

	// 动态规划使用
	for i := 0; i < len(amount); i++ {
		for j := 0; j < len(coins); j++ {
			
			if (i-coins[j]) >= 0 && dp[i][j]<dp[i-coins[j]]
			

		}
	}

}
func main() {
	fmt.Println(INT_MAX)
}
