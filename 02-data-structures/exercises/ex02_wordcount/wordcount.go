// Package ex02_wordcount — Exercise 2: maps + strings/runes.
package ex02_wordcount

import (
	"strings"
	"unicode/utf8"
)

// WordCount returns a map from each whitespace-separated word to how many times
// it appears. Words are compared case-INSENSITIVELY ("Go" and "go" are the same
// word, counted under the lowercase key "go").
//
//	WordCount("Go go GO rust") -> {"go": 3, "rust": 1}
//	WordCount("")              -> {} (empty, non-nil map)
//
// Hints:
//   - strings.Fields splits on any run of whitespace and drops empties.
//   - strings.ToLower normalizes case.
//   - Reading a missing key returns 0, so m[w]++ works without a pre-check.
//
// TODO: implement. Returns nil for now (which will fail the non-nil check).
func WordCount(text string) map[string]int {
	counts := make(map[string]int)
	splitted_string := strings.Fields(text)

	for _, v := range splitted_string {
		v = strings.ToLower(v)
		counts[v]++
	}

	return counts
}

// FirstRune returns the first UNICODE CHARACTER (rune) of s as a string, and
// ok=false for the empty string.
//
//	FirstRune("héllo") -> "h", true
//	FirstRune("étage") -> "é", true   (must NOT return a broken half-byte)
//	FirstRune("")      -> "",  false
//
// Hint: convert to []rune, or use utf8.DecodeRuneInString. Do NOT just take s[0]
// — that returns a single byte and corrupts multi-byte characters.
//
// TODO: implement.
func FirstRune(s string) (string, bool) {
	if s == "" {
		return "", false
	}

	rune, size := utf8.DecodeRuneInString(s)
	return string(rune), size > 0
}
