package main

import "fmt"

func LongestPalindromicString(s string) string {
	dp := make([][]bool, len(s))
	longest := ""
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = (j-i < 3) || dp[i+1][j-1]
			}
			if dp[i][j] && j-i+1 > len(longest) {
				longest = s[i : j+1]
			}
		}
	}
	return longest
}

func main() {
	fmt.Printf("%v\n", LongestPalindromicString("babaddtattarrattatddetartrateedredividerb"))
}
