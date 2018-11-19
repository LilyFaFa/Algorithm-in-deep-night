package main

import (
	"fmt"

	"math"
)

func main() {
	var number int
	for {
		n, _ := fmt.Scan(&number)
		if n == 0 {
			break
		} else {
			var target int
			fmt.Scan(&target)
			intputs := []int{}

			for i := 0; i < number; i++ {
				var a int
				fmt.Scan(&a)
				intputs = append(intputs, a)
			}

			result := Find(intputs, target)
			fmt.Println(result)
		}
	}
}

func Find(numbers []int, k int) int {
	left, right, mid := 0, len(numbers)-1, 0
	for {
		mid = int(math.Floor(float64((left + right) / 2)))
		if numbers[mid] > k {
			right = mid - 1
		} else if numbers[mid] < k {
			left = mid + 1
		} else {
			for mid+1 < len(numbers) && numbers[mid+1] == k {
				mid++
			}
			break
		}
		if left > right {
			mid = -1
			break
		}
	}
	return mid
}
