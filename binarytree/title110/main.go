package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	balancer := false
	treeHight(root, &balancer)
	return balancer
}

func treeHight(root *TreeNode, balancer *bool) {
	if root == nil {
		return 0
	}

	leftHight := treeHight(root.Left, balancer)
	rightHight := treeHight(root.Right, balancer)
	if leftHight > rightHight+1 || rightHight > leftHight+1 {
		*balancer = false
		return 0
	}

	if leftHight > rightHight {
		leftHight++
	} else if rightHight > leftHight {
		rightHight++
	}
}

func main() {

}
