package title2

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l11 := l1
	l22 := l2
	le := &ListNode{
		Val:  0,
		Next: nil,
	}
	l := le

	tem1 := 0
	tem2 := 0

	for l11 != nil && l22 != nil {
		tem := l11.Val + l22.Val
		tem1 = tem % 10
		tem2 = tem / 10

		l.Next = &ListNode{
			Val:  tem1,
			Next: nil,
		}

		l = l.Next
		l11 = l11.Next
		l22 = l22.Next

		if l11 != nil {
			l11.Val = l11.Val + tem2
		} else if l22 != nil {
			l22.Val = l22.Val + tem2
		}
	}

	if l11 != nil || l22 != nil {
		tmp1 := 0
		tmp2 := 0
		for l11 != nil {
			tmp1 = (l11.Val + tmp2) % 10
			tmp2 = (l11.Val + tmp2) / 10
			l.Next = &ListNode{
				Val:  tmp1,
				Next: nil,
			}
			l = l.Next
			l11 = l11.Next

		}
		for l22 != nil {
			tmp1 = (l22.Val + tmp2) % 10
			tmp2 = (l22.Val + tmp2) / 10
			l.Next = &ListNode{
				Val:  tmp1,
				Next: nil,
			}
			l = l.Next
			l22 = l22.Next
		}
		if tmp2 != 0 {
			l.Next = &ListNode{
				Val:  1,
				Next: nil,
			}
		}
	} else if tem2 != 0 {
		l.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}
	//remove useless listHead
	return le.Next
}
