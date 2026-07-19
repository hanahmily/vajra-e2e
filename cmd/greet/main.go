// Command greet prints greetings to stdout for the names given as arguments.
//
// This file defines the APPROVED CLI INTERFACE (flag/argument surface and I/O
// contract). Its greeting logic is delegated to the greeting package, whose
// stubbed implementation keeps the end-to-end test RED until the GREEN phase.
//
// Usage:
//
//	greet [name...]
//
// Contract:
//   - For each argument, one greeting line is written to stdout, in order.
//   - With no arguments, a single default greeting ("Hello, World!") is written.
//   - On success the process exits 0.
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/hanahmily/vajra-e2e/greeting"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout))
}

// run is the testable entry point: it writes greetings for args to out and
// returns the process exit code.
func run(args []string, out io.Writer) int {
	names := args
	if len(names) == 0 {
		names = []string{""} // triggers the default greeting
	}
	for _, line := range greeting.GreetAll(names) {
		fmt.Fprintln(out, line)
	}
	return 0
}
