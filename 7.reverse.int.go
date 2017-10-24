package leetcode

import (
	"math"
)

func reverse(x int) int {
	var digits []int
	var ret int64
	for {
		digits = append(digits, x%10)
		if x/10 != 0 {
			x = x / 10
		} else {
			break
		}
	}
	length := len(digits)
	for i := 0; i < length; i++ {
		ad := int(math.Pow10(length - 1 - i))
		ret += int64(digits[i] * ad)
	}
	if ret > int64(0x7fffffff) || ret < int64(-0x7fffffff) {
		return 0
	} else {
		return int(ret)
	}
}
