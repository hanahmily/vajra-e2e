# Requirements — package `sm` (module `smlive`)

## Goal
Provide a tiny, reusable integer-maximum helper as a Go package named `sm`
living at the module root of module `smlive`.

## Approved public interface
The milestone approves exactly one exported symbol. No other public API is in scope.

```go
package sm

// Max returns the larger of a and b.
// When a == b, either value (they are equal) is returned.
func Max(a, b int) int
```

- Package name: `sm`
- Module path: `smlive` (package sits at the module root)
- Import path for consumers: `smlive`

## Functional requirements
1. `Max(a, b int) int` MUST return `a` when `a >= b`, and `b` otherwise.
2. `Max(2, 5)` MUST return `5`.
3. `Max(9, 1)` MUST return `9`.
4. Given a non-empty slice of ints, folding `Max` over the slice MUST yield the
   maximum element (real end-to-end use case exercised by `TestE2E`).
5. `Max` MUST be pure: no side effects, no allocation, deterministic.

## Non-functional requirements
- No third-party dependencies; standard library only.
- `Max` operates on the built-in `int` type (platform word size).

## Out of scope
- Variadic max, generics, or max over floats/other types.
- Handling of empty slices (the caller seeds the fold with the first element).

## Contract / test surface (RED at assign)
The following files are protected and define the contract. They are RED until
the coder adds `sm.go` implementing `Max`:
- `contract_test.go` — asserts `Max(2,5)==5` and `Max(9,1)==9`.
- `e2e_test.go` — `TestE2E` reduces a slice to its maximum via `Max`.

At the assign step there is intentionally **no** `sm.go`, so `go test ./...`
fails to build (RED). The coder later adds `sm.go` to make it GREEN without
touching any protected file.

## Definition of done (for the coder, downstream)
- `sm.go` defines `func Max(a, b int) int` in `package sm`.
- `go test ./...` passes (GREEN).
- No protected file is modified.
