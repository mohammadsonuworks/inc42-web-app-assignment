package main

import (
	"testing"
)

// TestAdd tests the Add function with various cases
func TestAdd(t *testing.T) {
	// Define test cases
	testCases := []struct {
		a, b   int
		result int
	}{
		{1, 2, 3},
		{5, 10, 15},
		{-1, -1, -2},
		{-5, 5, 0},
	}

	// Loop through test cases and check if the output matches the expected result
	for _, tc := range testCases {
		got := Add(tc.a, tc.b)
		if got != tc.result {
			t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.result)
		}
	}
}
