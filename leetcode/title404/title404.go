package main

//树的节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	result := 0
	sum(root, &result, false)
	return result
}

func sum(root *TreeNode, result *int, left bool) {
	// 这步很重要
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if left {
			*result = *result + root.Val
		}
		return
	}
	if root.Right != nil {
		sum(root.Right, result, false)
	}
	if root.Left != nil {
		sum(root.Left, result, true)
	}
}
