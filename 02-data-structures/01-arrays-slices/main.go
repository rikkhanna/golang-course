package main

import "fmt"

func main() {
	// --- Arrays: fixed size, size is PART of the type ---
	// [3]int and [4]int are different, incompatible types. Arrays are VALUES:
	// assigning or passing one copies all elements. You'll rarely use them directly.
	var a [3]int // [0 0 0]
	a[0] = 1
	b := [...]int{5, 6, 7} // [...] lets the compiler count -> [3]int
	fmt.Println("arrays:", a, b, len(b))

	// --- Slices: the workhorse. A view (ptr, len, cap) over a backing array. ---
	s := []int{10, 20, 30} // slice literal (note: no size in the brackets)
	fmt.Printf("slice: %v len=%d cap=%d\n", s, len(s), cap(s))

	// make([]T, len, cap): allocate with a known size / preallocated capacity.
	buf := make([]int, 2, 8)
	fmt.Printf("make: %v len=%d cap=%d\n", buf, len(buf), cap(buf))

	// --- append: grows the slice; may reallocate when capacity is exceeded ---
	nums := make([]int, 0, 2)
	for i := 1; i <= 5; i++ {
		nums = append(nums, i*10)
		fmt.Printf("append %d -> %v len=%d cap=%d\n", i*10, nums, len(nums), cap(nums))
	}
	// Notice cap jumps (2 -> 4 -> 8...): when full, Go allocates a bigger array
	// and copies. ALWAYS assign append's result back: nums = append(nums, ...).

	// --- Slicing: s[low:high] shares the SAME backing array (half-open range) ---
	base := []int{0, 1, 2, 3, 4}
	mid := base[1:4] // [1 2 3], shares memory with base
	fmt.Printf("mid=%v len=%d cap=%d\n", mid, len(mid), cap(mid))

	// --- ALIASING GOTCHA: mutating mid changes base ---
	mid[0] = 999
	fmt.Println("after mid[0]=999 -> base:", base) // [0 999 2 3 4]

	// --- copy: make an INDEPENDENT copy to break aliasing ---
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	n := copy(dst, src) // copies min(len(dst), len(src)) elements
	dst[0] = -1
	fmt.Printf("copied %d, dst=%v src=%v (src unchanged)\n", n, dst, src)

	// --- Delete an element (idiomatic, order-preserving) ---
	xs := []int{10, 20, 30, 40}
	i := 1 // remove index 1 (value 20)
	xs = append(xs[:i], xs[i+1:]...)
	fmt.Println("after delete idx1:", xs) // [10 30 40]
}
