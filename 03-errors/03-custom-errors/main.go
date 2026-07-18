package main

import (
	"errors"
	"fmt"
)

// A custom error type is any type with an Error() string method. Use one when
// callers need STRUCTURED data about the failure (a field name, a code, a
// retry-after), not just a message.
type ValidationError struct {
	Field string
	Msg   string
}

// Implement the error interface. Pointer or value receiver both work; pointer
// is common so errors.As can match *ValidationError.
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation: field %q: %s", e.Field, e.Msg)
}

// An error type can also implement Unwrap() to participate in a chain, so
// errors.Is/As can see through it to an inner cause.
type QueryError struct {
	Query string
	Err   error // the wrapped cause
}

func (e *QueryError) Error() string { return fmt.Sprintf("query %q: %v", e.Query, e.Err) }
func (e *QueryError) Unwrap() error { return e.Err }

var ErrTimeout = errors.New("timeout")

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Msg: "must be non-negative"}
	}
	if age > 150 {
		return &ValidationError{Field: "age", Msg: "implausibly large"}
	}
	return nil
}

func runQuery(q string) error {
	// Pretend this timed out; wrap the sentinel inside our custom type.
	return &QueryError{Query: q, Err: ErrTimeout}
}

func main() {
	// --- errors.As extracts the concrete type so we can read its fields ---
	err := validateAge(-3)
	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Printf("field=%q msg=%q\n", ve.Field, ve.Msg)
	}

	// --- A custom type with Unwrap() lets errors.Is see the inner sentinel ---
	qerr := runQuery("SELECT 1")
	fmt.Println("qerr:", qerr)
	if errors.Is(qerr, ErrTimeout) {
		fmt.Println("errors.Is(qerr, ErrTimeout) -> true (via Unwrap)")
	}
	var qe *QueryError
	if errors.As(qerr, &qe) {
		fmt.Printf("extracted QueryError for query %q\n", qe.Query)
	}
}
