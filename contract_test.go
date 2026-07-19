package mathx

import "testing"

// TestClampContract asserts the approved public contract for Clamp.
// It is RED until mathx.go provides the implementation.
func TestClampContract(t *testing.T) {
	cases := []struct {
		name           string
		v, lo, hi, want int
	}{
		{"within range", 5, 0, 10, 5},
		{"below lo", -1, 0, 10, 0},
		{"above hi", 99, 0, 10, 10},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Clamp(tc.v, tc.lo, tc.hi); got != tc.want {
				t.Errorf("Clamp(%d, %d, %d) = %d, want %d", tc.v, tc.lo, tc.hi, got, tc.want)
			}
		})
	}
}
