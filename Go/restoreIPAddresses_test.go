package leetcode

import (
	"reflect"
	"testing"
)

func TestRestoreIPAddresses(t *testing.T) {
	var tests = []struct { // Test table
		in1 string
		out []string
	}{
		{"25525511135", []string{"255.255.11.135", "255.255.111.35"}},
		{"0000", []string{"0.0.0.0"}},
	}

	for i, item := range tests {
		ret := restoreIPAddresses(item.in1)
		if !reflect.DeepEqual(ret, item.out) {
			t.Errorf("%d. %v => %v, wanted: %v", i, item.in1, ret, item.out)
		}
	}
}
