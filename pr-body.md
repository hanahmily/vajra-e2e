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
- **`e2e/greet_e2e_test.go`** — RED end-to-end `TestE2E` that builds the binary
  and runs it like a real user over a slice of names, asserting stdout and exit
  code.

### Contract / requirements
- `Greet` trims surrounding whitespace; a non-empty name `n` → `"Hello, <n>!"`.
- An empty or whitespace-only name → `"Hello, World!"` (`DefaultName`).
- `GreetAll` returns one greeting per input, in order; nil/empty input → a
  non-nil empty slice.
- `greet Alice Bob` → `Hello, Alice!\nHello, Bob!\n`.
- `greet` (no args) → `Hello, World!\n`.

See `requirements.md` for the full contract.

### Test status (expected at this phase)
- `go build ./...` — passes (production code compiles).
- `go vet ./...` — clean.
- `go test ./...` — **FAIL (RED)**, as designed.
- `go test -run TestE2E ./...` — **FAIL (RED)**, as designed. The GREEN phase
  will implement the stubbed bodies to turn these green without altering the
  protected interface or test assertions.

### Protected files
See `.vajra-protected`.

### Revision notes (addressing review feedback)
This resubmission responds to the previous review, which asked to (1) clarify
the wording of `requirements.md` and (2) confirm the e2e test's references.

- **Requirements wording clarified.** `requirements.md` now spells out what the
  "contract phase" means, that only leading/trailing whitespace is trimmed
  (interior whitespace is preserved), the exact shape of the `"Hello, <n>!"`
  output, and precisely what the e2e test builds and drives.
- **e2e test references confirmed.** `e2e/greet_e2e_test.go` builds the `greet`
  binary from `cmd/greet` and drives it as a user would; it references **only**
  the `greet` CLI and, transitively, the `greeting` package
  (`greeting.Greet` / `greeting.GreetAll`). There is **no `calc` package and no
  `calc.Add`** anywhere in this feature, so the requested `calc.Add` reference
  does not apply and was not introduced — doing so would break the greeting
  contract, and the contract files are protected in any case. The protected
  interface and every test assertion are unchanged.
