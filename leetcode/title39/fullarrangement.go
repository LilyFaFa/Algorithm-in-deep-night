package main

/*
  numbers 的长度为n的全排列
*/
func arrange(n int, numbers []int, used []int, arr []int, result *([][]int)) {
	if len(arr) == n {
		*result = append(*result, arr)
		return
	}

	if len(used) == 0 {
		for i := 0; i < len(numbers); i++ {
			used = append(used, 0)
		}
	}

	//如果不重新创建一个变量，那么会导致变量被更改
	//var a []int
	for i := 0; i < len(numbers); i++ {
		if used[i] == 0 {
			a := append(arr, numbers[i])
			used[i] = 1
			arrange(n, numbers, used, a, result)
			used[i] = 0
		}

	}

}
