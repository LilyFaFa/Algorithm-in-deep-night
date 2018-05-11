package main

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	var arr []int
	var result [][]int
	//进行排序
	sort.Ints(candidates)

	combinationS(candidates, target, 0, arr, &result)
	return result
}

func combinationS(candidates []int, target int, start int, arr []int, result *([][]int)) {
	if target == 0 {
		/*
			如果这里不创建新的变量，不重新开辟空间会出错，可能和go的内存管理机制有关
		*/
		a := make([]int, len(arr))
		copy(a, arr)
		*result = append(*result, a)
		return
	}
	for i := start; i < len(candidates); i++ {
		//剪枝
		if candidates[i] > target {
			break
		}
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		target = target - candidates[i]
		/*
			以下两行可以使用注释的三行代替，但是都不可避免需要上层的创建新的变量
			分配空间，否则就会出错
		*/
		a := append(arr, candidates[i])
		combinationS(candidates, target, i+1, a, result)
		/*
			arr := append(arr, candidates[i])
			combinationS(candidates, target, i, arr, result)
			arr = arr[0 : len(arr)-1]
		*/
		target = target + candidates[i]
	}
}
