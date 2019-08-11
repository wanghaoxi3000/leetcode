package leetcode

import "testing"

func TestLongestCommonPrefix1(t *testing.T) {
	var tests = []struct { // Test table
		in  []string
		out string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"", "b"}, ""},
		{[]string{"aca", "cba"}, ""},
	}

	for i, item := range tests {
		ret := longestCommonPrefix1(item.in)
		if ret != item.out {
			t.Errorf("%d. %v => %v, wanted: %v", i, item.in, ret, item.out)
		}
	}
}

func TestLongestCommonPrefix2(t *testing.T) {
	var tests = []struct { // Test table
		in  []string
		out string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"", "b"}, ""},
		{[]string{"aca", "cba"}, ""},
	}

	for i, item := range tests {
		ret := longestCommonPrefix2(item.in)
		if ret != item.out {
			t.Errorf("%d. %v => %v, wanted: %q", i, item.in, ret, item.out)
		}
	}
}
