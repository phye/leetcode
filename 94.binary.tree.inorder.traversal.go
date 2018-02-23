package main

import (
	"fmt"
	base "github.com/phye/algobase"
)

func InOrder(root *base.TreeNode) []int {
	stack := make([]*base.TreeNode, 0)
	ret := make([]int, 0)
	cur := root
	for {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else if len(stack) > 0 {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret = append(ret, cur.Val)
			cur = cur.Right
		} else {
			return ret
		}
	}
}

func main() {
	bt := base.BuildTree([]int{1, 2, 3, 4, 5, 6, 7})
	ret := InOrder(bt)
	fmt.Printf("%v\n", ret)
}
