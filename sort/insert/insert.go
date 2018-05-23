package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 5, 2, 256, 2}
	fmt.Println(numbers)
	insertSort(numbers)
	fmt.Println(numbers)
}

func insertSort(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < i; j++ {
			if numbers[j] > numbers[i] {
				var m int
				tmp := numbers[i]
				for m = i; m > j; m-- {
					numbers[m] = numbers[m-1]
				}
				numbers[j] = tmp
				break
			}
		}
	}
}
