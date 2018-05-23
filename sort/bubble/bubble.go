package main

import (
	"fmt"
)

func main() {
	example := []int{2, 5, 1, 7, 2, 9, 3, 6}
	fmt.Println(example)
	sortSort(example)
	fmt.Println(example)
}

func sortSort(numbers []int) {
	var tmp int
	for i := 0; i < len(numbers); i++ {
		for j := len(numbers) - 1; j > i; j-- {
			if numbers[j-1] > numbers[j] {
				tmp = numbers[j]
				numbers[j] = numbers[j-1]
				numbers[j-1] = tmp
			}
		}
	}
}
