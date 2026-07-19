# Requirements — `calc` package

## Summary

Add a Go package named `calc` at the root of a module named `calclive`. The
package exposes a single pure function that adds two integers. This is the
module's first public interface, authored under the milestone-TDD workflow:
the tests land RED before any production code exists, and the coder later makes
them GREEN without touching the protected test files.

## Functional requirements

1. The module path is `calclive`, declared in `go.mod` targeting Go 1.25.
2. A package named `calc` lives at the module root.
3. The package exports `func Add(a, b int) int` returning the sum `a + b`.
4. `Add` is pure: no side effects, no shared state, deterministic for every
   pair of inputs.
5. Integer overflow follows Go's standard two's-complement `int` semantics; no
   special handling is required.
6. After the coder adds the implementation, `go test ./...` MUST pass GREEN.

## Approved public interface (final contract)

```go
package calc

// Add returns the sum of a and b.
func Add(a, b int) int
```

- Package name: `calc`
- Import path: `calclive` (the module root)
- Callers reference the function as `calc.Add(...)`.

This interface is **final**. The coder implements it in a new `calc.go` at the
module root and must **not** modify the protected test files or `.vajra-protected`.

## Test contract (RED at assign)

- `contract_test.go` — unit test asserting `calc.Add(2, 3) == 5`, plus a
  table of cases covering positive, negative, and zero operands and
  commutativity (`Add(a, b) == Add(b, a)`).
- `e2e_test.go` — `TestE2E` folds `calc.Add` over the slice
  `[]int{1, 2, 3, 4, 5}` from an initial `0`, asserting the running total is
  `15`. This simulates the canonical real use case: summing a slice on top of
  the public interface.

Both test files live in the external test package `calc_test` and import
`calclive`. At assign time `calc.go` does not exist, so the `calclive` package
has no buildable Go files and the suite fails to compile — `go test ./...` is
RED. This is intentional.

## Definition of done (for the coder)

- Add `calc.go` at the module root declaring `package calc` and implementing
  `Add(a, b int) int`.
- Do **not** edit `contract_test.go`, `e2e_test.go`, or `.vajra-protected`.
- `go test ./...` passes with both `TestAdd` and `TestE2E` GREEN.
