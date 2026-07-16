// Package ex02_stats — Exercise 2: variadic functions + closures.
package ex02_stats

// MinMax returns the smallest and largest of the given ints, and ok=false if
// no arguments were passed (in which case min and max should be 0).
//
// This is variadic: call it as MinMax(3, 1, 2) or MinMax(slice...).
//
// TODO: implement. Currently returns zero values so tests fail.
func MinMax(nums ...int) (min, max int, ok bool) {
	return 0, 0, false
}

// RunningAverage returns a CLOSURE. Each call to the returned function accepts
// one new number and returns the average of every number seen so far.
//
// Example:
//   avg := RunningAverage()
//   avg(10) // 10
//   avg(20) // 15
//   avg(30) // 20
//
// Hint: capture a running sum and a count in variables declared in
// RunningAverage; the returned func closes over them.
//
// TODO: implement. Currently the returned func always yields 0.
func RunningAverage() func(x float64) float64 {
	return func(x float64) float64 {
		return 0
	}
}
