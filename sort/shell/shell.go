package main

import (
	"fmt"
)

func main() {
	numbers := []int{3, 5, 7, 2, 3, 7, 8, 3, 5, 9, 0}
	fmt.Println(numbers)
	shell(numbers)
	fmt.Println(numbers)
}
func shell(numbers []int) {
	fmt.Println(numbers)
	insertSort(numbers, 8)
	fmt.Println(numbers)
	insertSort(numbers, 4)
	fmt.Println(numbers)
	insertSort(numbers, 1)
}

func insertSort(numbers []int, dis int) {
	//步长循环，从插入排序代码中该的
	for i := 0; i < dis; i++ {

		for m := i; m < len(numbers); m += dis {
			for j := i; j < m; j += dis {
				if numbers[j] > numbers[m] {
					var n int
					tmp := numbers[m]
					for n = m; n > j; n -= dis {
						numbers[n] = numbers[n-dis]
					}
					numbers[j] = tmp
					break
				}
			}
		}

	}
}
