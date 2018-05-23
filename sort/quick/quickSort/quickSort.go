package main

import (
	"fmt"
)

func main() {
	numbers := []int{4, 2, 5, 1, 6, 9, 7}
	fmt.Println(numbers)
	quickSort(numbers)
	fmt.Println(numbers)
}

func quickSort(numbers []int) {
	if len(numbers) == 1 {
		return
	}

	head := 1
	tail := len(numbers) - 1
	middle := 0

	for tail >= head {
		//这两个与的循序不能换
		for head < len(numbers) && numbers[head] < numbers[0] {
			head++
		}
		for tail > 0 && numbers[tail] > numbers[0] {
			tail--
		}

		if head < len(numbers) && tail > 0 {
			tmp := numbers[head]
			numbers[head] = numbers[tail]
			numbers[tail] = tmp
			tail--
			head++
		}
	}

	if tail > 0 {
		tmp := numbers[0]
		numbers[0] = numbers[tail]
		numbers[tail] = tmp
		middle = tail
	}

	if middle > 0 {
		quickSort(numbers[:middle])
	}
	if middle < len(numbers)-1 {
		quickSort(numbers[middle+1:])
	}
}
