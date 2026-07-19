package calc_test

import (
	"testing"

	"calclive"
)

// TestE2E simulates a real use case of the calc package: summing a slice of
// integers by folding calc.Add over the elements, starting from 0. This is the
// canonical way a caller builds a running total on top of the public interface.
//
// Like the contract test, this is RED at assign time (calc.go is absent) and
// turns GREEN once the coder implements calc.Add.
func TestE2E(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	sum := 0
	for _, n := range nums {
		sum = calc.Add(sum, n)
	}

	const want = 15 // 1 + 2 + 3 + 4 + 5
	if sum != want {
		t.Fatalf("folding calc.Add over %v = %d, want %d", nums, sum, want)
	}
}
