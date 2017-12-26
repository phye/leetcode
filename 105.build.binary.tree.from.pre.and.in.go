package main

import (
	"fmt"
)

type (
	TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	m := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		m[inorder[i]] = i
	}
	root := &TreeNode{preorder[0], nil, nil}
	idx := m[root.Val]
	root.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	return root
}

func main() {
	bt := buildTree(
		[]int{33, 27, 29, 3, 20, 30, 8, 39, 28},
		[]int{29, 27, 20, 3, 33, 8, 39, 30, 28})
	fmt.Printf("%v\n", bt.Val)
}
