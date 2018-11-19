package main

import (
	"fmt"
)

func main() {
	var N int
	for {
		n, _ := fmt.Scan(&N)
		if n == 0 {
			break
		} else {
			M := 0
			fmt.Scan(&M)
			members := [][]int{}
			for i := 0; i < N+1; i++ {
				member := make([]int, N+1)
				members = append(members, member)
			}
			for i := 0; i < M; i++ {
				x := 0
				fmt.Scan(&x)
				y := 0
				fmt.Scan(&y)
				members[x][y] = 1
			}
			changed := true
			for changed {
				changed = false
				for i := 1; i < len(members); i++ {
					c := search(&members, i)
					if c {
						changed = true
					}
				}
			}

			result := 0
			for i := 1; i < len(members); i++ {
				flag := true
				for j := 1; j < len(members); j++ {
					if j == i {
						continue
					}
					if members[j][i] == 0 {
						flag = false
						break
					}
				}
				if flag {
					result++
				}
			}
			fmt.Println(result)
		}
	}
}

func search(members *[][]int, x int) bool {
	changed := false
	for index1, member := range (*members)[x] {
		if member == 1 {
			for index2, m := range (*members)[index1] {
				if m == 1 && (*members)[x][index2] == 0 {
					changed = true
					(*members)[x][index2] = 1
				}
			}
		}
	}
	return changed
}
