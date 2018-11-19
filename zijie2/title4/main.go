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
			inputs := []int{}
			for i := 0; i < M; i++ {
				tmp := 0
				fmt.Scan(&tmp)
				inputs = append(inputs, tmp)
			}
			fmt.Println(0)
		}
	}
}

func check(intputs []int) bool {
	for _,input:=range inputs{
		if input< 0b10000000{
		}
	}
	return false
}

