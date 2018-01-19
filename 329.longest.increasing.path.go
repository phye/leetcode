package main

import "fmt"

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	mval := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			longestIncreasingPathAux(matrix, dp, i, j, &mval, m, n)
		}
	}
	return mval
}

func longestIncreasingPathAux(matrix [][]int, dp [][]int, i, j int, pmax *int, m, n int) {
	if dp[i][j] > 0 {
		return
	}
	tmp := 1
	if isValid(m, n, i-1, j) && matrix[i-1][j] < matrix[i][j] {
		longestIncreasingPathAux(matrix, dp, i-1, j, pmax, m, n)
		tmp = max(tmp, dp[i-1][j]+1)
	}
	if isValid(m, n, i+1, j) && matrix[i+1][j] < matrix[i][j] {
		longestIncreasingPathAux(matrix, dp, i+1, j, pmax, m, n)
		tmp = max(tmp, dp[i+1][j]+1)
	}
	if isValid(m, n, i, j-1) && matrix[i][j-1] < matrix[i][j] {
		longestIncreasingPathAux(matrix, dp, i, j-1, pmax, m, n)
		tmp = max(tmp, dp[i][j-1]+1)
	}
	if isValid(m, n, i, j+1) && matrix[i][j+1] < matrix[i][j] {
		longestIncreasingPathAux(matrix, dp, i, j+1, pmax, m, n)
		tmp = max(tmp, dp[i][j+1]+1)
	}
	*pmax = max(tmp, *pmax)
	dp[i][j] = tmp
}

func isValid(m, n, i, j int) bool {
	if i < 0 || j < 0 {
		return false
	}
	if i >= m || j >= n {
		return false
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	matrix := [][]int{
		[]int{9, 9, 4},
		[]int{6, 6, 8},
		[]int{2, 1, 1},
	}
	res := longestIncreasingPath(matrix)
	fmt.Printf("%v\n", res)
}
