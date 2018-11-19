package main

import (
	"fmt"
	"io"
	"sort"
)

func main() {
	for {
		var num int
		_, err := fmt.Scanf("%d", &num)
		if err == io.EOF {
			break
		} else {
			var numbers []int
			for i := 0; i < num; i++ {
				var a int
				fmt.Scanf("%d", &a)
				numbers = append(numbers, a)

			}

			var numbersTmp []int
			for i := 0; i < len(numbers); i++ {
				numbersTmp = append(numbersTmp, numbers[i])
			}
			sort.Ints(numbersTmp)
			for _, num1 := range numbers {
				sum := 0
				for _, num2 := range numbersTmp {
					if num2 < num1 {
						sum += num2
					}
				}
				fmt.Println(sum)
			}

		}
	}
}
