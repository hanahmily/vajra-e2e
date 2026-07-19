# Requirements — `strutil.Reverse`

## Summary

Add a Go package `strutil` at the module root (module `strutillive`) exposing a
single function that reverses a string by Unicode code point (rune), not by
byte. This document is the authoritative, **intentionally minimal and complete**
contract: everything required is listed here, and nothing beyond it is in scope.

## Public interface (approved, FINAL)

```go
package strutil

// Reverse returns s with its runes in reverse order. Reversal is rune-correct:
// multi-byte UTF-8 runes are reversed as whole runes and are never split into
// their constituent bytes.
func Reverse(s string) string
```

- Module path: `strutillive`.
- Package name: `strutil`.
- Package location: module root (the `.go` files live in the repository root).

## Functional requirements

- **R1 — ASCII reversal.** `Reverse("abc") == "cba"`.
- **R2 — Rune-correct Unicode reversal.** Reversal operates on runes, not
  bytes. `Reverse("Hello, 世界") == "界世 ,olleH"`. Multi-byte runes are never
  corrupted.
- **R3 — Involution / round-trip.** Reversing twice yields the original input:
  for every string `s`, `Reverse(Reverse(s)) == s`. This is the real-world
  use case exercised end-to-end.
- **R4 — Boundary inputs.** The empty string and single-rune strings reverse to
  themselves: `Reverse("") == ""`, `Reverse("a") == "a"`. (Covered by R3's
  round-trip over `""` and `"a"`.)

## Explicitly out of scope

The contract is deliberately kept to the requirements above. The following are
**NOT** in scope and must not be added — introducing them would be gold-plating,
not a contract gap:

- Normalization of combining marks / grapheme-cluster (as opposed to rune)
  reversal.
- Byte-level or word-level reversal variants.
- Any additional exported symbols, options, or error returns (`Reverse` never
  errors and takes no options).

## Test contract (RED at assign)

| File | Kind | Asserts |
| --- | --- | --- |
| `contract_test.go` | unit | R1 (`"abc"→"cba"`) and R2 (`"Hello, 世界"→"界世 ,olleH"`). |
| `e2e_test.go` | e2e (`TestE2E`) | R3/R4 round-trip over ASCII, Unicode, empty, single-rune, palindrome, and mixed inputs. |

`strutil.go` is intentionally **absent** at assign, so the package does not
compile and both `go test ./...` and `go test -run TestE2E ./...` exit non-zero
(RED). The coder turns the suite GREEN by adding `strutil.go` alone, without
touching the protected files listed in `.vajra-protected`.

## Reviewer expectation

Because this contract is minimal and complete, a correct review reports
`contract_gap_count == 0` and `issue_count == 0` once the implementation
satisfies R1–R4.
