package main

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	cout := 0
	for _, item := range g {
		index := sort.SearchInts(s, item)
		if index != len(s) {
			cout++
			s = append(s[:index], s[index+1:]...)
		}
	}
	return cout
}
