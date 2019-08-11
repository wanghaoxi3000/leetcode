package leetcode

import "testing"

func TestCheckInclusion1(t *testing.T) {
	var tests = []struct { // Test table
		in1 string
		in2 string
		out bool
	}{
		{"a", "b", false},
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
		{"adc", "dcda", true},
		{"dinitrophenylhydrazine", "acetylphenylhydrazine", false},
	}

	for i, item := range tests {
		ret := checkInclusion1(item.in1, item.in2)
		if ret != item.out {
			t.Errorf("%d. %v %v => %v, wanted: %v", i, item.in1, item.in2, ret, item.out)
		}
	}
}

func TestCheckInclusion2(t *testing.T) {
	var tests = []struct { // Test table
		in1 string
		in2 string
		out bool
	}{
		{"a", "b", false},
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
		{"adc", "dcda", true},
		{"dinitrophenylhydrazine", "acetylphenylhydrazine", false},
	}

	for i, item := range tests {
		ret := checkInclusion2(item.in1, item.in2)
		if ret != item.out {
			t.Errorf("%d. %v %v => %v, wanted: %v", i, item.in1, item.in2, ret, item.out)
		}
	}
}
