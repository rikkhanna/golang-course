package main

import "fmt"

type point struct {
	X, Y int
}

// Go passes everything by VALUE. This receives a COPY; mutating it is invisible
// to the caller.
func tryMutateValue(p point) {
	p.X = 100
}

// Take a POINTER to mutate the caller's value. &p at the call site, *p to deref.
func mutateViaPointer(p *point) {
	p.X = 100 // shorthand for (*p).X = 100; Go auto-dereferences struct fields
}

func main() {
	// & takes the address; * dereferences. There is NO pointer arithmetic in Go.
	x := 42
	p := &x                 // p is *int, points at x
	fmt.Println("*p =", *p) // 42
	*p = 7                  // write through the pointer
	fmt.Println("x =", x)   // 7 — x changed via p

	// --- Value vs pointer semantics ---
	pt := point{1, 2}
	tryMutateValue(pt)
	fmt.Println("after tryMutateValue:", pt) // {1 2} — unchanged (got a copy)
	mutateViaPointer(&pt)
	fmt.Println("after mutateViaPointer:", pt) // {100 2} — changed

	// --- new(T) allocates a zeroed T and returns a *T ---
	np := new(point) // *point pointing at {0 0}
	np.Y = 5
	fmt.Println("new(point):", *np) // {0 5}

	// --- Nil pointers: the zero value of any pointer type ---
	var np2 *point
	fmt.Println("nil pointer:", np2 == nil) // true
	// fmt.Println(np2.X) // <- panics: nil pointer dereference. Always check/init.

	// When to use pointers (rules of thumb):
	//   - You need to MUTATE the caller's value.
	//   - The value is large and copying is wasteful.
	//   - The type's identity matters (shared, not copied).
	// Otherwise prefer values — they're simpler and avoid nil surprises.
}
