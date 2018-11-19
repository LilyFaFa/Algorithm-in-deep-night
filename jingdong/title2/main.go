package main

import (
	"fmt"
	"io"
)

func main() {
	for {
		var num int
		_, err := fmt.Scanf("%d", &num)
		if err == io.EOF {
			break
		} else {
			inputs := [][]int{}
			for i := 0; i < num+1; i++ {
				input := make([]int, num+1)
				inputs = append(inputs, input)
			}
			for i := 0; i < num-1; i++ {
				var x, y int
				fmt.Scanf("%d", &x)
				fmt.Scanf("%d", &y)
				inputs[x][y] = 1
			}
			result := 0
			fir := []int{}
			second := []int{}
			for i := 1; i < len(inputs); i++ {
				for j := 1; j < len(inputs); j++ {
					if i == j {
						continue
					}
					fir = append(fir, i)
					second = append(second, j)
					dp(inputs, &result, fir, second)
					fir = []int{}
					second = []int{}
				}

			}

			fmt.Println(result)
		}
	}
}

func dp(inputs [][]int, max *int, fir []int, second []int) {
	if len(fir)*len(second) > *max {
		*max = len(fir) * len(second)
	}

	if len(fir)+len(second) == len(inputs)-1 {
		return
	}

	for i := 1; i < len(inputs); i++ {
		if inputs[fir[len(fir)-1]][i] == 1 {
			flag := true
			for _, s := range second {
				if s == i {
					flag = false
					break
				}
			}
			if !flag {
				fir = append(fir, i)
				dp(inputs, max, fir, second)
				fir = fir[:len(fir)-1]
			}
		}

	}
	for i := 1; i < len(inputs); i++ {
		if inputs[second[len(second)-1]][i] == 1 {
			flag := true
			for _, s := range fir {
				if s == i {
					flag = false
					break
				}
			}
			if !flag {
				second = append(second, i)
				dp(inputs, max, fir, second)
				second = second[:len(fir)-1]
			}
		}

	}
}
