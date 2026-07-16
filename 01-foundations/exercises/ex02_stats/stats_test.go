package ex02_stats

import "testing"

func TestMinMax(t *testing.T) {
	cases := []struct {
		name             string
		in               []int
		wantMin, wantMax int
		wantOK           bool
	}{
		{"empty", nil, 0, 0, false},
		{"single", []int{5}, 5, 5, true},
		{"ascending", []int{1, 2, 3}, 1, 3, true},
		{"mixed", []int{4, -2, 9, 0}, -2, 9, true},
		{"negatives", []int{-5, -1, -9}, -9, -1, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			min, max, ok := MinMax(tc.in...)
			if min != tc.wantMin || max != tc.wantMax || ok != tc.wantOK {
				t.Errorf("MinMax(%v) = (%d, %d, %t), want (%d, %d, %t)",
					tc.in, min, max, ok, tc.wantMin, tc.wantMax, tc.wantOK)
			}
		})
	}
}

func TestRunningAverage(t *testing.T) {
	avg := RunningAverage()
	steps := []struct {
		in   float64
		want float64
	}{
		{10, 10},
		{20, 15},
		{30, 20},
		{0, 15},
	}
	for i, s := range steps {
		if got := avg(s.in); got != s.want {
			t.Errorf("step %d: avg(%g) = %g, want %g", i, s.in, got, s.want)
		}
	}

	// A second, independent closure must NOT share state with the first.
	avg2 := RunningAverage()
	if got := avg2(100); got != 100 {
		t.Errorf("fresh closure: avg2(100) = %g, want 100", got)
	}
}
