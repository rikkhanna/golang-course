What is Rune ?

- A rune is a Unicode code point
A rune is just an alias for int32. It holds the numeric identity of a single Unicode character (a "code point"). 'A' is the rune 65, 'é' is 233, '🚀' is 128640.

var r rune = 'A'   // rune literals use SINGLE quotes
fmt.Println(r)     // 65   (it's a number!)
fmt.Printf("%c\n", r) // A  (%c formats a rune as its character)
Why Go needs a separate type: bytes ≠ characters
A Go string is a sequence of bytes, encoded as UTF-8. In UTF-8, a character can be 1 to 4 bytes:

Character	Bytes it takes	Rune value
h	1	104
é	2	233
🚀	4	128640
So the string "héllo" is 5 characters but 6 bytes. This is why in the example you studied:

s := "héllo"
len(s)                      // 6  ← BYTE count
utf8.RuneCountInString(s)   // 5  ← character count
s[0]                        // 104 ← a single BYTE (uint8), not a rune
Indexing a string gives you a byte, not a character. If you did s[1] on "héllo" you'd get half of the é — a broken, meaningless byte.

The two ways to get runes out of a string
1. range decodes runes for you — and the index jumps by how many bytes each took:

for i, r := range "héllo" {
    fmt.Printf("index %d = %c\n", i, r)
}
// index 0 = h
// index 1 = é
// index 3 = l   ← notice: jumped from 1 to 3, because é took 2 bytes
2. Convert to []rune when you need to index characters directly (e.g. reversing, or grabbing the "first character"):

rs := []rune("héllo")  // [104 233 108 108 111], length 5
string(rs[0])          // "h" — a whole character, safe
The one rule to remember
Byte = raw storage unit. Rune = one actual character.
Index/len a string → bytes. range or []rune a string → runes.

This is exactly what Exercise 2's FirstRune is testing: s[0] would return a byte and mangle the é/🚀 cases, so you must go through []rune (or utf8.DecodeRuneInString) to return a whole character. Want to keep going on the exercises, or any other part of runes/UTF-8 you want unpacked first?