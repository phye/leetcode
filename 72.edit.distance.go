package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = MinVal([]int{dp[i-1][j], dp[i][j-1], dp[i-1][j-1]})
			}
		}
	}
	return dp[m][n]
}

func MinVal(nums []int) int {
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}

func main() {
	fmt.Println("vim-go")

	fmt.Printf("%v\n", minDistance("a", "a"))
	//fmt.Printf("%v\n", EditDistance("hello", "helix"))
}
