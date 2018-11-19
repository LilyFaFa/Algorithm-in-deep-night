package main

import (
	"sort"
)

func increasingTriplet(nums []int) bool {
	tail, res := make([]int, len(nums)), 0
	for _, n := range nums {
		i := sort.SearchInts(tail[:res], n)
		tail[i] = n
		if i == res {
			res++
		}
	}
	return res >= 3
}
