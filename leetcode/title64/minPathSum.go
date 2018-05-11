package main

import (
	"fmt"
)

type point struct {
	x int
	y int
}

func minPathSum(grid [][]int) int {
	mem := make(map[point]int)
	return f(len(grid)-1, len(grid[0])-1, mem, grid)
}

func f(m int, n int, mem map[point]int, grid [][]int) int {

	if m == 1 && n == 0 {
		return grid[1][0] + grid[0][0]
	}
	if m == 0 && n == 1 {
		return grid[0][0] + grid[0][1]
	}
	if m == 0 && n == 0 {
		return grid[0][0]
	}
	// 使用一个超大值代表此路不通
	if m < 0 || n < 0 {
		return 100000000
	}

	p := point{m, n}
	num, exist := mem[p]
	if exist {
		return num
	} else {
		mem[p] = minSubPath(f(m-1, n, mem, grid), f(m, n-1, mem, grid)) + grid[m][n]
	}
	fmt.Println(mem[p])
	return mem[p]
}

func minSubPath(x int, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
