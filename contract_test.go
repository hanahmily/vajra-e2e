package calc_test

import (
	"testing"

	"calclive"
)

// TestAdd is the contract unit test for the calc package. It pins the approved
// public interface: calc.Add(a, b int) int must return a + b.
//
// This test is RED at assign time because calc.go does not yet exist, so the
// package has no exported Add symbol and the suite fails to compile. The coder
// makes it GREEN by adding calc.go — without editing this protected file.
func TestAdd(t *testing.T) {
	if got := calc.Add(2, 3); got != 5 {
		t.Fatalf("calc.Add(2, 3) = %d, want 5", got)
	}
}
