package main

import (
	"fmt"
)

func numIslands(grid [][]byte) int {
	fmt.Println(grid)
	var number int
	number = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				number += 1
				dpf(&grid, i, j)
			} else {
				continue
			}
		}
	}
	fmt.Println(grid)
	return number
}
func dpf(grid *[][]byte, x int, y int) {
	if x < 0 || y < 0 || x >= len(*grid) || y >= len((*grid)[0]) {
		return
	}
	if (*grid)[x][y] == '1' {
		(*grid)[x][y] = '0'
		dpf(grid, x+1, y)
		dpf(grid, x-1, y)
		dpf(grid, x, y+1)
		dpf(grid, x, y-1)
	}
	return

}

func main() {
	testCase := [][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	}
	result := numIslands(testCase)
	fmt.Println(result)

}
