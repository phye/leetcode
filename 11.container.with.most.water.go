package leetcode

func maxArea(height []int) int {
	ret := 0
	left := 0
	right := len(height) - 1
	for {
		if left >= right {
			return ret
		}
		var area int
		if height[left] < height[right] {
			area = (right - left) * height[left]
			left = left + 1
		} else {
			area = (right - left) * height[right]
			right = right - 1
		}
		if area > ret {
			ret = area
		}
	}
}
