package main

import (
	"fmt"
)

func main() {
	test := []int{4, 3, 5, 2, 6, 3, 5, 3, 5, 3}
	Sort(test, 0, len(test)-1)
	fmt.Println(test)
}

func Sort(numbers []int, head int, tail int) {

	if head >= tail {
		return
	}

	tmp := numbers[head]
	i := head + 1
	j := tail
	for i < j {
		for numbers[i] <= tmp && i < j {
			i++
		}
		for numbers[j] > tmp && i < j {
			j--
		}

		if i < j {
			tmp := numbers[i]
			numbers[i] = numbers[j]
			numbers[j] = tmp
			i++
			j--
		}
	}
	if i == j && numbers[i] > tmp {
		numbers[head] = numbers[i-1]
		numbers[i-1] = tmp
		Sort(numbers, head, i-2)
		Sort(numbers, i, tail)
		return

	}
	numbers[head] = numbers[j]
	numbers[j] = tmp

	Sort(numbers, head, j-1)
	Sort(numbers, j+1, tail)
}
