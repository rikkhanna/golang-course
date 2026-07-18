package main

import (
	"errors"
	"fmt"
	"os"
)

// Sentinel errors for a tiny "store".
var (
	ErrNotFound = errors.New("not found")
	ErrDenied   = errors.New("permission denied")
)

// lookup simulates a low-level operation returning a sentinel.
func lookup(key string) (string, error) {
	switch key {
	case "alice":
		return "admin", nil
	case "root":
		return "", ErrDenied
	default:
		return "", ErrNotFound
	}
}

// getRole wraps lower-level errors with CONTEXT using %w. The wrap keeps the
// original sentinel reachable so callers can still errors.Is it.
func getRole(key string) (string, error) {
	role, err := lookup(key)
	if err != nil {
		return "", fmt.Errorf("getRole %q: %w", key, err)
	}
	return role, nil
}

func main() {
	// --- Wrapping preserves the chain; %v/%s printing shows the full context ---
	_, err := getRole("ghost")
	fmt.Println("err:", err) // getRole "ghost": not found

	// --- errors.Is: does the chain contain this sentinel? (value match) ---
	if errors.Is(err, ErrNotFound) {
		fmt.Println("errors.Is(err, ErrNotFound) -> true (matched through the wrap)")
	}
	if !errors.Is(err, ErrDenied) {
		fmt.Println("errors.Is(err, ErrDenied) -> false")
	}

	// --- errors.As: find a specific TYPE in the chain and extract it ---
	// os.Open returns a *os.PathError wrapping the underlying cause.
	_, ferr := os.Open("/no/such/file/here")
	fmt.Println("open err:", ferr)

	var pathErr *os.PathError
	if errors.As(ferr, &pathErr) {
		// We recovered the concrete type and can read its fields.
		fmt.Printf("errors.As -> *os.PathError: op=%q path=%q\n", pathErr.Op, pathErr.Path)
	}
	// And errors.Is still matches the standard sentinel through os's wrapping.
	if errors.Is(ferr, os.ErrNotExist) {
		fmt.Println("errors.Is(ferr, os.ErrNotExist) -> true")
	}

	// --- %w vs %v: %v flattens to a string and BREAKS the chain ---
	broken := fmt.Errorf("getRole: %v", ErrNotFound) // note %v, not %w
	// This prints false: the sentinel is no longer reachable in the chain.
	fmt.Println("errors.Is on the flattened error:", errors.Is(broken, ErrNotFound))
}
