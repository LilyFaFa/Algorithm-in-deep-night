package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	return merge(t1, t2)
}

func merge(t1 *TreeNode, t2 *TreeNode) *TreeNode {

	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	t1.Val = t2.Val + t1.Val
	t1.Left = merge(t1.Left, t2.Left)
	t1.Right = merge(t1.Right, t2.Right)
	return t1
}
