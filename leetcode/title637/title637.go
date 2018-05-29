package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type tmp struct {
	time int
	sum  int
}

func averageOfLevels(root *TreeNode) []float64 {
	levelTmp := make(map[int]*tmp)
	dfs(root, 0, &levelTmp)
	result := make([]float64, len(levelTmp))
	for k, v := range levelTmp {
		result[k] = float64(v.sum) / float64(v.time)
	}
	return result
}

func dfs(root *TreeNode, level int, levelTmp *map[int]*tmp) {
	if root != nil {
		_, ok := (*levelTmp)[level]
		if ok {
			(*(*levelTmp)[level]).sum += root.Val
			(*(*levelTmp)[level]).time += 1
		} else {
			t := tmp{
				sum:  root.Val,
				time: 1,
			}
			(*levelTmp)[level] = &t
		}
		dfs(root.Left, level+1, levelTmp)
		dfs(root.Right, level+1, levelTmp)
	}
}
