package main

type point struct {
	x int
	y int
}

func uniquePaths(m int, n int) int {
	mem := make(map[point]int)

	return f(m, n, mem)
}

func f(m int, n int, mem map[point]int) int {

	if m == 2 && n == 1 {
		return 1
	}
	if m == 1 && n == 2 {
		return 1
	}
	if m == 1 && n == 1 {
		return 1
	}
	if m < 1 || n < 1 {
		return 0
	}

	p := point{m, n}
	num, exist := mem[p]
	if exist {
		return num
	} else {
		mem[p] = f(m-1, n, mem) + f(m, n-1, mem)
	}

	return mem[p]
}
