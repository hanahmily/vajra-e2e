# Requirements — `greeting` package and `greet` CLI

## Goal
Provide a small, deterministic greeting library (the `greeting` package) and a
command-line tool (`greet`) that consumes it. This document captures the
approved contract for the milestone-TDD **contract phase**: the requirements,
the public interface, and the RED unit and end-to-end tests that lock the
contract. "Contract phase" means the public signatures and the tests are
final, while the function bodies remain unimplemented stubs; no production
behavior is implemented in this phase.

## Functional requirements

### `greeting` package (`greeting/greeting.go`)
- `const DefaultName = "World"` — the name used when a caller supplies an empty
  or whitespace-only name.
- `func Greet(name string) string`
  - Leading and trailing whitespace in `name` is trimmed before use; interior
    whitespace is preserved.
  - A non-empty (post-trim) name `n` yields exactly `"Hello, <n>!"` — the
    literal text `Hello, `, then `n`, then a single `!`, with no trailing
    newline.
  - An empty or whitespace-only name (i.e. one that is empty after trimming)
    yields `"Hello, World!"` (`DefaultName`).
- `func GreetAll(names []string) []string`
  - Returns exactly `len(names)` elements, in the same order.
  - `result[i] == Greet(names[i])` for every `i`.
  - A nil or empty input yields a non-nil, empty slice (length 0).

### `greet` CLI (`cmd/greet/main.go`)
- `greet [name...]`.
- For each argument, one greeting line is written to stdout, in order.
- With no arguments, a single default greeting (`Hello, World!`) is written.
- On success the process exits 0.

## Contract scope
The contract is intentionally minimal and complete: the rules above are the
entire contract. No additional edge cases (locale, punctuation, encoding, etc.)
are in scope. There are no contract gaps to close in review.

## Test contract (RED in this phase)
- `greeting/greeting_test.go` — unit tests for the per-name rules, ordering, and
  the non-nil empty-slice guarantee.
- `e2e/greet_e2e_test.go` — `TestE2E` builds the `greet` binary from
  `cmd/greet` (which delegates to the `greeting` package) and runs it like a
  real user over a slice of names, asserting stdout and exit code.

### Note on test references (review clarification)
The end-to-end test references **only** the `greet` CLI and, transitively, the
`greeting` package (`greeting.Greet` / `greeting.GreetAll`). There is no `calc`
package and no `calc.Add` symbol in this feature; the earlier review note asking
to confirm a `calc.Add` reference does not apply to the greeting contract and no
such reference was added. The four contract files are protected (see
`.vajra-protected`) and their assertions are unchanged.

## Expected verification status at contract phase
- `go build ./...` — passes (production code compiles).
- `go vet ./...` — clean.
- `go test ./...` — **FAIL (RED)** by design (stubbed bodies).
- `go test -run TestE2E ./...` — **FAIL (RED)** by design.

The GREEN phase implements the stubbed bodies to turn these green without
altering the protected interface or test assertions.

## Protected files
See `.vajra-protected`.
