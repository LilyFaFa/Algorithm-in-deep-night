package main

import (
	"fmt"
)

func rotate(matrix [][]int) {
	circle := len(matrix)
	for i := 0; i < circle/2; i++ {
		for j := i; j < circle-i-1; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[circle-1-j][i]
			matrix[circle-1-j][i] = matrix[circle-1-i][circle-1-j]
			matrix[circle-1-i][circle-1-j] = matrix[j][circle-1-i]
			matrix[j][circle-1-i] = tmp
		}
	}

}

func main() {
	var m int
	var test [][]int
	fmt.Scanf("%d", &m)
	for j := 0; j < m; j++ {
		var items []int
		for i := 0; i < m; i++ {
			var item int
			fmt.Scanf("%d", &item)
			items = append(items, item)
		}
		test = append(test, items)
	}

	fmt.Println(test)
	rotate(test)
	fmt.Println(test)
}
