package main

import "fmt"

func UniquePaths(m, n int) int {
	P := make([][]int, m)
	for i := 0; i < m; i++ {
		P[i] = make([]int, n)
		P[i][0] = 1
	}
	for j := 0; j < n; j++ {
		P[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			P[i][j] = P[i-1][j] + P[i][j-1]
		}
	}
	return P[m-1][n-1]
}

func main() {
	fmt.Printf("%v\n", UniquePaths(3, 3))
}
