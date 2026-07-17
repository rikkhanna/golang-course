// Package ex01_classify — Exercise 1: control flow + functions.
package ex01_classify

import (
	"strconv"
)

// Classify returns:
//   - "fizzbuzz" if n is divisible by both 3 and 5
//   - "fizz"     if divisible by 3 only
//   - "buzz"     if divisible by 5 only
//   - the number as a decimal string otherwise (e.g. 7 -> "7")
//
// Hints:
//   - Use the modulo operator: n%3 == 0.
//   - Convert an int to its decimal string with strconv.Itoa (import "strconv").
//   - An expressionless `switch { case ...: }` reads nicely here.
//
// TODO: implement this. Right now it returns "" so the tests fail.
func Classify(n int) string {

	switch {
	case n%3 == 0 && n%5 == 0:
		return "fizzbuzz"
	case n%3 == 0:
		return "fizz"
	case n%5 == 0:
		return "buzz"
	}
	return strconv.Itoa(n)
}
