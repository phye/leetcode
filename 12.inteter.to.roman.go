package leetcode

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	p := 10
	ret := ""
	units := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}
	for {
		res := num / 10
		rem := num % 10
		if res == 0 && rem == 0 {
			return ret
		}
		rdigit := ""
		if rem == 0 {

		} else if rem > 0 && rem < 4 {
			rdigit = strings.Repeat(units[p/10], rem)
		} else if rem == 4 {
			rdigit = units[p/10] + units[p/2]
		} else if rem == 5 {
			rdigit = units[p/2]
		} else if rem < 9 {
			rdigit = units[p/2] + strings.Repeat(units[p/10], rem-5)
		} else if rem == 9 {
			rdigit = units[p/10] + units[p]
		}
		if rem != 0 {
			ret = rdigit + ret
		}
		p = p * 10
		num = num / 10
	}
}
