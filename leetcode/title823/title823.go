package main

import (
	"sort"
)

func numFactoredBinaryTrees(candidates []int) int {
	var r int64
	var ss int64
	ss = 1000000007
	result := make(map[int]int64)
	//进行排序
	sort.Ints(candidates)

	//构建一个hash map
	hashMap := make(map[int]bool)
	for _, in := range candidates {
		hashMap[in] = true
		result[in] = 1
	}

	for ind, inv := range candidates {
		for i := 0; i < ind; i++ {
			if inv%candidates[i] == 0 {
				_, ok := hashMap[inv/candidates[i]]
				if ok {
					result[inv] = result[inv] + result[candidates[i]]*result[inv/candidates[i]]
					result[inv] = result[inv] % ss
				}
			}
		}
	}

	for _, re := range result {
		r = r + re
	}
	r = r % ss
	return int(r)

}
