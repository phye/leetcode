package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	buf := make([]int, n)
	nums1 = append(nums1, buf...)
	for k := m + n - 1; k >= 0; k-- {
		if n == 0 {
			break
		}
		if nums2[n-1] > nums1[m-1] {
			nums1[k] = nums2[n-1]
			n -= 1
		} else {
			nums1[k] = nums1[m-1]
			m -= 1
		}
	}
	return nums1
}

func main() {
	fmt.Printf("%v\n", merge([]int{0, 3, 5, 6, 9}, 5, []int{1, 2, 4, 7}, 4))
}
