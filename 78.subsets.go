package main

import "fmt"

func subsets(nums []int) [][]int {
	var ret [][]int
	if len(nums) == 0 {
		ret = make([][]int, 0)
		ret = append(ret, []int{})
		return ret
	}
	ret = subsets(nums[:len(nums)-1])
	olen := len(ret)
	for i := 0; i < olen; i++ {
		e := []int{}
		//e = append(e, ret[i]...)
		e = append(ret[i], nums[len(nums)-1])
		ret = append(ret, e)
	}
	return ret
}

func subsetsIterative(nums []int) [][]int {
	ret := [][]int{}
	ret = append(ret, []int{})
	for i := 0; i < len(nums); i++ {
		olen := len(ret)
		for j := 0; j < olen; j++ {
			e := []int{}
			e = append(e, ret[j]...)
			e = append(e, nums[i])
			ret = append(ret, e)
		}
	}
	return ret
}

func main() {
	//fmt.Printf("%v\n", subsets([]int{1}))
	fmt.Printf("%v\n", subsets([]int{1, 2, 3}))
	fmt.Printf("%v\n", subsetsIterative([]int{1, 2, 3}))
}
