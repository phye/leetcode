package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool, 0)
	for _, word := range wordDict {
		m[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	// Note: i-1 denotes the 0 based index in s
	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] {
				if m[s[j:i]] {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[len(s)]
}

func main() {
	str := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	//str := "aaaaaaaaaab"
	dict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	fmt.Printf("%v\n", wordBreak(str, dict))
}
