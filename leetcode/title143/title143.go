package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	var p1 *ListNode
	var p2 *ListNode
	p1 = head
	p2 = head
	for p1.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	f := p1.Next
	fmt.Println(f)
	revertList(f)
	fmt.Println(f)
	last := f
	first := head

	for last.Next != nil {
		next := first.Next
		first.Next = last
		last = last.Next
		first.Next.Next = next
		first = next
	}
	first.Next = last
}

//反转指针
func revertList(head *ListNode) {
	var pre *ListNode
	var next *ListNode
	p := head
	//这样也不行
	//for head != nil {
	//	head = head.Next
	//}
	for p != nil {
		next = p.Next
		p.Next = pre
		pre = p
		p = next
	}
	//切记,这样会改变haed指针指向的变量的值
	//*head=*pre
	//这样也不会改变真实head
	//head = pre
	//这样也会改变head指向的变量的值
	//head.Next = pre.Next
	//head.Val = pre.Val
}

func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	fmt.Println(l, l.Next, l.Next.Next, l.Next.Next.Next)
	revertList(l)
	//fmt.Println(l, l.Next, l.Next.Next, l.Next.Next.Next)
	fmt.Println(l, l.Next)
}
