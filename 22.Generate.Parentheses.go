package main

import "fmt"

func generateParenthesis(n int) []string {
	mem := [][]string{[]string{"("}}
	for i := 1; i < 2*n; i++ {
		mem = append(mem, []string{})
		for j := 0; j < len(mem[i-1]); j++ {
			lcnt := 0
			for _, c := range mem[i-1][j] {
				if c == '(' {
					lcnt += 1
				}
			}
			if lcnt > len(mem[i-1][j])-lcnt {
				if lcnt < n {
					mem[i] = append(mem[i], mem[i-1][j]+"(")
					mem[i] = append(mem[i], mem[i-1][j]+")")
				} else {
					mem[i] = append(mem[i], mem[i-1][j]+")")
				}
			} else if lcnt == len(mem[i-1][j])-lcnt {
				mem[i] = append(mem[i], mem[i-1][j]+"(")
			} else {
				// Should never occur
			}
		}
	}
	return mem[2*n-1]
}

func main() {
	fmt.Printf("%v", generateParenthesis(4))
}
