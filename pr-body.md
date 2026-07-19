# Add `calc` package with `Add` to module `calclive`

## Summary
Introduces a new Go package `calc` at the root of module `calclive`, exposing
`func Add(a, b int) int` that returns `a + b`.

This PR is opened as a **draft** at the contract (assign) stage of a
milestone-TDD workflow. It currently contains only the approved public
interface and its **RED** tests; the implementation (`calc.go`) is added in a
follow-up commit by the coder so the suite turns GREEN.

## Public interface
```go
package calc

// Add returns the sum of a and b.
func Add(a, b int) int
```

## Requirements pinned by tests
- **R1 — Addition:** `calc.Add(2, 3) == 5` (`contract_test.go`).
- **R2 — Composability:** folding `calc.Add` over `[1 2 3 4 5]` yields `15`
  (`e2e_test.go`), demonstrating a real summation use case.

## Contract scope
The contract is minimal and intentionally complete — no additional edge cases
(overflow, variadic, etc.) are in scope. Expected review outcome:
`contract_gap_count == 0`, `issue_count == 0`.

## Test plan
- At assign: `go test ./...` fails RED (no `calc.go` yet — module `calclive`
  has no buildable package for the tests to import).
- After implementation: `go test ./...` passes GREEN with `calc.go` added and
  no protected file modified.

## Protected files
- `contract_test.go`
- `e2e_test.go`
