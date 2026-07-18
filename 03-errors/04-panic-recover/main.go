package main

import "fmt"

// panic/recover is NOT Go's error handling — errors-as-values is. panic is for:
//   - unrecoverable programmer bugs (index out of range, nil deref),
//   - impossible states you'd rather crash than continue past,
//   - occasionally, unwinding across many layers in a self-contained subsystem
//     (e.g. a parser) that recovers at its boundary and converts to an error.

func main() {
	// --- recover must run inside a DEFERRED function to stop a panic ---
	if r, ok := safeDivide(10, 2); ok {
		fmt.Println("result:", r)
	}
	if _, ok := safeDivide(1, 0); !ok { // panics inside; recover -> ok=false
		fmt.Println("1/0 was recovered, ok=false")
	}
	fmt.Println("main still running after recovered panic")

	// --- A panic that is NOT recovered crashes the program (commented out) ---
	// var p *int
	// _ = *p // panic: nil pointer dereference -> program exits non-zero
}

// safeDivide converts a panic into a normal (value, ok) result. This pattern —
// recover at a boundary and translate to a return value — is the legitimate use
// of recover. Note: for THIS example plain checking (if b==0) is better; we
// force a panic only to demonstrate recover.
func safeDivide(a, b int) (result int, ok bool) {
	defer func() {
		// recover() returns nil if there's no panic in flight; otherwise it
		// returns the panic value and stops the unwinding.
		if r := recover(); r != nil {
			fmt.Printf("  recovered from panic: %v\n", r)
			result, ok = 0, false // named returns let us set the result here
		}
	}()

	// This panics when b == 0 ("integer divide by zero").
	return a / b, true
}
