package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	var result bool
	result = false
	has(root, 0, sum, &result)
	return result
}

func has(root *TreeNode, tmpsum int, sum int, result *bool) {
	if *result {
		return
	}
	if root == nil {
		return
	}
	tmpsum += root.Val
	if root.Left == nil && root.Right == nil {
		if tmpsum == sum {
			*result = true
		}
		return
	}
	if root.Left != nil {
		has(root.Left, tmpsum, sum, result)
	}
	if root.Right != nil {
		has(root.Right, tmpsum, sum, result)
	}
}

func main() {
	t := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 4,
						},
					},
				},
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	result := hasPathSum(&t, 10)
	fmt.Println(result)
}
