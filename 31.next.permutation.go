package main

import (
	"fmt"
)

func nextPermutation(nums []int) []int {
	n := len(nums)
	ret := make([]int, n)
	var i, j int
	// From the end, find the first one with index such that num[i-1] < num[i]
	for i = n - 1; i > 0 && nums[i-1] >= nums[i]; i-- {
	}
	if i == 0 {
		for k := 0; k < n; k++ {
			ret[k] = nums[n-k-1]
		}
		return ret
	}

	// From the end, find the first one which is larger than nums[i-1]
	for j = n - 1; nums[i-1] >= nums[j]; j-- {
	}

	// swap ret at i-1 and j
	nums[i-1], nums[j] = nums[j], nums[i-1]

	// reversly sort nums[i:] and store in ret
	copy(ret[:i], nums[:i])
	for j := n - 1; j >= i; j-- {
		ret[n-j-1+i] = nums[j]
	}
	return ret
}

func main() {
	nums := []int{1, 2, 3, 4}
	for i := 0; i < 25; i++ {
		fmt.Printf("%v\n", nums)
		nums = nextPermutation(nums)
	}
}
