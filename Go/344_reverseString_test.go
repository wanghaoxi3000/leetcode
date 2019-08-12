package leetcode

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	var tests = []struct {
		in  []byte
		out []byte
	}{
		{
			[]byte{'h', 'e', 'l', 'l', 'o'},
			[]byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			[]byte{'H', 'a', 'n', 'n', 'a', 'h'},
			[]byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}

	for i, item := range tests {
		reverseString(item.in)
		if !reflect.DeepEqual(item.in, item.out) {
			t.Errorf("%d. %v, wanted: %v", i, string(item.in), string(item.out))
		}
	}
}
