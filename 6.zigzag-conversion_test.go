package leetcode

import (
	"testing"
)

func TestZigZagConvert(t *testing.T) {
	ret := convert("PAYPALISHIRING", 3)

	if ret != "PAHNAPLSIIGYIR" {
		t.Fatalf("Bug in zig zag convert. Expected: %s, Acutal: %s", "PAHNAPLSIIGYIR", ret)
	}

	ret = convert("AB", 2)
	if ret != "AB" {
		t.Fatalf("Expected: %v, Actual: %v", "AB", ret)
	}
}
