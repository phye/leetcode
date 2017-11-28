package main

import "fmt"

func letterCombinations(digits string) []string {
	m := map[byte]string{
		'1': "",
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
		'0': " ",
	}
	if len(digits) == 0 {
		return []string{}
	}
	ret := make([]string, 0)
	ret = append(ret, "")
	for i := 0; i < len(digits); i++ {
		chars := m[digits[i]]
		nret := make([]string, 0)
		for k := 0; k < len(ret); k++ {
			for j := 0; j < len(chars); j++ {
				nret = append(nret, ret[k]+string(chars[j]))
			}
		}
		ret = nret
	}
	return ret
}

func main() {
	fmt.Printf("%v\n", letterCombinations("23"))
}
