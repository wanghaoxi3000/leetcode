package leetcode

import "testing"

func TestMultiply(t *testing.T) {
	var tests = []struct { // Test table
		in1 string
		in2 string
		out string
	}{
		{"2", "3", "6"},
		{"123", "456", "56088"},
		{"9", "99", "891"},
	}

	for i, item := range tests {
		ret := multiply(item.in1, item.in2)
		if ret != item.out {
			t.Errorf("%d. %v %v => %v, wanted: %v", i, item.in1, item.in2, ret, item.out)
		}
	}
}
