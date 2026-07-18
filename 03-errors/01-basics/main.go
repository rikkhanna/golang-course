package main

import (
	"errors"
	"fmt"
)

// The idiom: fallible functions return (result, error). error is the LAST return.
// On success return a valid value + nil error; on failure return the zero value
// + a non-nil error.

func divide(a, b int) (int, error) {
	if b == 0 {
		// errors.New creates a simple error from a string.
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Sentinel error: a package-level error value callers can compare against.
// Convention: name starts with "Err".
var ErrEmpty = errors.New("input is empty")

// parseFirst returns the first byte of s, or ErrEmpty. fmt.Errorf builds an
// error with formatting (here without wrapping — just a message).
func parseFirst(s string) (byte, error) {
	if s == "" {
		return 0, ErrEmpty
	}
	if len(s) > 32 {
		return 0, fmt.Errorf("input too long: %d bytes", len(s))
	}
	return s[0], nil
}

func main() {
	// --- The standard check pattern ---
	q, err := divide(10, 2)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("10/2 =", q)
	}

	// Handle the error immediately; don't let a bad value flow onward.
	if _, err := divide(1, 0); err != nil {
		fmt.Println("error:", err) // division by zero
	}

	// --- Comparing against a sentinel with == (works when NOT wrapped) ---
	if _, err := parseFirst(""); err == ErrEmpty {
		fmt.Println("got ErrEmpty (compared with ==)")
	}

	// A formatted, non-sentinel error.
	if _, err := parseFirst("this string is definitely longer than thirty-two bytes"); err != nil {
		fmt.Println("error:", err)
	}

	b, err := parseFirst("go")
	if err != nil {
		fmt.Println("unexpected:", err)
		return
	}
	fmt.Printf("first byte: %c\n", b)
}
