package main

func climbStairs(n int) int {
	var mem map[int]int
	mem = make(map[int]int)
	return f(n, mem)
}

/*
这里我们没有将子问题记住，所以这个递归就没有被记忆化，时间复杂度比较高，运行就会超时
func f(i int) int {
	if i == 0 || i == 1 {
		return 1
	}
	return f(i-1) + f(i-2)
}
*/

/*
  使用mem做记忆
*/
func f(i int, mem map[int]int) int {
	if i == 0 || i == 1 {
		mem[i] = 1
		return 1
	}
	_, exit := mem[i]
	if exit {
		return mem[i]
	} else {
		mem[i] = f(i-1, mem) + f(i-2, mem)
		return mem[i]
	}
}
