package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {

	for {
		var aim string
		_, err := fmt.Scanf("%s", &aim)
		if err == io.EOF {
			break
		} else {
			reader := bufio.NewReader(os.Stdin)
			in1, _, _ := reader.ReadLine()
			inputs1 := strings.Split(string(in1), " ")

			s := []int{}
			resultMap := make(map[int]int)

			for i := 0; i < 3; i++ {
				nu := minDistance(inputs1[i], aim)
				s = append(s, nu)
				resultMap[i] = nu
			}

			sort.Ints(s)
			for i := 3; i < len(inputs1); i++ {
				nu := minDistance(inputs1[i], aim)
				resultMap[i] = nu
				if nu < s[2] {
					s = s[:2]
					s = append(s, nu)
				}
				sort.Ints(s)
			}

			results := []string{}
			for j := 0; j < 3; j++ {
				flag := true
				for i := 0; i < len(inputs1); i++ {
					if len(results) == 3 {
						flag = false
						break
					}
					if resultMap[i] == s[j] {
						results = append(results, inputs1[i])
					}
				}

				if !flag {
					break
				}
			}
			fmt.Printf("%s %s %s", results[0], results[1], results[2])
		}
	}

}

var one = "qwertasdfgzxcv"

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	cost := 0

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = min(dp[i-1][j]+3, dp[i][j-1]+3)

			if word1[i-1] == word2[j-1] {
				cost = 0
			} else {
				if strings.Contains(one, string(word2[j-1])) && strings.Contains(one, string(word1[i-1])) || !strings.Contains(one, string(word2[j-1])) && !strings.Contains(one, string(word1[i-1])) {
					cost = 1
				} else {
					cost = 2
				}
			}

			dp[i][j] = min(dp[i][j], dp[i-1][j-1]+cost)
		}
	}

	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
