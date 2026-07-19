# Draft PR

## Title
feat(greeting): add greeting package and `greet` CLI (contract + RED tests)

## Body
Adds a new `greeting` Go package and a `greet` command-line tool that consumes
it. This PR lands the **contract phase** of the milestone-TDD workflow: the
approved public interface plus failing (RED) unit and end-to-end tests. No
production behavior is implemented yet — that follows in the GREEN phase.

### What's included
- **`greeting/greeting.go`** — approved public interface:
  - `const DefaultName = "World"`
  - `func Greet(name string) string`
  - `func GreetAll(names []string) []string`
  - Bodies are stubs (return zero values) so tests are RED.
- **`cmd/greet/main.go`** — CLI interface: prints one greeting line per
  argument to stdout, or a single default greeting when no args are given;
  exits 0 on success.
- **`greeting/greeting_test.go`** — RED unit tests covering the per-name rules.
- **`e2e/greet_e2e_test.go`** — RED end-to-end test that builds the binary and
  runs it like a real user, asserting stdout and exit code.

### Contract / requirements
- `Greet` trims surrounding whitespace; a non-empty name `n` → `"Hello, <n>!"`.
- An empty or whitespace-only name → `"Hello, World!"` (`DefaultName`).
- `GreetAll` returns one greeting per input, in order; nil/empty input → a
  non-nil empty slice.
- `greet Alice Bob` → `Hello, Alice!\nHello, Bob!\n`.
- `greet` (no args) → `Hello, World!\n`.

### Test status (expected at this phase)
- `go build ./...` — passes (production code compiles).
- `go vet ./...` — clean.
- `go test ./greeting/ ./e2e/` — **FAIL (RED)**, as designed. The GREEN phase
  will implement the stubbed bodies to turn these green without altering the
  protected interface or test assertions.

### Protected files
See `.vajra-protected`.
