package main

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	balancer := true
	treeHight(root, &balancer)
	return balancer

}

/*
求树高
*/
func treeHight(root *TreeNode, balancer *bool) int {
	if root == nil {
		return 0
	}

	leftHight := treeHight(root.Left, balancer)
	/*
		这一步操作放在这里也可以，放在这里就可以只要判断左子树不平衡就可以认为不平衡
			if !(*balancer) {
				return 0
			}
	*/
	rightHight := treeHight(root.Right, balancer)
	if ((leftHight - rightHight) > 1) || ((leftHight - rightHight) < -1) {
		*balancer = false
	}
	//如果子树不平衡，提前结束,会提高时间复杂度
	if !(*balancer) {
		return 0
	}

	if leftHight > rightHight {
		return leftHight + 1
	} else {
		return rightHight + 1
	}
}
