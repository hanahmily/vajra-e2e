package calc_test

import (
	"testing"

	calc "calclive"
)

// TestAdd is the RED contract test for the calc package. It pins the single,
// intentionally complete requirement of this milestone: calc.Add(2, 3) == 5.
//
// This test is a protected contract file (see .vajra-protected). It fails to
// build until the coder adds calc.go with func Add(a, b int) int, and it must
// not be modified during implementation.
func TestAdd(t *testing.T) {
	if got := calc.Add(2, 3); got != 5 {
		t.Fatalf("calc.Add(2, 3) = %d, want 5", got)
	}
}
