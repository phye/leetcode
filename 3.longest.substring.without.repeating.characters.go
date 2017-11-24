package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	longest := 0
	lm := make(map[byte]int)
	// last unrepeated char
	last := 0
	for i := 0; i < len(s); i++ {
		if prev, ok := lm[s[i]]; ok {
			if last < prev+1 {
				last = prev + 1
			}
		}
		if i-last+1 > longest {
			longest = i - last + 1
		}
		lm[s[i]] = i
	}
	return longest
}

func main() {
	fmt.Printf("%v\n", lengthOfLongestSubstring("abba"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("pwwkew"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("abcdef"))
}
