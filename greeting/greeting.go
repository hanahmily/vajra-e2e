// Package greeting produces friendly, deterministic greeting strings.
//
// This file defines the APPROVED PUBLIC INTERFACE for the feature. The bodies
// are intentionally left as unimplemented stubs (they return zero values) so
// that the accompanying unit and end-to-end tests start RED. Production
// behavior is implemented in a later (GREEN) phase, not here.
package greeting

// DefaultName is the name used when a caller supplies an empty or
// whitespace-only name.
const DefaultName = "World"

// Greet returns a greeting for the given name.
//
// Contract:
//   - Surrounding whitespace in name is trimmed before use.
//   - A non-empty (post-trim) name n yields exactly "Hello, <n>!".
//   - An empty or whitespace-only name yields "Hello, World!" (see DefaultName).
//
// Examples:
//
//	Greet("Alice")   == "Hello, Alice!"
//	Greet("  Bob  ") == "Hello, Bob!"
//	Greet("")        == "Hello, World!"
func Greet(name string) string {
	// TODO(contract): implement in the GREEN phase. Stub returns zero value so
	// unit/e2e tests are RED.
	_ = name
	return ""
}

// GreetAll returns one greeting per input name, preserving input order and
// applying the same per-name rules as Greet.
//
// Contract:
//   - The result has exactly len(names) elements, in the same order.
//   - result[i] == Greet(names[i]) for every i.
//   - A nil or empty input yields a non-nil, empty slice (len 0).
func GreetAll(names []string) []string {
	// TODO(contract): implement in the GREEN phase. Stub returns zero value so
	// unit/e2e tests are RED.
	_ = names
	return nil
}
