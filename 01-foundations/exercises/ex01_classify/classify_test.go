package ex01_classify

import "testing"

// This is a "table-driven" test: a slice of cases, one loop. It's the dominant
// testing style in Go. t.Run gives each case its own named subtest.
func TestClassify(t *testing.T) {
	cases := []struct {
		name string
		in   int
		want string
	}{
		{"three", 3, "fizz"},
		{"five", 5, "buzz"},
		{"fifteen", 15, "fizzbuzz"},
		{"nine", 9, "fizz"},
		{"ten", 10, "buzz"},
		{"seven", 7, "7"},
		{"one", 1, "1"},
		{"zero", 0, "fizzbuzz"}, // 0 is divisible by everything
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Classify(tc.in)
			if got != tc.want {
				t.Errorf("Classify(%d) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
