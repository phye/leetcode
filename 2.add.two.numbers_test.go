package leetcode

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{2, nil}
	l1.Next = &ListNode{4, nil}
	l1.Next.Next = &ListNode{3, nil}

	l2 := &ListNode{5, nil}
	l2.Next = &ListNode{6, nil}
	l2.Next.Next = &ListNode{4, nil}

	ret := addTwoNumbers(l1, l2)
	printNumber(ret)

	if ret.Val != 7 || ret.Next.Val != 0 || ret.Next.Next.Val != 8 {
		t.Fatalf("Bug exists in addTwoNumbers")
	}
}
