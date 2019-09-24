package string

import (
	"testing"
)

func TestEvalRPN(t *testing.T) {
	var tests = []struct { // Test table
		in1 []string
		out int
	}{
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}

	for i, item := range tests {
		ret := evalRPN(item.in1)
		if ret != item.out {
			t.Errorf("%d. %v => %v, wanted: %v", i, item.in1, ret, item.out)
		}
	}
}
