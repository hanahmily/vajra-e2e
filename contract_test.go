package strutil

import "testing"

// contract_test.go is a PROTECTED contract file. It pins the public API of the
// strutil package: Reverse(s string) string returns s with its runes in
// reverse order (rune-correct, not byte-correct).
//
// This suite is intentionally RED at assign time: strutil.go does not exist
// yet, so the package fails to compile and `go test ./...` is non-zero. The
// coder makes it GREEN by adding strutil.go, without editing this file.

func TestReverseASCII(t *testing.T) {
	const in, want = "abc", "cba"
	if got := Reverse(in); got != want {
		t.Fatalf("Reverse(%q) = %q, want %q", in, got, want)
	}
}

func TestReverseUnicode(t *testing.T) {
	// Rune-correct reversal: multi-byte runes must be reversed whole, never
	// split into their constituent bytes.
	const in, want = "Hello, 世界", "界世 ,olleH"
	if got := Reverse(in); got != want {
		t.Fatalf("Reverse(%q) = %q, want %q", in, got, want)
	}
}
