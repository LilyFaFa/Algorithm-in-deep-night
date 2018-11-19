package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	frenq := make(map[int]int)
	for _, n := range nums {
		fre, ok := frenq[n]
		if ok {
			frenq[n] = fre + 1
		} else {
			frenq[n] = 1
		}
	}

	barrel := make(map[int][]int)
	times := []int{}

	for k, v := range frenq {
		_, ok := barrel[v]
		if ok {
			barrel[v] = append(barrel[v], k)
		} else {
			times = append(times, v)
			items := []int{k}
			barrel[v] = items
		}
	}

	sort.Ints(times)
	result := []int{}

	if k > len(frenq) {
		k = len(frenq)
	}
	fmt.Println(frenq)
	fmt.Println(times)
	fmt.Println(k)
	flag := 0
	i := 0
	for i < k {
		result = append(result, barrel[times[len(times)-flag-1]]...)
		i = i + len(barrel[times[len(times)-flag-1]])
		flag++
	}
	return result

}
func main() {
	input := []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}
	r := topKFrequent(input, 10)
	fmt.Println(r)
}
