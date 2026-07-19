# Requirements — `calc` package (module `calclive`)

## Goal
Add a Go package named `calc` at the module root of a module named `calclive`,
exposing a single addition function.

## Approved public interface
```go
// Package calc lives at the root of module "calclive".
package calc

// Add returns the sum of a and b.
func Add(a, b int) int
```

- Import path: `calclive` (module root).
- Package name: `calc`.
- The interface is defined and frozen by the protected RED tests; no separate
  interface stub file is introduced at assign.

## Functional requirements
1. **R1 — Addition.** `calc.Add(a, b)` returns `a + b` for all `int` inputs.
   - Pinned by the contract test: `calc.Add(2, 3) == 5`.
2. **R2 — Composability over a sequence.** Folding `calc.Add` over a slice of
   integers accumulates their total, so it is usable as a real summation
   primitive.
   - Pinned by the e2e test: folding `calc.Add` over `[1 2 3 4 5]` yields `15`.

## Scope boundaries
- The contract is **minimal and intentionally complete**. There are **no extra
  edge cases in scope** (no overflow handling, no variadic form, no other
  arithmetic). A reviewer evaluating this contract must find
  `contract_gap_count == 0` and `issue_count == 0`.
- Production behavior is **not** implemented at assign: `calc.go` does not exist
  yet, so `go test ./...` fails RED. The coder makes the suite GREEN by adding
  `calc.go` only, without touching the protected files.

## Protected files (`.vajra-protected`)
- `contract_test.go` — RED unit test for R1.
- `e2e_test.go` — RED end-to-end test for R2.

## Definition of done
- `calc.go` implements `func Add(a, b int) int` returning `a + b`.
- `go test ./...` passes GREEN.
- No protected file is modified.
