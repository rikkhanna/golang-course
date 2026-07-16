package main

import "fmt"

func main() {
	// --- if ---
	// No parentheses around the condition. Braces are mandatory.
	x := 7
	if x%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	// if can carry an initializer statement, scoped to the if/else. This is
	// extremely common with functions that return a value + something to check.
	if half := x / 2; half > 2 {
		fmt.Println("half is big:", half)
	}

	// --- for: Go's ONLY loop keyword ---
	// 1) classic three-part
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println("sum 1..5 =", sum)

	// 2) condition-only (this is Go's "while")
	n := 1
	for n < 100 {
		n *= 2
	}
	fmt.Println("first power of 2 >= 100:", n)

	// 3) infinite loop with break
	tries := 0
	for {
		tries++
		if tries == 3 {
			break
		}
	}
	fmt.Println("tries:", tries)

	// 4) range over a slice: index + value. Use _ to discard either.
	nums := []int{10, 20, 30}
	for idx, val := range nums {
		fmt.Printf("nums[%d]=%d\n", idx, val)
	}

	// --- switch ---
	// No fallthrough by default (opposite of C). Each case breaks automatically.
	switch day := 3; day {
	case 0, 6: // multiple values per case
		fmt.Println("weekend")
	case 1, 2, 3, 4, 5:
		fmt.Println("weekday")
	default:
		fmt.Println("invalid")
	}

	// Expressionless switch = a clean if/else-if ladder.
	score := 82
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	default:
		fmt.Println("C or below")
	}
}
