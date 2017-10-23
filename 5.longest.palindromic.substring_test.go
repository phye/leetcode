package leetcode

import (
	"testing"
)

func TestFindLongestPalindromic(t *testing.T) {
	input := "babad"
	result := FindLongestPalindromic(input)
	if result != "bab" {
		t.Fatalf("Bug in palindromic. Expected: %s, actual: %s", "bad", result)
	}
}
