package sm

import "testing"

// TestMax pins the approved contract for Max. It is RED until sm.go defines
// func Max(a, b int) int (the package fails to build without it).
func TestMax(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{"second larger", 2, 5, 5},
		{"first larger", 9, 1, 9},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := Max(c.a, c.b); got != c.want {
				t.Fatalf("Max(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
			}
		})
	}
}
