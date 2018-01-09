package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var nh, cur, fast *ListNode
	for fast = head; fast != nil && fast.Next != nil; {
		fast = fast.Next.Next
		cur = head
		head = head.Next
		cur.Next = nh
		nh = cur
	}
	var l, r *ListNode
	if fast != nil {
		l, r = nh, head.Next
	} else {
		l, r = nh, head
	}
	for ; l != nil && r != nil; l, r = l.Next, r.Next {
		if l.Val != r.Val {
			return false
		}
	}
	return true
}

func main() {
	l := &ListNode{1, nil}
	l.Next = &ListNode{1, nil}
	fmt.Printf("%v\n", isPalindrome(l))
}
