package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	ret := [][]int{}
	existing := make(map[int]map[int]int)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		k := i + 1
		j := len(nums) - 1
		if _, ok := existing[nums[i]]; !ok {
			existing[nums[i]] = make(map[int]int)
		}
		for {
			if k >= j {
				break
			}
			if nums[i]+nums[k]+nums[j] < 0 {
				k = k + 1
			} else if nums[i]+nums[k]+nums[j] > 0 {
				j = j - 1
			} else {
				if val, ok := existing[nums[i]][nums[k]]; !ok || val != nums[j] {
					ret = append(ret, []int{nums[i], nums[k], nums[j]})
					existing[nums[i]][nums[k]] = nums[j]
				}
				k = k + 1
				j = j - 1
			}
		}
	}
	return ret
}

func main() {
	nums := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	ret := threeSum(nums)
	fmt.Printf("%v \n", ret)
	nums = []int{0, 0, 0}
	ret = threeSum(nums)
	fmt.Printf("%v \n", ret)
}
