package calc_test

import (
	"testing"

	"calclive"
)

// TestAdd is the contract unit test for the calc package. It pins the approved
// public interface: calc.Add(a, b int) int must return a + b.
//
// This test is RED at assign time because calc.go does not yet exist, so the
// calclive package has no buildable Go files and the suite fails to compile.
// The coder makes it GREEN by adding calc.go — without editing this protected
// file.
func TestAdd(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{"positive", 2, 3, 5},
		{"with zero", 7, 0, 7},
		{"negative", -4, -6, -10},
		{"mixed sign", -8, 3, -5},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := calc.Add(tc.a, tc.b); got != tc.want {
				t.Fatalf("calc.Add(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
			// Add is commutative for ints.
			if got := calc.Add(tc.b, tc.a); got != tc.want {
				t.Fatalf("calc.Add(%d, %d) = %d, want %d (commutativity)", tc.b, tc.a, got, tc.want)
			}
		})
	}
}
