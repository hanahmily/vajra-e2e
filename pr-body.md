# Add `strutil.Reverse` (rune-correct string reversal)

Add a new Go package `strutil` at the module root (module `strutillive`)
exposing `func Reverse(s string) string`, which returns the input reversed by
Unicode code point.

## What

- `Reverse(s string) string` — reverses `s` rune-by-rune, so multi-byte UTF-8
  characters are preserved intact rather than split into bytes.

## Why

A small, dependency-free primitive for reversing user-facing text correctly.
Naive byte-level reversal corrupts any non-ASCII string; this reverses whole
runes so `Reverse("Hello, 世界") == "界世 ,olleH"`.

## Contract (TDD)

This PR opens as a **draft** carrying a RED test suite that defines the contract
up front; the implementation lands on top of it:

- `contract_test.go` — `Reverse("abc") == "cba"` and the Unicode case above.
- `e2e_test.go` — `TestE2E` round-trips the feature: `Reverse(Reverse(s)) == s`
  across ASCII, Unicode, empty, single-rune, and mixed inputs.

`contract_test.go` and `e2e_test.go` are protected (listed in
`.vajra-protected`) and pin the public API; only `strutil.go` is added to turn
the suite GREEN.

## Testing

- Before implementation: `go test ./...` fails RED (package does not compile —
  `Reverse` undefined).
- After implementation: `go test ./...` and `go test -cover ./...` pass GREEN.
