package main

import (
	"fmt"
	base "github.com/phye/algobase"
)

func ValidateBST(root *base.TreeNode) bool {
	return isBSTValid(root, nil, nil)
}

func isBSTValid(root, min, max *base.TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val < min.Val {
		return false
	}
	if max != nil && root.Val > max.Val {
		return false
	}
	return isBSTValid(root.Left, min, root) && isBSTValid(root.Right, root, max)
}

func main() {
	bt := base.BuildTree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Printf("%v\n", ValidateBST(bt))

	bst := base.BuildTree([]int{5, 2, 8, 1, 3, 6, 9})
	fmt.Printf("%v\n", ValidateBST(bst))
}
