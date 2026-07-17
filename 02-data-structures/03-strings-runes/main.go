package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// A string is an immutable sequence of BYTES, conventionally UTF-8 encoded.
	s := "héllo" // the 'é' is 2 bytes in UTF-8

	// --- len is BYTE count, not character count ---
	fmt.Printf("len(bytes)=%d runeCount=%d\n", len(s), utf8.RuneCountInString(s))
	// -> len=6, runeCount=5

	// --- Indexing gives a BYTE (uint8), not a character ---
	fmt.Printf("s[0]=%d (%c)\n", s[0], s[0]) // 104 (h)

	// --- Ranging over a string yields (byteIndex, rune) ---
	// Note the index jumps by 2 at the multi-byte 'é'.
	for i, r := range s {
		fmt.Printf("  index %d: rune %c (U+%04X)\n", i, r, r)
	}

	// --- Conversions ---
	bs := []byte(s) // copy into a mutable byte slice
	rs := []rune(s) // decode into runes; len(rs) == rune count
	fmt.Printf("[]byte len=%d, []rune len=%d\n", len(bs), len(rs))
	fmt.Println("back to string:", string(rs))

	// Strings are immutable: s[0] = 'H' won't compile. Build new strings instead.
	// strings.Builder is the efficient way to concatenate in a loop.
	var b strings.Builder
	for i := 0; i < 3; i++ {
		b.WriteString("ab")
	}
	fmt.Println("builder:", b.String()) // ababab

	// --- Handy strings functions ---
	fmt.Println(strings.ToUpper("go"), strings.Contains("gopher", "ph"),
		strings.Split("a,b,c", ","), strings.TrimSpace("  hi  "))

	// --- Reverse a string CORRECTLY (by runes, not bytes) ---
	fmt.Println("reversed:", reverseRunes("héllo"))
}

// reverseRunes reverses by runes so multi-byte characters aren't corrupted.
// Reversing bytes would split the 'é' into invalid UTF-8.
func reverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
