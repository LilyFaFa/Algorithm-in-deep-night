package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 4, 2, 6, 3, 6, 3, 8, 5, 9, 0, 7}
	mergeSort(numbers, 0, len(numbers)-1)
	fmt.Println(numbers)
}
func mergeSort(numbers []int, first int, tail int) {
	var middle int
	if first < tail {
		middle = (first + tail) / 2
		mergeSort(numbers, first, middle)
		mergeSort(numbers, middle+1, tail)
		merge(numbers, first, middle, tail)
	}
}

func merge(numbers []int, first int, middle int, tail int) {
	tmp := []int{}
	i := first
	j := middle + 1
	for j <= tail && i <= middle {
		if numbers[i] < numbers[j] {
			tmp = append(tmp, numbers[i])
			i++
			continue
		} else {
			tmp = append(tmp, numbers[j])
			j++
			continue
		}
	}
	if i != middle+1 {
		tmp = append(tmp, numbers[i:middle+1]...)
	} else if j != tail+1 {
		tmp = append(tmp, numbers[j:tail+1]...)
	}
	for i := 0; i < len(tmp); i++ {
		numbers[i+first] = tmp[i]
	}
}
