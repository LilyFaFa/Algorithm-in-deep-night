package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	var max int
	max = -100000000
	maxResult(root, &max)
	if max == -100000000 {
		max = 0
	}
	return max
}

func maxResult(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	tmp := root.Val
	rightlargest := maxResult(root.Left, max)
	leftlargest := maxResult(root.Right, max)
	if rightlargest > 0 {
		tmp = tmp + rightlargest
	}
	if leftlargest > 0 {
		tmp = tmp + leftlargest
	}
	if tmp > *max {
		*max = tmp
	}

	if rightlargest > leftlargest && rightlargest > 0 {
		return root.Val + rightlargest
	} else if rightlargest < leftlargest && leftlargest > 0 {
		return root.Val + leftlargest
	}
	return root.Val
}
func main() {

}
