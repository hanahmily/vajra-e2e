package sm

import "testing"

// TestE2E simulates a real use case: reducing a slice of ints to its maximum by
// folding Max over the elements. RED until sm.go provides Max.
func TestE2E(t *testing.T) {
	nums := []int{3, 7, 2, 8, 5, 8, 1}

	got := nums[0]
	for _, n := range nums[1:] {
		got = Max(got, n)
	}

	const want = 8
	if got != want {
		t.Fatalf("running max over %v = %d, want %d", nums, got, want)
	}
}
