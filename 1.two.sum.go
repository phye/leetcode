package leetcode

func twoSum(nums []int, target int) []int {
	/* Brute Force
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
	*/

	lm := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		lm[nums[i]] = i
		compl := target - nums[i]
		if idx, ok := lm[compl]; ok {
			return []int{idx, i}
		}
	}
	return []int{}
}
