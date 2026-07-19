# Add `calc` package with `Add(a, b int) int`

## Title

Add `calc` package with `Add(a, b int) int` at module root

## Body

### What

Introduces a Go package `calc` at the root of module `calclive` exposing a
single pure function:

```go
func Add(a, b int) int // returns a + b
```

### Why

Establishes the module's first public interface and a TDD contract for it,
following the milestone-tdd workflow: the tests are authored and land RED
before any production code is written, then the implementation makes them
GREEN without touching the protected test files.

### Contract (authored at assign, RED)

- `contract_test.go` — asserts `calc.Add(2, 3) == 5`.
- `e2e_test.go` — `TestE2E` folds `calc.Add` over `[]int{1, 2, 3, 4, 5}` and
  asserts the running total is `15` (a realistic "sum a slice" use case).
- Both are external `calc_test` tests importing `calclive`.
- `.vajra-protected` lists these files so the implementation cannot alter them.

### Implementation (coder)

- Adds `calc.go` (`package calc`) with `Add`.
- Leaves protected files untouched.

### Verification

- At assign: `go test ./...` fails RED (no `calc.go`, package does not compile).
- After implementation: `go test ./...` passes GREEN.
