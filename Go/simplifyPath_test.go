package leetcode

import "testing"

func TestSimplifyPath(t *testing.T) {
	var tests = []struct { // Test table
		in1 string
		out string
	}{
		{"/home/", "/home"},
		{"/../", "/"},
		{"/home//foo/", "/home/foo"},
		{"/a/./b/../../c/", "/c"},
		{"/a/../../b/../c//.//", "/c"},
		{"/a/../../b/../c//.//", "/c"},
		{"/a//b////c/d//././/..", "/a/b/c"},
	}

	for i, item := range tests {
		ret := simplifyPath1(item.in1)
		if ret != item.out {
			t.Errorf("simplifyPath1 %d. %v => %v, wanted: %v", i, item.in1, ret, item.out)
		}
	}

	for i, item := range tests {
		ret := simplifyPath2(item.in1)
		if ret != item.out {
			t.Errorf("simplifyPath2 %d. %v => %v, wanted: %v", i, item.in1, ret, item.out)
		}
	}
}
