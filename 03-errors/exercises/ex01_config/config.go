// Package ex01_config — Exercise 1: sentinel errors + wrapping.
package ex01_config

import (
	"errors"
	"fmt"
	"strconv"
)

// ErrMissingKey is a sentinel returned when a key isn't present.
var ErrMissingKey = errors.New("missing key")

// Config is a simple string->string settings store.
type Config map[string]string

// Get returns the value for key. On a missing key it wraps ErrMissingKey with
// context so callers get BOTH a readable message and errors.Is detectability.
func (c Config) Get(key string) (string, error) {
	// Map indexing (comma-ok): v is the VALUE, ok is a BOOL (present?).
	v, ok := c[key]
	if !ok {
		// %w embeds the sentinel in the chain -> errors.Is(err, ErrMissingKey) is true.
		return "", fmt.Errorf("get %q: %w", key, ErrMissingKey)
	}
	return v, nil
}

// GetInt returns the value for key parsed as an int.
func (c Config) GetInt(key string) (int, error) {
	// Reuse Get. If the key is missing, its error already wraps ErrMissingKey,
	// so we just propagate it unchanged (no need to re-wrap).
	s, err := c.Get(key)
	if err != nil {
		return 0, err
	}
	// Present but maybe not a number: wrap the strconv error so the caller can
	// still errors.As it to a *strconv.NumError.
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("get int %q: %w", key, err)
	}
	return n, nil
}
