package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	for _, c := range s {
		if c == '(' {
			stack = append(stack, c)
		} else if c == '{' {
			stack = append(stack, c)
		} else if c == '[' {
			stack = append(stack, c)
		} else if c == ')' {
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else if c == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else if c == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return false
	} else {
		return true
	}
}

func main() {
	fmt.Printf("%v", isValid("()"))
}
