// Package ex01_slices — Exercise 1: slices without aliasing bugs.
package ex01_slices

// Dedup returns a new slice containing the elements of in with consecutive AND
// non-consecutive duplicates removed, preserving first-seen order.
//
//	Dedup([1,2,2,3,1,3]) -> [1,2,3]
//	Dedup([]int{})       -> []int{} (empty, non-nil is fine too)
//
// Requirements:
//   - Do NOT mutate the input slice `in` (the test checks it's untouched).
//   - Preserve the order in which each value first appears.
//
// Hint: a map[int]struct{} as a "seen" set + append to a fresh result slice.
//
// TODO: implement. Returns nil for now.
func Dedup(in []int) []int {
	if in == nil {
		return nil
	}
	seen := map[int]struct{}{}
	result := []int{}
	for _, v := range in {
	  if _, ok := seen[v]; !ok {
		seen[v] = struct{}{}
		result = append(result, v)
	  }
	}
	return result
}

// RotateLeft returns a NEW slice rotated left by k positions.
//
//	RotateLeft([1,2,3,4,5], 2) -> [3,4,5,1,2]
//
// Requirements:
//   - Do NOT mutate `in`.
//   - k may be larger than len(in) or 0; handle with modulo. Empty input -> empty.
//
// Hint: normalize k with k % len(in), then append(in[k:], in[:k]...) into a
// fresh slice (careful not to alias `in`).
//
// TODO: implement. Returns nil for now.
func RotateLeft(in []int, k int) []int {
	if in == nil {
		return nil
	}

	l := len(in)
	if l == 0 {
		return []int{}
	}
	k = k % l
	if k < 0 {
		k += l
	}
	result := make([]int, l)
	copy(result, in[k:])
	copy(result[l-k:], in[:k])
	return result
}
