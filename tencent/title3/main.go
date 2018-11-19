package main

import (
	"fmt"
)

func main() {
	var n int
	for {
		l, _ := fmt.Scan(&n)
		if l == 0 {
			break
		} else {
			var m int
			fmt.Scan(&m)
			inputs := [][]int{}
			for i := 0; i < n; i++ {
				input := make([]int, n)
				inputs = append(inputs, input)
			}

			for i := 0; i < m; i++ {
				var a int
				fmt.Scan(&a)
				var b int
				fmt.Scan(&b)
				if a == b {
					continue
				} else {
					inputs[a-1][b-1] = 1
				}
			}

			fmt.Println(inputs)
			totle := 0
			out := make([]int, n)
			for i := 0; i < n; i++ {
				cur := 0
				for j := 0; j < n; j++ {
					if inputs[i][j] == 1 {
						cur++
					}
				}
				out[i] = cur
			}

			for i := 0; i < n; i++ {
				cur := 0
				for j := 0; j < n; j++ {
					if inputs[j][i] == 1 {
						cur++
					}
				}
				if cur > out[i] {
					fmt.Println(i)
					totle++
				}
			}
			fmt.Println(totle)
		}
	}
}
