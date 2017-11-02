package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{1}
		} else {
			return []int{-1, -1}
		}
	}
	start := 0
	end := len(nums) - 1
	for {
		mid := (end + start) / 2
		if nums[mid] == target {
			var l, r int
			for l = mid; l >= 0 && nums[l] == nums[mid]; l-- {
			}
			for r = mid; r < len(nums) && nums[r] == nums[mid]; r++ {
			}
			return []int{l + 1, r - 1}
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
		if start > end {
			return []int{-1, -1}
		}
	}
}

func main() {
	ret := searchRange([]int{2, 2}, 2)
	fmt.Printf("%v\n", ret)
}
