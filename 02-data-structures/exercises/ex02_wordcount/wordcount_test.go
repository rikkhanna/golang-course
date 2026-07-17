package ex02_wordcount

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want map[string]int
	}{
		{"empty", "", map[string]int{}},
		{"single", "go", map[string]int{"go": 1}},
		{"case insensitive", "Go go GO rust", map[string]int{"go": 3, "rust": 1}},
		{"extra spaces", "  a   b  a ", map[string]int{"a": 2, "b": 1}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := WordCount(tc.in)
			if got == nil {
				t.Fatalf("WordCount(%q) = nil, want a non-nil map", tc.in)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("WordCount(%q) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}

func TestFirstRune(t *testing.T) {
	cases := []struct {
		name   string
		in     string
		want   string
		wantOK bool
	}{
		{"ascii", "hello", "h", true},
		{"accented", "étage", "é", true},
		{"empty", "", "", false},
		{"emoji", "🚀go", "🚀", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := FirstRune(tc.in)
			if got != tc.want || ok != tc.wantOK {
				t.Errorf("FirstRune(%q) = (%q, %t), want (%q, %t)",
					tc.in, got, ok, tc.want, tc.wantOK)
			}
		})
	}
}
