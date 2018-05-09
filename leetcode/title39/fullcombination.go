package main

func combination(n int, numbers []int, start int, arr []int, result *([][]int)) {
	if len(arr) == n {
		*result = append(*result, arr)
		return
	}

	if start >= len(numbers) {
		return
	}

	for i := start; i < len(numbers); i++ {
		a := append(arr, numbers[i])
		combination(n, numbers, i+1, a, result)
	}
}
