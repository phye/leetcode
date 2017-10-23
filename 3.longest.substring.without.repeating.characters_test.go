package leetcode

import (
	"testing"
)

func TestLengthOfLongestSubString(t *testing.T) {
	input := "abcabcbb"
	if lengthOfLongestSubString(input) != 3 {
		t.Fatalf("Bug in Longest Sub String")
	}
}
