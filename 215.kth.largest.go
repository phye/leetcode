package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	q := partition(nums)
	if q == k-1 {
		return nums[q]
	} else if q < k-1 {
		return findKthLargest(nums[q+1:], k-q-1)
	} else {
		return findKthLargest(nums[:q], k)
	}
}

func partition(nums []int) int {
	r := len(nums) - 1
	x := nums[r]
	i := -1
	for j := 0; j < r; j++ {
		if nums[j] > x {
			tmp := nums[j]
			nums[j] = nums[i+1]
			nums[i+1] = tmp
			i++
		}
	}
	tmp := nums[r]
	nums[r] = nums[i+1]
	nums[i+1] = tmp
	return i + 1
}

func main() {
	fmt.Printf("%v\n", findKthLargest([]int{2, 1}, 1))
}
