package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 使用快速排序，这里值更改链表的值，不更改head和tail的指针指向
// 所以不需要使用head和tail的指针的指针
func quickSort(head *ListNode, tail *ListNode) {
	if head == tail {
		return
	}
	pre := head
	index := head.Next

	//pre 指向的值和其前面的都是比 head 要小的值
	//pre 和 index 之间的值都是比 head 要大的值
	for index != tail {
		if index.Val < head.Val {
			tmp := pre.Next.Val
			pre.Next.Val = index.Val
			index.Val = tmp
		}
		index = index.Next
	}

	// 将head值pre值交换一下
	tmp := head.Val
	head.Val = pre.Val
	pre.Val = tmp

	// 这里的pre就是指向middle值，严格意义上而言，这里的pre应该使用pre前面的一个指针
	quickSort(head, pre)
	quickSort(pre.Next, tail)
	return
}

// 使用插入排序，新建一个链表头，将旧链表的节点一次移动到新链表中
// 因为涉及到需要插入节点，所以会构造一个伪头部
func insertSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 为了操作方便，添加一个假头
	var pstart *ListNode

	for tmp := head; tmp != nil; {
		next := tmp
		pre := pstart
		for pre.Next != nil {
			if pre.Next.Val < tmp.Val {
				pre = pre.Next
			} else {
				break
			}
		}
		tmp.Next = pre.Next
		pre.Next = tmp
		tmp = next
	}
	return pstart.Next
}

// 归并排序
func mergeSort(head *ListNode, tail *ListNode) {

}
