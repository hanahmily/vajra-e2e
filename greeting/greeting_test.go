package greeting_test

import (
	"reflect"
	"testing"

	"github.com/hanahmily/vajra-e2e/greeting"
)

func TestGreet(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{name: "normal name", input: "Alice", want: "Hello, Alice!"},
		{name: "trims surrounding whitespace", input: "  Bob  ", want: "Hello, Bob!"},
		{name: "empty name uses default", input: "", want: "Hello, World!"},
		{name: "whitespace-only name uses default", input: "   ", want: "Hello, World!"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := greeting.Greet(tc.input); got != tc.want {
				t.Fatalf("Greet(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func TestGreetAll(t *testing.T) {
	got := greeting.GreetAll([]string{"Alice", "", "  Bob "})
	want := []string{"Hello, Alice!", "Hello, World!", "Hello, Bob!"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GreetAll(...) = %#v, want %#v", got, want)
	}
}

func TestGreetAllEmptyIsNonNil(t *testing.T) {
	got := greeting.GreetAll(nil)
	if got == nil {
		t.Fatal("GreetAll(nil) returned nil, want non-nil empty slice")
	}
	if len(got) != 0 {
		t.Fatalf("GreetAll(nil) len = %d, want 0", len(got))
	}
}

func TestDefaultNameConstant(t *testing.T) {
	if greeting.DefaultName != "World" {
		t.Fatalf("DefaultName = %q, want %q", greeting.DefaultName, "World")
	}
}
