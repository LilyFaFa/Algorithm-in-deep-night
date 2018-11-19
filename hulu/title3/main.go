package main

import (
	"fmt"
	"io"
)

func main() {
	for {
		var m int
		_, err := fmt.Scanf("%d", &m)
		if err == io.EOF {
			break
		} else {
			var n int
			var k int
			fmt.Scanf("%d", &n)
			fmt.Scanf("%d", &k)
			depMap := make(map[int]int)
			ready := []int{}

			depend := make([][]int, m+1)
			for i := 0; i < m+1; i++ {
				d := make([]int, m+1)
				depend = append(depend, d)
			}
			for i := 0; i < k; i++ {
				var x, y int
				fmt.Scanf("%d", &x)
				fmt.Scanf("%d", &y)
				if _, ok := depMap[x]; ok {
					depMap[x]++
				} else {
					depMap[x] = 1
				}
				depend[x][y] = 1
			}

			for i := 1; i <= m; i++ {
				if _, ok := depMap[i]; !ok {
					ready = append(ready, i)
				}
			}

			finished := 0
			result := 0
			for finished != m {
				result++
				max := n
				if len(ready) < max {
					max = len(ready)
				}

				for i := 0; i < max; i++ {
					j := ready[i]
					finished++
					for k := 1; k <= m; k++ {
						if depend[k][j] == 1 {
							depMap[k] = depMap[k] - 1
							if depMap[k] == 0 {
								ready = append(ready, k)
							}
						}
					}
				}
				ready = ready[max:]
			}
			fmt.Println(result)
		}
	}
}
