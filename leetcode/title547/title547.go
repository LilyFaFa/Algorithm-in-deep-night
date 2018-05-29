package main

import (
	"fmt"
)

func findCircleNum(M [][]int) int {
	number := 0
	tmp := make(map[int]bool)
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[0]); j++ {
			_, ok1 := tmp[i]
			_, ok2 := tmp[j]
			if M[i][j] == 1 && !ok1 && !ok2 {
				tmp[i] = true
				tmp[j] = true
				number++
				dfs(M, i, &tmp)
				dfs(M, j, &tmp)
			}
		}
	}
	fmt.Println(M)
	return number
}

func dfs(M [][]int, x int, tmp *map[int]bool) {
	if x >= len(M) || x < 0 {
		return
	}
	for h := 0; h < len(M); h++ {
		_, ok := (*tmp)[h]
		if ok {
			continue
		}
		if M[x][h] == 1 {
			(*tmp)[h] = true
			dfs(M, h, tmp)
		}
	}
}

func main() {
	testCase := [][]int{
		{1, 0, 0, 1},
		{0, 1, 1, 0},
		{0, 1, 1, 1},
		{1, 0, 1, 1},
	}
	result := findCircleNum(testCase)
	fmt.Println(result)
}
