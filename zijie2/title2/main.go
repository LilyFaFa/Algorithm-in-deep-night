package main

import (
	"fmt"
)

func main() {
	var M int
	for {
		n, _ := fmt.Scan(&M)
		if n == 0 {
			break
		} else {
			inputs := [][]int{}
			for i := 0; i < M; i++ {
				tmp := []int{}
				for j := 0; j < M; j++ {
					a := 0
					fmt.Scan(&a)
					tmp = append(tmp, a)
				}
				inputs = append(inputs, tmp)
			}
			result := numOfGroup(inputs)
			fmt.Println(result)
		}
	}
}
func numOfGroup(members [][]int) int {
	var number int
	number = 0
	for i := 0; i < len(members); i++ {
		for j := 0; j < len(members[0]); j++ {
			if members[i][j] == 1 {
				number += 1
				dpf(&members, i, j)
			} else {
				continue
			}
		}
	}
	return number
}

func dpf(members *[][]int, x int, y int) {
	if x < 0 || y < 0 || x >= len(*members) || y >= len((*members)[0]) {
		return
	}
	if (*members)[x][y] == 1 {
		(*members)[x][y] = 0
		dpf(members, x+1, y)
		dpf(members, x-1, y)
		dpf(members, x, y+1)
		dpf(members, x, y-1)
	}
	return

}
