package radix

import (
	"testing"
)

func TestRadixSort(t *testing.T) {
	test_cases := []struct {
		input    []int
		expected []int
	}{
		{[]int{100, 3, 8}, []int{3, 8, 100}},
		{[]int{100, 3, 8, -9}, []int{-9, 3, 8, 100}},
		{[]int{100, 3, 0, 8, -9}, []int{-9, 0, 3, 8, 100}},
		{[]int{-9999, 0, -999, 999, 999, 998, 100, 1000, 10000}, []int{-9999, -999, 0, 100, 998, 999, 999, 1000, 10000}},
		{[]int{}, []int{}},
		{[]int{1000}, []int{1000}},
		{[]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}, []int{-5467984, -784, 0, 0, 42, 59, 74, 238, 905, 959, 7586, 7586, 9845}},
	}

	for _, tc := range test_cases {
		Ints(tc.input)

		verifySlice(tc.input, tc.expected, t)
	}
}

func verifySlice(actual []int, expected []int, t *testing.T) {
	if len(actual) != len(expected) {
		t.Fatalf("Expected actual %v to have the same length as expected %v", actual, expected)
	}

	t.Logf("actual %v, expected %v", actual, expected)

	for i, val := range expected {
		actual_val := actual[i]

		if val != actual_val {
			t.Errorf("Unexpected value in position %d", i)
		}
	}
}
