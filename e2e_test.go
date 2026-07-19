package strutil

import "testing"

// e2e_test.go is a PROTECTED end-to-end file. It exercises the feature the way
// a real caller would: reversing a string and then reversing the result must
// yield the original string (Reverse is its own inverse / an involution).
//
// Like contract_test.go this is RED at assign time and turns GREEN once the
// coder adds strutil.go. The coder must not edit this file.

func TestE2E(t *testing.T) {
	inputs := []string{
		"abc",
		"Hello, 世界",
		"",
		"a",
		"racecar",
		"naïve café",
	}
	for _, in := range inputs {
		once := Reverse(in)
		twice := Reverse(once)
		if twice != in {
			t.Fatalf("round-trip failed: Reverse(Reverse(%q)) = %q (single reverse = %q), want %q",
				in, twice, once, in)
		}
	}
}
