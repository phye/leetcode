package main

import "fmt"

func permute(nums []int) [][]int {
	ret := _permute(nums)
	return ret
}

func _permute(nums []int) [][]int {
	ret := make([][]int, 0)
	if len(nums) == 0 {
		return ret
	} else if len(nums) == 1 {
		ret = append(ret, []int{})
		ret[0] = append(ret[0], nums[0])
		return ret
	} else {
		for i := 0; i < len(nums); i++ {
			_nums := []int{}
			_nums = append(_nums, nums[:i]...)
			_nums = append(_nums, nums[i+1:]...)
			subs := _permute(_nums)
			for j := 0; j < len(subs); j++ {
				ret = append(ret, append(subs[j], nums[i]))
			}
		}
		return ret
	}
}

func main() {
	fmt.Printf("%v\n", permute([]int{1, 2, 3}))
}
