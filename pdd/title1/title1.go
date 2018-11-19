package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scanf("%s", &input)
	for i := 0; i <= len(input)/4; i++ {
		fmt.Print(string(input[i]))
	}
	fmt.Println()
	for j := 1; j < len(input)/4; j++ {
		fmt.Print(string(input[len(input)-j]))
		for a := 1; a < len(input)/4; a++ {
			fmt.Print(" ")
		}
		fmt.Print(string(input[len(input)/4+j]))
		fmt.Println()
	}

	for i := 3 * (len(input) / 4); i >= len(input)/2; i-- {
		fmt.Print(string(input[i]))
	}
}
