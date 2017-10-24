package leetcode

import (
	"testing"
)

func TestReverse(t *testing.T) {
	ret := reverse(1354)
	if ret != 4531 {
		t.Fatalf("Got %v, expected: %v", ret, 4531)
	}
}
