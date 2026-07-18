// Package ex02_validation — Exercise 2: a custom error type + errors.As.
package ex02_validation

import "fmt"

// FieldError is a custom error type that reports WHICH field failed and why.
// It satisfies the built-in error interface via the Error() method below.
type FieldError struct {
	Field string
	Msg   string
}

// Error makes *FieldError an error. Pointer receiver -> the type that satisfies
// `error` is *FieldError, which is why Validate returns &FieldError{...} and the
// test matches with `var fe *FieldError; errors.As(err, &fe)`.
func (e *FieldError) Error() string {
	return fmt.Sprintf("field %q: %s", e.Field, e.Msg)
}

// User is the thing we validate.
type User struct {
	Name string
	Age  int
}

// Validate returns nil if u is valid, otherwise a *FieldError for the FIRST
// problem found. Guard clauses check in order; nil is returned only at the end.
func Validate(u User) error {
	if u.Name == "" {
		return &FieldError{Field: "name", Msg: "must not be empty"}
	}
	if u.Age < 0 {
		return &FieldError{Field: "age", Msg: "must not be negative"}
	}
	if u.Age > 150 {
		return &FieldError{Field: "age", Msg: "is implausibly large"}
	}
	return nil
}
