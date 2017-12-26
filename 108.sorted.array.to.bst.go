package main

import "fmt"

type (
	TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
)

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := (0 + len(nums) - 1) / 2
	root := &TreeNode{nums[mid], nil, nil}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

func main() {
	arr := []int{-10, -3, 0, 5, 9}
	root := sortedArrayToBST(arr)
	fmt.Printf("%v\n", root.Val)
}
