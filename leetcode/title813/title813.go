package main

import (
	"fmt"
)

func largestSumOfAverages(A []int, K int) float64 {
	dp := make([][]float64, K+1)
	for i := 0; i < K+1; i++ {
		tmp := make([]float64, len(A))
		dp[i] = tmp
	}
	//丢弃k=0，初始化k==1
	dp[1][0] = float64(A[0])
	for i := 1; i < len(A); i++ {
		dp[1][i] = (dp[1][i-1]*float64(i) + float64(A[i])) / float64(i+1)
	}

	for k := 2; k < K+1; k++ {
		for i := 0; i < len(A); i++ {
			if k == 1 && i == 0 {
				continue
			}
			var max float64

			for j := 0; j < i; j++ {
				var sum float64
				for m := j + 1; m <= i; m++ {
					sum += float64(A[m])
				}
				sum = sum / float64(i-j)
				if max < sum+dp[k-1][j] {
					max = sum + dp[k-1][j]
				}
			}
			dp[k][i] = max
		}
	}

	return dp[K][len(A)-1]
}

func main() {
	example := []int{4, 1, 7, 5, 6, 2, 3}
	result := largestSumOfAverages(example, 4)
	fmt.Println(result)
}
