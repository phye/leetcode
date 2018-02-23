package main

import (
	"fmt"
	base "github.com/phye/algobase"
)

func buildTree(preorder []int, inorder []int) *base.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	m := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		m[inorder[i]] = i
	}
	root := &base.TreeNode{preorder[0], nil, nil}
	idx := m[root.Val]
	root.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	return root
}

func buildTreeFromInAndPost(inorder []int, postorder []int) *base.TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	m := make(map[int]int)
	for k, v := range inorder {
		m[v] = k
	}
	root := &base.TreeNode{postorder[len(postorder)-1], nil, nil}
	idx := m[root.Val]
	root.Left = buildTreeFromInAndPost(inorder[:idx], postorder[:idx])
	root.Right = buildTreeFromInAndPost(inorder[idx+1:], postorder[idx:len(postorder)-1])
	return root
}

func main() {
	bt := buildTree(
		[]int{33, 27, 29, 3, 20, 30, 8, 39, 28},
		[]int{29, 27, 20, 3, 33, 8, 39, 30, 28})
	fmt.Printf("%v\n", bt.Val)

	bt2 := buildTreeFromInAndPost(
		[]int{9, 3, 15, 20, 7},
		[]int{9, 15, 7, 20, 3},
	)
	fmt.Printf("%v\n", bt2.Val)
}
