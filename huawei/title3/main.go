package main

import (
	"fmt"
)

func main() {
	var l, s, t, m int
	fmt.Scanf("%d", &l)
	fmt.Scanf("%d", &s)
	fmt.Scanf("%d", &t)
	fmt.Scanf("%d", &m)

	dp := make(map[int]int)
	var inputs []int
	for i := 0; i < m; i++ {
		var num int
		fmt.Scanf("%d", &num)
		if num <= t && num >= s {
			dp[num] = 1
		} else {
			dp[num] = 10000000000000
		}
		inputs = append(inputs, num)
	}

	inputs = append(inputs, l)
	dp[l] = 100000000000

	for i := 1; i < len(inputs); i++ {
		var j, k int
		min := 10000000000

		if inputs[i] <= t {
			continue
		}

		j = inputs[i] - t
		k = inputs[i] - s

		for p := j; p <= k; p++ {
			if num, ok := dp[p]; ok && num < min {
				min = num
			}
		}

		if dp[inputs[i]] > min+1 {
			dp[inputs[i]] = min + 1
		}
	}

	fmt.Println(dp[l])
}
