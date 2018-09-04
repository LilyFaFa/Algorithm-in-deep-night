package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = &ListNode{
		Val: 0,
	}
	tail = tail.Next
	root := convertNode(head, tail)
	return root
}

func convertNode(head *ListNode, tail *ListNode) *TreeNode {
	h1 := head
	h2 := head

	if head.Next == tail {
		tree := &TreeNode{
			Val: head.Val,
		}
		return tree
	}

	for h1.Next != tail {
		h1 = h1.Next
		h2 = h2.Next
		if h1.Next != tail {
			h1 = h1.Next
		}
	}

	tree := &TreeNode{
		Val: h2.Val,
	}
	tree.Left = convertNode(head, h2)
	if h2.Next != tail {
		tree.Right = convertNode(h2.Next, tail)
	}
	return tree
}
