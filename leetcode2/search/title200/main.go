package search

func numIslands(grid [][]byte) int {
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				result++
				search(&grid, i, j)
			}
		}
	}
	return result
}

func search(grid *[][]byte, x int, y int) {
	// 1.1 递归终止条件
	if x >= len(*grid) || y >= len((*grid)[0]) || x < 0 || y < 0 {
		return
	}

	// 1.2 递归终止条件
	if (*grid)[x][y] == '0' {
		return
	}

	(*grid)[x][y] = '0'
	search(grid, x+1, y)
	search(grid, x, y+1)
	search(grid, x-1, y)
	search(grid, x, y-1)
}
