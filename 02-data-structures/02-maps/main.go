package main

import (
	"fmt"
	"sort"
)

func main() {
	// --- Creating maps ---
	// make(map[K]V) gives an empty, writable map. A map literal seeds values.
	ages := make(map[string]int)
	ages["alice"] = 30
	ages["bob"] = 25

	scores := map[string]int{"x": 1, "y": 2} // literal
	fmt.Println("ages:", ages, "scores:", scores)

	// --- Reading: a missing key returns the VALUE TYPE's zero value ---
	fmt.Println("missing key 'carol':", ages["carol"]) // 0, no error, no panic

	// --- comma-ok: distinguish "absent" from "present but zero" ---
	if age, ok := ages["alice"]; ok {
		fmt.Println("alice is", age)
	}
	if _, ok := ages["carol"]; !ok {
		fmt.Println("carol not in map")
	}

	// --- delete: safe even if the key is absent ---
	delete(ages, "bob")
	delete(ages, "nobody") // no-op, no panic
	fmt.Println("after delete:", ages)

	// --- nil map: reading is fine, WRITING PANICS ---
	var nm map[string]int                        // nil map (zero value)
	fmt.Println("read nil map:", nm["anything"]) // 0, fine
	// nm["x"] = 1 // <- would panic: "assignment to entry in nil map"
	// Fix: nm = make(map[string]int) before writing.

	// --- Iteration order is RANDOMIZED. Never depend on it. ---
	// To iterate deterministically, collect keys and sort them.
	keys := make([]string, 0, len(scores))
	for k := range scores { // ranging a map: k is key; k,v for key+value
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s=%d ", k, scores[k])
	}
	fmt.Println()

	// --- A common pattern: map as a set ---
	seen := map[string]struct{}{} // struct{} uses zero memory for the value
	for _, w := range []string{"a", "b", "a", "c", "b"} {
		seen[w] = struct{}{}
	}
	fmt.Println("unique count:", len(seen))
}
