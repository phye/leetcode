package main

import (
	"fmt"
)

func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}
	stack := make([]int, 0)
	num := 0
	var op byte
	op = '+'
	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			num = num*10 + int(s[i]-'0')
		}
		if (!isDigit(s[i]) && s[i] != ' ') || i == len(s)-1 {
			switch op {
			case '-':
				stack = append(stack, -num)
			case '+':
				stack = append(stack, num)
			case '*':
				stack = append(stack[:len(stack)-1], stack[len(stack)-1]*num)
			case '/':
				stack = append(stack[:len(stack)-1], stack[len(stack)-1]/num)
			}
			op = s[i]
			num = 0
		}
	}
	ret := 0
	for _, v := range stack {
		ret += v
	}
	return ret
}

func isDigit(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Printf("%v\n", calculate("100000000/1/2/3/4/5/6/7/8/9/10"))
}
