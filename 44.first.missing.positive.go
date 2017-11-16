package main

import "fmt"

func left(p int) int {
	return 2*p + 1
}

func right(p int) int {
	return 2*p + 2
}

func parent(p int) int {
	return (p - 1) / 2
}

// Given two child tree already in a min heap, preserve a min heap
func minHeapify(A []int, i int, hpsize int) {
	l := left(i)
	r := right(i)
	min := i
	if l < hpsize && A[min] > A[l] {
		min = l
	}
	if r < hpsize && A[min] > A[r] {
		min = r
	}
	if min != i {
		tmp := A[min]
		A[min] = A[i]
		A[i] = tmp
		minHeapify(A, min, hpsize)
	}
}

func buildMinHeap(A []int) {
	hpsize := len(A)
	for i := len(A) / 2; i >= 0; i-- {
		minHeapify(A, i, hpsize)
	}
}

func firstMissingPositive(nums []int) int {
	buildMinHeap(nums)
	smallest := 0
	hpsize := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[0] > smallest {
			if nums[0] == smallest+1 {
				smallest = nums[0]
			} else {
				return smallest + 1
			}
		}
		tmp := nums[0]
		nums[0] = nums[hpsize-1]
		nums[hpsize-1] = tmp
		hpsize = hpsize - 1
		minHeapify(nums, 0, hpsize)
	}
	return smallest + 1
}

func main() {
	fmt.Printf("%v ", firstMissingPositive([]int{3, 4, -1, 1}))
}
