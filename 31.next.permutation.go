package main

import (
	"fmt"
)

func nextPermutation(nums []int) {
	n := len(nums)
	var i, j int
	// From the end, find the first one with index such that num[i-1] < num[i]
	for i = n - 1; i > 0 && nums[i-1] >= nums[i]; i-- {
	}
	// If no such one were find, nums is already reversly sorted
	if i == 0 {
		for k := 0; k < n/2; k++ {
			nums[k], nums[n-k-1] = nums[n-k-1], nums[k]
		}
		return
	}

	// From the end, find the first one which is larger than nums[i-1]
	for j = n - 1; nums[i-1] >= nums[j]; j-- {
	}

	// swap nums at i-1 and j
	nums[i-1], nums[j] = nums[j], nums[i-1]

	// reversly sort nums[i:] via swap
	for k := 0; k < (n-i)/2; k++ {
		nums[i+k], nums[n-1-k] = nums[n-1-k], nums[i+k]
	}
	return
}

func main() {
	nums := []int{1, 2, 3, 4}
	for i := 0; i < 25; i++ {
		fmt.Printf("%v\n", nums)
		nextPermutation(nums)
	}
}
