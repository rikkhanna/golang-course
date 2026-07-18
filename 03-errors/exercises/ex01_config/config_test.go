package ex01_config

import (
	"errors"
	"strconv"
	"testing"
)

func TestGet(t *testing.T) {
	c := Config{"host": "localhost", "port": "8080"}

	if got, err := c.Get("host"); err != nil || got != "localhost" {
		t.Errorf(`Get("host") = (%q, %v), want ("localhost", nil)`, got, err)
	}

	_, err := c.Get("missing")
	if err == nil {
		t.Fatal(`Get("missing") returned nil error, want a wrapped ErrMissingKey`)
	}
	if !errors.Is(err, ErrMissingKey) {
		t.Errorf("Get error = %v; errors.Is(err, ErrMissingKey) = false, want true", err)
	}
	// Context requirement: the message should mention the key.
	if !contains(err.Error(), "missing") {
		t.Errorf("Get error = %q; want it to name the missing key", err.Error())
	}
}

func TestGetInt(t *testing.T) {
	c := Config{"port": "8080", "bad": "not-a-number"}

	if got, err := c.GetInt("port"); err != nil || got != 8080 {
		t.Errorf(`GetInt("port") = (%d, %v), want (8080, nil)`, got, err)
	}

	// Missing key must still be detectable as ErrMissingKey.
	if _, err := c.GetInt("nope"); !errors.Is(err, ErrMissingKey) {
		t.Errorf("GetInt missing: errors.Is(err, ErrMissingKey) = false, want true (err=%v)", err)
	}

	// Bad value: the underlying strconv error must be reachable in the chain.
	_, err := c.GetInt("bad")
	if err == nil {
		t.Fatal(`GetInt("bad") returned nil error, want a wrapped parse error`)
	}
	var numErr *strconv.NumError
	if !errors.As(err, &numErr) {
		t.Errorf("GetInt bad value: errors.As(&strconv.NumError) = false, want true (err=%v)", err)
	}
}

func contains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
