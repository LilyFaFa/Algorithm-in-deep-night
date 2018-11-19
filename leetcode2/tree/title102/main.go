package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var nextLevel []*TreeNode
	result := [][]int{}
	if root == nil {
		return result
	} else {
		first := []int{root.Val}
		if root.Left != nil {
			nextLevel = append(nextLevel, root.Left)
		}
		if root.Right != nil {
			nextLevel = append(nextLevel, root.Right)
		}
		result = append(result, first)
	}

	for len(nextLevel) != 0 {
		var cur []int
		l := len(nextLevel)
		for i := 0; i < l; i++ {
			node := nextLevel[i]
			cur = append(cur, node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		result = append(result, cur)
		nextLevel = nextLevel[l:]
	}
	return result
}
