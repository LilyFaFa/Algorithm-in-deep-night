package main

type point struct {
	x int
	y int
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	//入口堵住，那么就没有任何一条路径
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	mem := make(map[point]int)
	return f(obstacleGrid, len(obstacleGrid), len(obstacleGrid[0]), mem)
}

func f(obstacleGrid [][]int, m int, n int, mem map[point]int) int {

	if m == 2 && n == 1 && obstacleGrid[1][0] == 0 {
		return 1
	}
	if m == 1 && n == 2 && obstacleGrid[0][1] == 0 {
		return 1
	}
	if m == 1 && n == 1 && obstacleGrid[0][0] == 0 {
		return 1
	}
	if m < 1 || n < 1 {
		return 0
	}

	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	p := point{m, n}
	num, exist := mem[p]
	if exist {
		return num
	} else {
		mem[p] = f(obstacleGrid, m-1, n, mem) + f(obstacleGrid, m, n-1, mem)
	}

	return mem[p]
}
