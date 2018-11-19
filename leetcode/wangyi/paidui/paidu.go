package main

import (
	"fmt"
	"sort"
)

func main() {
	var m int
	fmt.Scanf("%d", &m)
	var input []int
	for i := 0; i < m; i++ {
		var in int
		fmt.Scanf("%d", &in)
		input = append(input, in)
	}
	fmt.Println(paidui(input))
}

func paidui(input []int) int {
	sort.Ints(input)
	output := make([]int, len(input))
	if len(input) == 1 {
		return 0
	}
	if len(input) == 2 {
		if input[0] > input[1] {
			return input[0] - input[1]
		} else {
			return input[1] - input[0]
		}
	}
	var middle int

	if len(input)%2 == 0 {
		middle = len(input) / 2
		output[middle] = input[len(input)-1]
		flag := 1
		j := len(input) - 2
		i := 0
		for i < j {
			output[middle-flag] = input[i]
			i++
			if i >= j {
				break
			}
			output[middle+flag] = input[i]
			i++
			if i >= j {
				break
			}
			flag++
			output[middle-flag] = input[j]
			if i >= j {
				break
			}
			j--
			output[middle+flag] = input[j]
			if i >= j {
				break
			}
			j--
			flag++
		}
		output[0] = input[j]
	} else {
		middle = len(input) / 2
		output[middle] = input[len(input)-1]
		flag := 1
		j := len(input) - 2
		i := 0
		for i < j {
			output[middle-flag] = input[i]
			i++
			if i > j {
				break
			}
			output[middle+flag] = input[i]
			i++
			if i > j {
				break
			}
			flag++
			output[middle-flag] = input[j]
			j--
			if i > j {
				break
			}
			output[middle+flag] = input[j]
			j--
			if i > j {
				break
			}
			flag++
		}
	}
	result := 0
	for i := 0; i < len(output)-1; i++ {
		if output[i+1]-output[i] < 0 {
			result = result - (output[i+1] - output[i])
		} else {
			result = result + (output[i+1] - output[i])
		}
	}
	fmt.Println(output)
	return result
}
