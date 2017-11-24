package main

import "fmt"

func isMatch(T string, P string) bool {
	m := len(T)
	n := len(P)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if P[j-1] != '*' {
				//if i == 0 {
				//	continue
				//}
				if i > 0 && (P[j-1] == '.' || T[i-1] == P[j-1]) {
					dp[i][j] = dp[i-1][j-1]
				}
				// Otherwise, dp[i][j] must be false
			} else {
				if j < 2 {
					continue
				}
				// Zero match
				if dp[i][j-2] {
					dp[i][j] = true
					continue
				}

				// At least one match
				if i > 0 && (P[j-2] == '.' || T[i-1] == P[j-2]) {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
		fmt.Printf("%v\n", dp[i])
	}
	return dp[m][n]
}

func main() {
	fmt.Printf("%v\n", isMatch("aa", "a"))
	//fmt.Printf("%v\n", isMatch("aaa", "ab*a"))
	//fmt.Printf("%v\n", isMatch("ab", ".*"))
	//fmt.Printf("%v\n", isMatch("aab", "c*a*b"))
	//fmt.Printf("%v\n", isMatch("aa", "a"))
	//fmt.Printf("%v\n", isMatch("aa", "aa"))
	//fmt.Printf("%v\n", isMatch("aaa", "aa"))
	//fmt.Printf("%v\n", isMatch("aa", "a*"))
	//fmt.Printf("%v\n", isMatch("aa", ".*"))
}
