package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fake := &ListNode{}
	fake.Next = head
	h1 := fake
	h2 := fake
	for i := 0; i < n; i++ {
		h1 = h1.Next
	}
	for h1.Next != nil {
		h1 = h1.Next
		h2 = h2.Next
	}
	h2.Next = h2.Next.Next
	return fake.Next
}
