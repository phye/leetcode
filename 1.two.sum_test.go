package leetcode

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	input := []int{2, 7, 11, 15}
	target := 9
	ret := twoSum(input, target)
	fmt.Printf("Indexes returned are %v\n", ret)
	if len(ret) != 2 {
		t.Fatalf("Two sum failed")
	}
}
