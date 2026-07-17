package ex01_slices

import (
	"testing"
)

// equalInts treats nil and empty slices as equal (both length 0). This matches
// idiomatic Go, where a function may return nil OR []int{} for "no elements"
// and callers don't care. reflect.DeepEqual would wrongly distinguish them.
func equalInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestDedup(t *testing.T) {
	cases := []struct {
		name string
		in   []int
		want []int
	}{
		{"empty", []int{}, []int{}},
		{"no dups", []int{1, 2, 3}, []int{1, 2, 3}},
		{"consecutive", []int{1, 2, 2, 3}, []int{1, 2, 3}},
		{"scattered", []int{1, 2, 2, 3, 1, 3}, []int{1, 2, 3}},
		{"all same", []int{7, 7, 7}, []int{7}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			in := make([]int, len(tc.in))
			copy(in, tc.in) // preserves length; non-nil snapshot for mutation check
			got := Dedup(in)
			if !equalInts(got, tc.want) {
				t.Errorf("Dedup(%v) = %v, want %v (order must be first-seen)", tc.in, got, tc.want)
			}
			if !equalInts(in, tc.in) {
				t.Errorf("Dedup mutated its input: got %v, was %v", in, tc.in)
			}
		})
	}
}

func TestRotateLeft(t *testing.T) {
	cases := []struct {
		name string
		in   []int
		k    int
		want []int
	}{
		{"by 2", []int{1, 2, 3, 4, 5}, 2, []int{3, 4, 5, 1, 2}},
		{"by 0", []int{1, 2, 3}, 0, []int{1, 2, 3}},
		{"full wrap", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"over wrap", []int{1, 2, 3}, 4, []int{2, 3, 1}},
		{"empty", []int{}, 3, []int{}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			in := make([]int, len(tc.in))
			copy(in, tc.in)
			got := RotateLeft(in, tc.k)
			if !equalInts(got, tc.want) {
				t.Errorf("RotateLeft(%v, %d) = %v, want %v", tc.in, tc.k, got, tc.want)
			}
			if !equalInts(in, tc.in) {
				t.Errorf("RotateLeft mutated its input: got %v, was %v", in, tc.in)
			}
		})
	}
}
