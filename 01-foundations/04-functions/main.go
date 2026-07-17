package main

import "fmt"

// Basic function: params share a type can be grouped (a, b int).
func add(a, b int) int {
	return a + b
}

// Multiple return values — the Go idiom for "result + error" (more in section 03).
// Here we return the quotient and a boolean "ok".
func safeDiv(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

// Named return values: the names are pre-declared as zero-valued vars. A bare
// `return` returns their current values. Use sparingly — clear for short funcs,
// confusing for long ones.
func minMax(nums []int) (min, max int) {
	if len(nums) == 0 {
		return // returns 0, 0
	}
	min, max = nums[0], nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return // returns the named values
}

// Variadic: `nums ...int` collects any number of trailing int args into a slice.
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Functions are first-class values. This one returns a closure that captures
// and mutates `count` across calls.
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	fmt.Println("add:", add(2, 3))

	if q, ok := safeDiv(10, 2); ok {
		fmt.Println("safeDiv:", q)
	}
	if _, ok := safeDiv(10, 0); !ok {
		fmt.Println("safeDiv: division by zero avoided")
	}

	lo, hi := minMax([]int{4, 9, 1, 7})
	fmt.Printf("min=%d max=%d\n", lo, hi)

	fmt.Println("sum:", sum(1, 2, 3, 4))
	// Spread an existing slice into a variadic call with `...`
	xs := []int{5, 6, 7}
	fmt.Println("sum(slice...):", sum(xs...))

	next := counter()
	fmt.Println("counter:", next(), next(), next(), next()) // 1 2 3

	// --- defer ---
	// Deferred calls run in LIFO order when the function returns. Classic use:
	// pairing acquire/release (open/close, lock/unlock) right next to each other.
	fmt.Print("defer order (expect 2 1 0): ")
	deferOrder()
	fmt.Println()

	defer fmt.Println("deferred: runs LAST, when main returns")
	fmt.Println("end of main body")
}

// deferOrder shows two things about defer:
//  1. Deferred calls fire in LIFO (last-in, first-out) order.
//  2. A deferred call's ARGUMENTS are evaluated when `defer` runs, not when the
//     call finally executes — so each iteration captures the current i, and
//     they print in reverse: 2 1 0.
func deferOrder() {
	for i := 0; i < 3; i++ {
		defer fmt.Printf("%d ", i)
	}
}
