package leetcode

import "testing"

func TestReverseWords(t *testing.T) {
	var tests = []struct { // Test table
		in1 string
		out string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world!  ", "world! hello"},
		{"a good   example", "example good a"},
	}

	for i, item := range tests {
		ret := reverseWords(item.in1)
		if ret != item.out {
			t.Errorf("%d. %v => %v, wanted: %v", i, item.in1, ret, item.out)
		}
	}
}
