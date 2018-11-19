package main

import (
	"fmt"
	"io"
	//"math"
)

func main() {
	result := []int{}
	for {
		var m int
		var n int
		_, err := fmt.Scanf("%d", &m)
		if err == io.EOF {
			break
		} else {
			fmt.Scanf("%d", &n)
			if n > m {
				n = m
			}
			dp := make([][]int, m+1)
			for i := 0; i < m+1; i++ {
				a := make([]int, n+1)
				dp[i] = a
			}

			for i := 0; i < m+1; i++ {
				dp[i][0] = 1
				dp[i][1] = 1
			}
			for i := 0; i < n+1; i++ {
				dp[0][i] = 1
				dp[1][i] = 1
			}

			for i := 2; i < m+1; i++ {
				if m < n {
					n = m
				}
				k := n
				if i < k {
					k = i
				}
				for j := 2; j <= k; j++ {
					if j < i-j {
						dp[i][j] = dp[i][j-1] + dp[i-j][j]
					} else {
						dp[i][j] = dp[i][j-1] + dp[i-j][i-j]
					}
				}
			}

			result = append(result, dp[m][n])
		}
	}
	for _, r := range result {
		fmt.Println(r)
	}
}
