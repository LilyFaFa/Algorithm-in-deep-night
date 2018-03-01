package title2

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {

	testCase1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	testCase2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  6,
				Next: nil,
			},
		},
	}

	actual := addTwoNumbers(testCase1, testCase2)
	expected := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}

	if actual.Val != expected.Val {
		t.Errorf("expected=%v, actual=%v", expected, actual)
	}
	if actual.Next.Val != expected.Next.Val {
		t.Errorf("expected=%v, actual=%v", expected.Next, actual.Next)
	}
	if actual.Next.Next.Val != expected.Next.Next.Val {
		t.Errorf("expected=%v, actual=%v", expected.Next.Next, actual.Next.Next)
	}
	if actual.Next.Next.Next.Val != expected.Next.Next.Next.Val {
		t.Errorf("expected=%v, actual=%v", expected.Next.Next.Next, actual.Next.Next.Next)
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	testCase1 := &ListNode{
		Val:  0,
		Next: nil,
	}
	testCase2 := &ListNode{
		Val:  0,
		Next: nil,
	}
	expected := &ListNode{
		Val:  0,
		Next: nil,
	}

	actual := addTwoNumbers(testCase1, testCase2)
	if actual.Val != expected.Val {
		t.Errorf("expected=%v, actual=%v", expected, actual)
	}
}
