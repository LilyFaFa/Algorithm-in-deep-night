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
			if len(numbers) < 3 {
				fmt.Println(0)
				continue
			}
			sort.Ints(numbers)
			result := 0
			for i := 0; i < len(numbers)-2; i++ {
				for j := i + 1; j < len(numbers)-1; j++ {
					for m := j + 1; m < len(numbers); m++ {
						if numbers[j]-numbers[i] > 18000 || numbers[m]-numbers[j] > 18000 || numbers[m]-18000 < numbers[i] {
							result++
						}
					}
				}
			}
			fmt.Println(result)
		}

	}
}
