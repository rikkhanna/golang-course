package main

import "fmt"

// Package-level constants and vars can use `var`/`const` but NOT `:=`.
// Constants are compile-time values: numbers, strings, booleans.
const pi = 3.14159       // untyped constant — adapts to context
const greeting = "hello" // typed as string when used

// A typed constant block. iota auto-increments within a const block (0,1,2,...),
// a common way to build enum-like sequences.
type Weekday int

const (
	Sunday  Weekday = iota // 0
	Monday                 // 1
	Tuesday                // 2
)

func main() {
	// --- Declaration styles ---
	var explicit int = 10 // full form: keyword, name, type, value
	var inferred = 20     // type inferred from the value (int)
	short := 30           // short form: only inside functions; type inferred
	fmt.Println(explicit, inferred, short)

	// --- Zero values: variables declared without a value are NOT garbage ---
	var (
		i int     // 0
		f float64 // 0
		b bool    // false
		s string  // "" (empty, not nil)
		p *int    // nil
	)
	fmt.Printf("zero values: i=%d f=%g b=%t s=%q p=%v\n", i, f, b, s, p)

	// --- Basic types you'll actually use ---
	// int (machine-word size), int64/int32/..., uint*, float64 (default float),
	// bool, string, byte (=uint8), rune (=int32, a Unicode code point).
	var count int = 100
	var ratio float64 = 3.0 / 4.0
	var initial rune = 'A' // 65
	fmt.Printf("count=%d ratio=%.2f initial=%c(%d)\n", count, ratio, initial, initial)

	// --- Conversions are ALWAYS explicit. Go never auto-converts numeric types. ---
	var whole int = 7
	var asFloat float64 = float64(whole) / 2 // without float64(), this won't compile
	fmt.Printf("asFloat=%.1f\n", asFloat)

	// --- Constants in action ---
	fmt.Printf("pi=%v greeting=%q Tuesday=%d\n", pi, greeting, Tuesday)
}
