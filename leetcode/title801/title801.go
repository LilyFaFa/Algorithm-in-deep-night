package main

import (
	"fmt"
)

func minSwap(A []int, B []int) int {
	tmp := [][]int{}
	for i := 0; i < len(A); i++ {
		t := []int{10000, 10000}
		tmp = append(tmp, t)
	}
	tmp[0][0] = 0
	tmp[0][1] = 1

	for i := 1; i < len(A); i++ {

		if A[i-1] < A[i] && B[i-1] < B[i] {
			tmp[i][0] = tmp[i-1][0]
		}

		if A[i-1] < B[i] && B[i-1] < A[i] {
			if tmp[i][0] > tmp[i-1][1] {
				tmp[i][0] = tmp[i-1][1]
			}
		}

		if A[i-1] < B[i] && B[i-1] < A[i] {
			tmp[i][1] = tmp[i-1][0] + 1
		}

		if A[i-1] < A[i] && B[i-1] < B[i] {
			if tmp[i][1] > tmp[i-1][1] {
				tmp[i][1] = tmp[i-1][1] + 1
			}
		}
	}

	if tmp[len(A)-1][0] < tmp[len(A)-1][1] {
		return tmp[len(A)-1][0]
	} else {
		return tmp[len(A)-1][1]
	}

}

func main() {
	tA := []int{0, 4, 4, 5, 9}
	tB := []int{0, 1, 6, 8, 10}
	result := minSwap(tA, tB)
	fmt.Println(result)
}
