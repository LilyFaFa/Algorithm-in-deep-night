package main

import (
	"strconv"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(t *TreeNode) string {
	result := ""

	if t == nil {
		return result
	} else {
		p(t, &result)
	}

	return result
}

func p(t *TreeNode, s *string) {

	if t == nil {
		return
	}

	*s = *s + strconv.Itoa(t.Val)
	if t.Right != nil || t.Left != nil {
		*s = *s + "("
		p(t.Left, s)
		*s = *s + ")"
	}
	if t.Right != nil {
		*s = *s + "("
		p(t.Right, s)
		*s = *s + ")"
	}

}

/*
func tree2str(t *TreeNode) string {
	result := ""
	if t == nil {
		return result
	} else {
		result = strconv.Itoa(t.Val)
	}
	p(t, &result, false)
	p(t, &result, true)
	return result
}

func p(t *TreeNode, s *string, right bool) {
	if right {
		if t == nil {
			*s = *s + "("
		} else {
			*s = *s + "("
		}
	} else {
		if t == nil {
			*s = *s + ")"
		} else {
			*s = *s + ")(" + strconv.Itoa(t.Val) + ")"
		}
	}
	p(t.Left, s, right)
	p(t.Left, s, false)
}
*/
