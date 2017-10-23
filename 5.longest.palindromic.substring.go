package leetcode

import (
//"fmt"
)

func FindLongestPalindromic(s string) string {
	src := []byte(s)
	longest := string(src[0])

	l := len(src)
	mem := make(map[int]map[int]bool)
	for i := 0; i < l; i++ {
		mem[i] = make(map[int]bool)
		mem[i][i] = true
	}

	// P(i,j) = P(i+1, j-1) if Si == Sj
	for k := 0; k < l; k++ {
		for j := 1; j <= k; j++ {
			i := k - j
			if src[i] != src[k] {
				mem[i][k] = false
				continue
			}
			if k == i+1 {
				mem[i][k] = true
			} else {
				mem[i][k] = mem[i+1][k-1]
			}
			if mem[i][k] == true && k+1-i > len(longest) {
				longest = string(src[i : k+1])
			}
		}
	}
	return longest
}
