package mathx_test

import (
	"testing"

	"mathxlive"
)

// TestE2E simulates a real use case: a caller clamps a slice of raw values
// into a fixed [lo, hi] window (e.g. normalizing sensor readings into a
// display range) and expects each element bounded accordingly.
// It is RED until mathx.go provides the implementation.
func TestE2E(t *testing.T) {
	const lo, hi = 0, 10

	input := []int{-5, -1, 0, 3, 5, 10, 11, 99}
	want := []int{0, 0, 0, 3, 5, 10, 10, 10}

	got := make([]int, len(input))
	for i, v := range input {
		got[i] = mathx.Clamp(v, lo, hi)
	}

	if len(got) != len(want) {
		t.Fatalf("clamped slice length = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("clamped[%d] = %d, want %d (input %d)", i, got[i], want[i], input[i])
		}
	}
}
