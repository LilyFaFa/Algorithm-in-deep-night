package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	// 如果跟节点的值比key大，那么递归处理左子树，相反处理右子树
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Right == nil && root.Left == nil {
			return nil
		} else if root.Right == nil {
			return root.Left
		} else if root.Left == nil {
			return root.Right
		} else {
			//寻找右子树的最小值替换跟节点值为新的跟节点
			var min *TreeNode
			min = root.Right
			for min.Left != nil {
				min = min.Left
			}
			root.Val = min.Val
			root.Right = deleteNode(root.Right, min.Val)
			return root
		}
	}
	return root

}
