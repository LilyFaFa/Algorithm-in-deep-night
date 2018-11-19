package main

import (
	"fmt"
	"io"
)

func main() {
	for {
		var number int
		_, err := fmt.Scanf("%d", &number)
		if err == io.EOF {
			break
		} else {
			var numbers []int
			for i := 0; i < number; i++ {
				var a int
				fmt.Scanf("%d", &a)
				numbers = append(numbers, a)
			}
			result := 1000000000000
			change(numbers, &result, 0, 0)
			if result == 1000000000000 {
				result = -1
			}
			fmt.Println(result)
		}
	}
}
func change(numbers []int, result *int, curresult int, cur int) {
	if cur == len(numbers) {
		if *result > curresult {
			*result = curresult
		}
		return
	}

	if cur == 0 {
		change(numbers, result, 0, 1)
		numbers[0] = numbers[0] * -1
		change(numbers, result, 1, 1)
		return
	}

	if numbers[cur] >= numbers[cur-1] {
		change(numbers, result, curresult, cur+1)
	}

	if numbers[cur]*(-1) >= numbers[cur-1] {
		numbers[cur] = numbers[cur] * (-1)
		change(numbers, result, curresult+1, cur+1)
	}

}
