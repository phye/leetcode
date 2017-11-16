package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	l := 2*n + 2*(m-2)
	if m == 1 {
		return matrix[0]
	}
	if n == 1 {
		ret := make([]int, m)
		for i := 0; i < m; i++ {
			ret[i] = matrix[i][0]
		}
		return ret
	}
	ret := make([]int, l)
	for k := 0; k < l; k++ {
		var i, j int
		if k < n {
			i = 0
			j = k
		} else if k < n+m-1 {
			i = k - n + 1
			j = n - 1
		} else if k < n*2-1+m-2 {
			i = m - 1
			j = n - 2 - (k - n - m + 1)
		} else {
			i = l - k
			j = 0
		}
		if i < 0 || j < 0 {
			break
		}
		ret[k] = matrix[i][j]
	}
	if m > 1 {
		for i := 1; i < m-1; i++ {
			matrix[i] = matrix[i][1 : n-1]
		}
		inner := spiralOrder(matrix[1 : m-1])
		ret = append(ret, inner...)
	}
	return ret
}

func main() {
	fmt.Printf("%v\n", spiralOrder([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}))
}
