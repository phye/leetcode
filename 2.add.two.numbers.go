package leetcode

import (
	"fmt"
)

type (
	ListNode struct {
		Val  int
		Next *ListNode
	}
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ret *ListNode
	var tail *ListNode
	adv := 0
	for {
		val := 0
		cur := &ListNode{}
		if l1 != nil {
			val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val = val + l2.Val
			l2 = l2.Next
		}
		val = val + adv
		//fmt.Printf("%v\n", val)
		if val >= 10 {
			cur.Val = val - 10
			adv = 1
		} else {
			cur.Val = val
			adv = 0
		}
		if ret == nil {
			ret = cur
			tail = ret
		} else {
			tail.Next = cur
			tail = tail.Next
		}
		if l1 == nil && l2 == nil && adv == 0 {
			return ret
		}
	}
}

func printNumber(ln *ListNode) {
	for cur := ln; cur != nil; cur = cur.Next {
		fmt.Printf("%v ", cur.Val)
	}
	fmt.Printf("\n")
}
