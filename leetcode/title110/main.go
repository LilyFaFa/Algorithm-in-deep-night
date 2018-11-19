package main

import (
	"fmt"
)

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

	result := isBalanced(&t)
	fmt.Println(result)
}
