package main

import "fmt"

func trap(height []int) int {
	ret := 0
	if len(height) < 3 {
		return 0
	}
	stack := []int{height[0]}
	for i := 1; i < len(height); i++ {
		if height[i] < stack[0] {
			stack = append(stack, height[i])
		} else {
			for j := 1; j < len(stack); j++ {
				ret += stack[0] - stack[j]
			}
			stack = []int{height[i]}
		}
	}
	// No water can be trapped for stack of less than 2 elements
	if len(stack) > 2 {
		right := len(stack) - 1
		for j := right - 1; j > 0; j-- {
			if stack[j] < stack[right] {
				if j == 1 {
					for k := right - 1; k > 0; k-- {
						ret += stack[right] - stack[k]
					}
				}
			} else {
				for k := right - 1; k > j; k-- {
					ret += stack[right] - stack[k]
				}
				right = j
			}
		}
	}
	return ret
}

func main() {
	fmt.Printf("%v\n", trap([]int{5, 4, 1, 2}))
	fmt.Printf("%v\n", trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Printf("%v\n", trap([]int{6, 8, 5, 0, 0, 6, 5}))
	fmt.Printf("%v\n", trap([]int{4, 2, 3}))
	fmt.Printf("%v\n", trap([]int{}))
}
