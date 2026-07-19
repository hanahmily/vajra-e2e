# Requirements — `mathx.Clamp`

## Goal
Provide a small, dependency-free math helper package `mathx` at the root of the
Go module `mathxlive`, exposing a single public function that bounds an integer
to an inclusive range.

## Approved public interface
The following interface is the complete, approved contract for this milestone.
No additional exported symbols are in scope.

```go
package mathx

// Clamp returns v constrained to the inclusive range [lo, hi]:
//   - if v < lo, it returns lo
//   - if v > hi, it returns hi
//   - otherwise it returns v
func Clamp(v, lo, hi int) int
```

- **Import path:** `mathxlive` (the module-root package, named `mathx`).
- **Callers** use it as `mathx.Clamp(...)`.

## Functional requirements
1. `Clamp(v, lo, hi)` returns `v` when `lo <= v <= hi`.
2. `Clamp(v, lo, hi)` returns `lo` when `v < lo`.
3. `Clamp(v, lo, hi)` returns `hi` when `v > hi`.
4. The function is pure: no allocation, no I/O, no global state.

## Acceptance criteria (contract)
Asserted by `contract_test.go` (unit) and `e2e_test.go` (end-to-end):

| Input                | Expected |
|----------------------|----------|
| `Clamp(5, 0, 10)`    | `5`      |
| `Clamp(-1, 0, 10)`   | `0`      |
| `Clamp(99, 0, 10)`   | `10`     |

The e2e test exercises a real use case: clamping every element of a slice of
values into a fixed `[lo, hi]` window and comparing against the expected slice.

## Scope / non-goals
- The contract is **minimal and intentionally complete**. There are **no extra
  edge cases in scope** (e.g. inverted ranges where `lo > hi`, other integer
  widths, generics). Reviewers should therefore report `contract_gap_count=0`
  and `issue_count=0`.
- No production implementation (`mathx.go`) is created at the assign step; the
  test suite must fail **RED** until the coder adds it.

## Protected files
`contract_test.go` and `e2e_test.go` are protected (see `.vajra-protected`).
The coder implements `mathx.go` to turn the suite GREEN **without** modifying
the protected files.
