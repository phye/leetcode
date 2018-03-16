package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) == 0 {
		return res
	}

	used := make([]bool, len(nums))
	cur := make([]int, 0)
	sort.Ints(nums)
	dfs(nums, used, &cur, &res)

	return res
}

func dfs(nums []int, used []bool, cur *[]int, res *[][]int) {
	if len(*cur) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, *cur)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
			continue
		}
		used[i] = true
		*cur = append(*cur, nums[i])
		dfs(nums, used, cur, res)
		used[i] = false
		*cur = (*cur)[:len(*cur)-1]
	}
}

func main() {
	nums := []int{0, 0, 0, 1, 9}
	ret := permuteUnique(nums)
	for _, v := range ret {
		fmt.Printf("%v\n", v)
	}
}
