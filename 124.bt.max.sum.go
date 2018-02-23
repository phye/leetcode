package main

import (
	"fmt"
	base "github.com/phye/algobase"
)

// A path in this problem is: 0 or more steps up, reaching one highest node in the path, then 0 or
// more steps down. Once the path goes down, it cannot go up again.

func maxSumPath(root *base.TreeNode) (int, []int) {
	ret := -0x7fffffff
	path := make([]int, 0)
	maxSumPathAux(root, &ret, &path)
	return ret, path
}

func maxSumPathAux(root *base.TreeNode, pmax *int, path *[]int) int {
	if root == nil {
		return 0
	}
	l := maxSumPathAux(root.Left, pmax, path)
	if l < 0 {
		l = 0
	}
	r := maxSumPathAux(root.Right, pmax, path)
	if r < 0 {
		r = 0
	}
	// TODO: Add finer control over path generation
	if *pmax < l+r+root.Val {
		*path = append(*path, root.Val)
	}
	*pmax = max(*pmax, l+r+root.Val)
	// Return maximum path sum ending at root
	return root.Val + max(l, r)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	bt := base.BuildTree([]int{3, 9, 20, -5, -4, -15, 7})
	sum, path := maxSumPath(bt)
	fmt.Printf("%v %v\n", sum, path)
}
