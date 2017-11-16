package main

import "fmt"

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if _exist(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func _exist(board [][]byte, word string, i int, j int, k int) bool {
	if k >= len(word) {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}
	if board[i][j] == word[k] {
		k += 1
		c := board[i][j]
		board[i][j] = '#'
		res := _exist(board, word, i-1, j, k) || _exist(board, word, i+1, j, k) || _exist(board, word, i, j-1, k) || _exist(board, word, i, j+1, k)
		board[i][j] = c
		return res
	}
	return false
}

func main() {
	board := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}
	fmt.Printf("%v\n", exist([][]byte{[]byte{'a', 'a'}}, "aaa"))
	fmt.Printf("%v\n", exist(board, "ABCCED"))
}
