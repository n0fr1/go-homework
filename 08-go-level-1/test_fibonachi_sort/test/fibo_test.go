package test

import (
	"reflect"
	"testProject/test_fibonachi_sort/fibo"
	"testing"
)

func TestFibo(t *testing.T) {

	cases := []struct {
		name string
		in   uint64
		out  uint64
	}{
		{
			name: "one",
			in:   1,
			out:  0,
		},

		{
			name: "two",
			in:   2,
			out:  1,
		},

		{
			name: "three",
			in:   3,
			out:  1,
		},

		{
			name: "four",
			in:   4,
			out:  2,
		},

		{
			name: "five",
			in:   5,
			out:  3,
		},

		{
			name: "six",
			in:   6,
			out:  5,
		},
	}

	for _, tc := range cases {

		t.Run(tc.name, func(t *testing.T) {
			result := fibo.Count(tc.in)

			if !reflect.DeepEqual(tc.out, result) {
				t.Error("Expected ", tc.out, " got ", result)
			}

		})

	}

}
