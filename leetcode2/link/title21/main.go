package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeSortList(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	result := &ListNode{}
	tail := result

	head1 := list1
	head2 := list2

	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			tail.Next = head1
			head1 = head1.Next
		} else {
			tail.Next = head2
			head2 = head2.Next
		}
		tail = tail.Next
	}

	if head1 != nil {
		tail.Next = head1
	} else if head2 != nil {
		tail.Next = head2
	}
	return result.Next

}
