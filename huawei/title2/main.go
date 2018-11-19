package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input1 []int
	var input int
	a := 0
	b := 0

	for i := 0; i < 4; i++ {
		fmt.Scanf("%d", &input)
		input1 = append(input1, input)
	}

	for i := 0; i < 4; i++ {
		fmt.Scanf("%d", &input)
		for index, num := range input1 {
			if num == input && index == i {
				a++
			} else if num == input {
				b++
			}
		}
	}
	fmt.Println(strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B")

}
