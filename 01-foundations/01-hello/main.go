// Package main defines an executable program. The compiler produces a runnable
// binary only for package main, and only when it also has a func main().
package main

// imports pull in other packages. "fmt" is the standard formatting/IO package.
// An unused import is a COMPILE ERROR in Go — try adding "os" here without using
// it and `go run` will refuse to build.
import (
	"fmt"
	// "os"
)

// main is the entry point. No arguments, no return value.
func main() {
	// Println writes its args separated by spaces, then a newline.
	fmt.Println("hello, go")

	// Printf uses C-style format verbs. %s = string, %d = int, %v = any value's
	// default format, %q = quoted string, \n = newline (Printf does NOT add one).
	name := "gopher"
	fmt.Printf("hello, %s! the answer %v is %d\n", name, 01, 42)
}
