package main

import "fmt"

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	P := make([]int, len(s))
	_numDecodings(s, len(s)-1, P)
	return P[len(s)-1]
}

func _numDecodings(s string, i int, P []int) {
	if i < 0 || P[i] > 0 {
		return
	}
	if i == 0 {
		P[i] = 1
		return
	} else {
		if s[i] == '0' {
			if s[i-1] != '1' && s[i-1] != '2' {
				P[i] = 0
				return
			} else {
				_numDecodings(s, i-1, P)
				if i >= 2 && P[i-1] != P[i-2] {
					P[i] = P[i-2]
					return
				} else {
					P[i] = P[i-1]
					return
				}
			}
		}
		_numDecodings(s, i-1, P)

		if s[i-1] == '1' || (s[i-1] == '2' && s[i] < '7') {
			if i >= 2 {
				P[i] = P[i-1] + P[i-2]
			} else {
				P[i] = P[i-1] + 1
			}
		} else {
			P[i] = P[i-1]
		}
	}
	return
}

func main() {
	fmt.Println("vim-go")
}
