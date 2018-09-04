package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	h := head
	lenth := 0
	for head.Next != nil {
		lenth++
		head = head.Next
	}
	lenth++
	k = k % lenth
	move := lenth - k
	t := head
	t.Next = h

	for i := 0; i < move; i++ {
		h = h.Next
		t = t.Next
	}
	t.Next = nil
	return h
}
