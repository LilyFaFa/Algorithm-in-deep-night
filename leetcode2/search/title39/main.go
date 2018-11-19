package search

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	tmp := []int{}
	search(candidates, target, 0, tmp, &result)
	return result
}

func search(candidates []int, target int, cur int, curSum int, tmp []int, result *[][]int) {
	// 1.1 递归终止条件
	if curSum > target {
		return
	}
	// 1.2 递归满足条件
	if curSum == target {
		*result = append(*result, tmp)
		return
	}
	// 2. 递归函数处理，这里很关建，只能从当前位置开始，不能从0开始
	for i := cur; i < len(candidates); i++ {
		t := []int{}
		t = append(t, tmp...)
		t = append(t, candidates[i])
		search(candidates, target, i, curSum+candidates[i], t, result)
	}
}
