package main

import (
	"fmt"
	"sort"
)

func eventualSafeNodes(graph [][]int) []int {
	safty := make(map[int]bool)
	result := []int{}
	flag := 1
	for flag == 1 {
		flag = 0
		for i := 0; i < len(graph); i++ {
			if len(graph[i]) == 0 {
				_, ok := safty[i]
				if !ok {
					safty[i] = true
					result = append(result, i)
					flag = 1
				}
				continue
			}
			for j := 0; j < len(graph[i]); j++ {
				_, ok := safty[graph[i][j]]
				if ok {
					graph[i] = append(graph[i][:j], graph[i][j+1:]...)
					flag = 1
				}
			}
		}
	}
	fmt.Println(graph)
	sort.Ints(result)
	return result
}
func main() {
	example := [][]int{
		{1, 2},
		{2, 3},
		{5},
		{0},
		{5},
		{},
		{},
	}
	result := eventualSafeNodes(example)
	fmt.Println(result)
}
