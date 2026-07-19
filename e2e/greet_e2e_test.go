// Package e2e drives the compiled `greet` binary the way a real user would:
// it builds the command, runs it with arguments, and asserts on its stdout and
// exit code. It fails (RED) until the greeting package is implemented.
package e2e

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// buildGreet compiles cmd/greet into a temp binary and returns its path.
func buildGreet(t *testing.T) string {
	t.Helper()
	bin := filepath.Join(t.TempDir(), "greet")
	cmd := exec.Command("go", "build", "-o", bin, "github.com/hanahmily/vajra-e2e/cmd/greet")
	cmd.Dir = repoRoot(t)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("build greet: %v\n%s", err, out)
	}
	return bin
}

// repoRoot returns the module root (parent of this e2e directory).
func repoRoot(t *testing.T) string {
	t.Helper()
	wd, err := filepath.Abs("..")
	if err != nil {
		t.Fatalf("resolve repo root: %v", err)
	}
	return wd
}

func run(t *testing.T, bin string, args ...string) string {
	t.Helper()
	var stdout bytes.Buffer
	cmd := exec.Command(bin, args...)
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		t.Fatalf("run %q %v: %v", bin, args, err)
	}
	return stdout.String()
}

// TestE2E drives the compiled binary through its real use cases: greeting a
// slice of names in one invocation, defaulting when given none, and trimming
// surrounding whitespace. It stays RED until the greeting package is
// implemented in the GREEN phase.
func TestE2E(t *testing.T) {
	bin := buildGreet(t)

	t.Run("greets a slice of names in order", func(t *testing.T) {
		names := []string{"Alice", "Bob", "Carol"}
		got := run(t, bin, names...)
		want := "Hello, Alice!\nHello, Bob!\nHello, Carol!\n"
		if got != want {
			t.Fatalf("greet %v stdout = %q, want %q", names, got, want)
		}
	})

	t.Run("no args uses default", func(t *testing.T) {
		got := run(t, bin)
		want := "Hello, World!\n"
		if got != want {
			t.Fatalf("greet (no args) stdout = %q, want %q", got, want)
		}
	})

	t.Run("trims surrounding whitespace", func(t *testing.T) {
		got := strings.TrimRight(run(t, bin, "  Carol  "), "\n")
		if got != "Hello, Carol!" {
			t.Fatalf("greet \"  Carol  \" stdout = %q, want %q", got, "Hello, Carol!")
		}
	})
}
