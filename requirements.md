# Requirements — `calc` package

## Summary

Add a Go package named `calc` at the root of a module named `calclive`. The
package exposes a single pure function that adds two integers.

## Functional requirements

1. The module path is `calclive` (declared in `go.mod`, Go 1.25).
2. A package named `calc` lives at the module root.
3. The package exports `func Add(a, b int) int` that returns `a + b`.
4. `Add` is pure: no side effects, no state, deterministic for all inputs.
5. After the coder adds the implementation, `go test ./...` MUST pass GREEN.

## Approved public interface (final contract)

```go
package calc

// Add returns the sum of a and b.
func Add(a, b int) int
```

- Package name: `calc`
- Import path: `calclive`
- Callers reference the function as `calc.Add(...)`.

This interface is **final**. The coder implements it in a new `calc.go` at the
module root and must **not** modify the protected test files.

## Test contract (RED at assign)

- `contract_test.go` — unit test asserting `calc.Add(2, 3) == 5`.
- `e2e_test.go` — `TestE2E` folding `calc.Add` over the slice
  `[]int{1, 2, 3, 4, 5}` from an initial `0`, asserting the total is `15`.

Both test files live in the external test package `calc_test` and import
`calclive`. At assign time `calc.go` does not exist, so the suite fails to
compile and `go test ./...` is RED. This is intentional.

## Definition of done (for the coder)

- Add `calc.go` at the module root with `package calc` and the `Add` function.
- Do not edit `contract_test.go`, `e2e_test.go`, or `.vajra-protected`.
- `go test ./...` passes (both `TestAdd` and `TestE2E` GREEN).
