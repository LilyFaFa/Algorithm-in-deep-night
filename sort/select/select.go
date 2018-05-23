package main

import (
	"fmt"
)

func main() {
	numbers := []int{3, 4, 7, 2, 4, 1, 7, 3}
	fmt.Println(numbers)
	selectSort(numbers)
	fmt.Println(numbers)

}

func selectSort(numbers []int) {

	for i := 0; i < len(numbers); i++ {
		min := i
		for j := len(numbers) - 1; j >= i; j-- {
			if numbers[j] < numbers[min] {
				min = j
			}
		}
		tmp := numbers[min]
		numbers[min] = numbers[i]
		numbers[i] = tmp
	}
}
