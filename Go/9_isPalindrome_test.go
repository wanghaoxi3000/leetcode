package leetcode

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		in1 int
		out bool
	}{
		{121, true},
		{-121, false},
		{10, false},
	}

	for i, item := range tests {
		ret := isPalindrome(item.in1)
		if ret != item.out {
			t.Errorf("%d. %v => %v, wanted: %v", i, item.in1, item.out, item.out)
		}
	}
}
