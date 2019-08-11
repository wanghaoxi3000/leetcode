package leetcode

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	var tests = []struct { // Test table
		in  string
		out int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for i, item := range tests {
		ret := lengthOfLongestSubstring(item.in)
		if ret != item.out {
			t.Errorf("%d. %q => %d, wanted: %d", i, item.in, ret, item.out)
		}
	}
}
