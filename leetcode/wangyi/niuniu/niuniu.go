package main

import (
	"fmt"
	"sort"
)

func main() {
	var m int
	var n int
	fmt.Scanf("%d", &m)
	fmt.Scanf("%d", &n)
	mn := make(map[int]int)
	val := make([]int, m)
	power := make([]int, n)
	for i := 0; i < m; i++ {
		var k int
		fmt.Scanf("%d", &k)
		var v int
		fmt.Scanf("%d", &v)
		mn[k] = v
		val[i] = k
	}
	for i := 0; i < n; i++ {
		var k int
		fmt.Scanf("%d", &k)
		power[i] = k
	}
	sort.Ints(val)
	niuniu(val, mn, power)
}
func niuniu(input []int, mn map[int]int, power []int) {
	var value []int
	for _, i := range input {
		v := mn[i]
		value = append(value, v)
	}
	for i := 0; i < len(input)-1; i++ {
		if value[i+1] < value[i] {
			value[i+1] = value[i]
		}
	}
	for i := 0; i < len(power)-1; i++ {
		k := sort.SearchInts(input, power[i])

		if k >= len(input) {
			k = len(input) - 1
		}
		if power[i] < input[k] {
			k = k - 1
		}
		fmt.Println(value[k])
	}
	k := sort.SearchInts(input, power[len(power)-1])
	if k >= len(input) {
		k = len(input) - 1
	}
	if power[len(power)-1] < input[k] {
		k = k - 1
	}
	fmt.Println(value[k])
}
