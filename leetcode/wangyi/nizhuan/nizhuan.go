package main

import (
	"fmt"
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
	output := nizhuan(input)
	for i := 0; i < len(output)-1; i++ {
		fmt.Print(output[i])
		fmt.Print(" ")
	}
	fmt.Print(output[len(output)-1])

}
func nizhuan(intput []int) []int {
	if len(intput) < 2 {
		return intput
	}
	output := make([]int, len(intput))
	var middle int
	if len(output)%2 == 1 {
		middle = len(intput) / 2
	} else {
		middle = len(intput) / 2
	}
	output[middle] = intput[0]
	if len(intput)%2 == 0 {
		for i := 1; i < len(intput); i++ {
			if i%2 == 1 {
				output[middle-i/2-1] = intput[i]
			} else {
				output[middle+i/2] = intput[i]
			}
		}
	} else {
		for i := 1; i < len(intput); i++ {
			if i%2 == 0 {
				output[middle-i/2] = intput[i]
			} else {
				output[middle+i/2+1] = intput[i]
			}
		}
	}
	return output
}

/*
package main

import (
	"fmt"
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
	output := nizhuan(input)
	for i := 0; i < len(output)-1; i++ {
		fmt.Print(output[i])
		fmt.Print(" ")
	}
	fmt.Print(output[len(output)-1])

}
func nizhuan(intput []int) []int {
	if len(intput) < 2 {
		return intput
	}
	output := make([]int, len(intput))
	var middle int
	if len(output)%2 == 1 {
		middle = len(intput) / 2
	} else {
		middle = len(intput) / 2
	}
	output[middle] = intput[0]
	if len(intput)/2 == 0 {
		for i := 1; i < len(intput); i++ {
			if i%2 == 1 {
				output[middle-i/2-1] = intput[i]
			} else {
				output[middle+i/2] = intput[i]
			}
		}
	} else {
		for i := 1; i < len(intput); i++ {
			if i%2 == 0 {
				output[middle-i/2] = intput[i]
			} else {
				output[middle+i/2+1] = intput[i]
			}
		}
	}
	return output
}

*/
