package string

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	var tests = []struct { // Test table
		in1 string
		out bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}

	for i, item := range tests {
		ret := isValid(item.in1)
		if ret != item.out {
			t.Errorf("%d. %v => %v, wanted: %v", i, item.in1, ret, item.out)
		}
	}
}
