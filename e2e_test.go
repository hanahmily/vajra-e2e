package calc_test

import (
	"testing"

	calc "calclive"
)

// TestE2E simulates a real use case of the calc package end to end: a caller
// accumulating a running total by folding calc.Add over a slice of integers,
// exactly as summation code in the wild would. Summing 1..5 must yield 15.
//
// This test is a protected contract file (see .vajra-protected). It is RED
// until calc.Add exists and must not be modified during implementation.
func TestE2E(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	sum := 0
	for _, n := range nums {
		sum = calc.Add(sum, n)
	}

	if want := 15; sum != want {
		t.Fatalf("folding calc.Add over %v = %d, want %d", nums, sum, want)
	}
}
