package simple

import (
	"fmt"
	"testing"
)

func TestAbs(t *testing.T) {
	cases := []struct {
		name string
		x    int
		want int
	}{
		{"positive", 5, 5},
		{"zero", 0, 0},
		{"negative", -5, 5},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Abs(c.x)
			if got != c.want {
				t.Errorf("want: %d, got: %d", c.want, got)
			}
		})
	}
}

func TestMin(t *testing.T) {
	cases := []struct {
		values []int
		want   int
	}{
		{[]int{10}, 10},
		{[]int{10, 20}, 10},
		{[]int{20, 10}, 10},
		{[]int{10, 20, 30}, 10},
		{[]int{30, 10, 20}, 10},
		{[]int{20, 30, 10}, 10},
	}
	for _, c := range cases {
		t.Run(fmt.Sprint(c.values), func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("unwanted panic: %#v", err)
				}
			}()

			got := Min(c.values...)
			if got != c.want {
				t.Errorf("want: %d, got: %d", c.want, got)
			}
		})
	}

	panicCases := []struct {
		values []int
		want   string
	}{
		{[]int{}, "no values"},
	}
	for _, c := range panicCases {
		t.Run(fmt.Sprint(c.values), func(t *testing.T) {
			defer func() {
				err := recover()
				if err != c.want {
					t.Errorf("wanted panic: %#v, got panic: %#v", c.want, err)
				}
			}()

			Min(c.values...)
		})
	}
}

func TestMax(t *testing.T) {
	cases := []struct {
		values []int
		want   int
	}{
		{[]int{10}, 10},
		{[]int{20, 10}, 20},
		{[]int{10, 20}, 20},
		{[]int{30, 20, 10}, 30},
		{[]int{10, 30, 20}, 30},
		{[]int{20, 10, 30}, 30},
	}
	for _, c := range cases {
		t.Run(fmt.Sprint(c.values), func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("unwanted panic: %#v", err)
				}
			}()

			got := Max(c.values...)
			if got != c.want {
				t.Errorf("want: %d, got: %d", c.want, got)
			}
		})
	}

	panicCases := []struct {
		values []int
		want   string
	}{
		{[]int{}, "no values"},
	}
	for _, c := range panicCases {
		t.Run(fmt.Sprint(c.values), func(t *testing.T) {
			defer func() {
				err := recover()
				if err != c.want {
					t.Errorf("wanted panic: %#v, got panic: %#v", c.want, err)
				}
			}()

			Max(c.values...)
		})
	}
}

func TestPow(t *testing.T) {
	cases := []struct {
		base     int
		exponent int
		want     int
	}{
		{5, 0, 1}, {5, 1, 5}, {5, 2, 25}, {5, 3, 125}, {5, 4, 625},
		{5, 5, 3125}, {5, 6, 15625}, {5, 7, 78125}, {5, 8, 390625},
		{-5, 0, 1}, {-5, 1, -5}, {-5, 2, 25}, {-5, 3, -125}, {-5, 4, 625},
		{-5, 5, -3125}, {-5, 6, 15625}, {-5, 7, -78125}, {-5, 8, 390625},
	}
	for _, c := range cases {
		caseName := fmt.Sprintf("(%d)^(%d)", c.base, c.exponent)
		t.Run(caseName, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("unwanted panic: %#v", err)
				}
			}()

			got := Pow(c.base, c.exponent)
			if got != c.want {
				t.Errorf("want: %d, got: %d", c.want, got)
			}
		})
	}

	panicCases := []struct {
		base     int
		exponent int
		want     string
	}{
		{5, -1, "invalid exponent: -1"},
	}
	for _, c := range panicCases {
		caseName := fmt.Sprintf("(%d)^(%d)", c.base, c.exponent)
		t.Run(caseName, func(t *testing.T) {
			defer func() {
				err := recover()
				if err != c.want {
					t.Errorf("wanted panic: %#v, got panic: %#v", c.want, err)
				}
			}()

			Pow(c.base, c.exponent)
		})
	}
}

func TestCeil(t *testing.T) {
	cases := []struct {
		divident int
		dividor  int
		want     int
	}{
		{7, 3, 3},
		{6, 3, 2}, {5, 3, 2}, {4, 3, 2},
		{3, 3, 1}, {2, 3, 1}, {1, 3, 1},
		{0, 3, 0}, {-1, 3, 0}, {-2, 3, 0},
		{-3, 3, -1}, {-4, 3, -1}, {-5, 3, -1},
		{-6, 3, -2}, {-7, 3, -2},
		{7, -3, -2}, {6, -3, -2},
		{5, -3, -1}, {4, -3, -1}, {3, -3, -1},
		{2, -3, 0}, {1, -3, 0}, {0, -3, 0},
		{-1, -3, 1}, {-2, -3, 1}, {-3, -3, 1},
		{-4, -3, 2}, {-5, -3, 2}, {-6, -3, 2},
		{-7, -3, 3},
		{2, 1, 2}, {1, 1, 1}, {0, 1, 0}, {-1, 1, -1}, {-2, 1, -2},
		{2, -1, -2}, {1, -1, -1}, {0, -1, -0}, {-1, -1, 1}, {-2, -1, 2},
	}
	for _, c := range cases {
		caseName := fmt.Sprintf("Ceil[%d/%d]", c.divident, c.dividor)
		t.Run(caseName, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("unwanted panic: %#v", err)
				}
			}()

			got := Ceil(c.divident, c.dividor)
			if got != c.want {
				t.Errorf("want: %d, got: %d", c.want, got)
			}
		})
	}

	panicCases := []struct {
		divident int
		dividor  int
		want     string
	}{
		{1, 0, "divide by zero"},
	}
	for _, c := range panicCases {
		caseName := fmt.Sprintf("Ceil[%d/%d]", c.divident, c.dividor)
		t.Run(caseName, func(t *testing.T) {
			defer func() {
				err := recover()
				if err != c.want {
					t.Errorf("wanted panic: %#v, got panic: %#v", c.want, err)
				}
			}()

			Ceil(c.divident, c.dividor)
		})
	}
}

func TestEucDiv(t *testing.T) {
	cases := []struct {
		divident  int
		dividor   int
		quotient  int
		remainder int
	}{
		{0, 3, 0, 0}, {0, -3, 0, 0},
		{7, 1, 7, 0}, {7, -1, -7, 0}, {-7, 1, -7, 0}, {-7, -1, 7, 0},
		{7, 3, 2, 1}, {7, -3, -2, 1}, {-7, 3, -3, 2}, {-7, -3, 3, 2},
		{7, 10, 0, 7}, {7, -10, 0, 7}, {-7, 10, -1, 3}, {-7, -10, 1, 3},
	}
	for _, c := range cases {
		caseName := fmt.Sprintf("(%d)_divided_by_(%d)", c.divident, c.dividor)
		t.Run(caseName, func(t *testing.T) {
			// assert the expectation
			{
				divident := c.quotient*c.dividor + c.remainder
				if divident != c.divident {
					t.Fatalf("wrong expectation: (%d)*(%d)+(%d)=(%d) (not %d)",
						c.quotient, c.dividor, c.remainder, divident, c.divident)
				}
			}
			{
				maxRem := Abs(c.dividor)
				if c.remainder < 0 || c.remainder >= maxRem {
					t.Fatalf("wrong expectation: remainder %d should be between 0<= and <%d",
						c.remainder, maxRem)
				}
			}

			defer func() {
				if err := recover(); err != nil {
					t.Errorf("unwanted panic: %#v", err)
				}
			}()

			quotient, remainder := EucDiv(c.divident, c.dividor)
			if quotient != c.quotient {
				t.Errorf("quotient want: %d, got: %d", c.quotient, quotient)
			}
			if remainder != c.remainder {
				t.Errorf("remainder want: %d, got: %d", c.remainder, remainder)
			}
		})
	}

	panicCases := []struct {
		divident int
		dividor  int
		want     string
	}{
		{1, 0, "divide by zero"},
	}
	for _, c := range panicCases {
		caseName := fmt.Sprintf("(%d)_divided_by_(%d)", c.divident, c.dividor)
		t.Run(caseName, func(t *testing.T) {
			defer func() {
				err := recover()
				if err != c.want {
					t.Errorf("wanted panic: %#v, got panic: %#v", c.want, err)
				}
			}()

			EucDiv(c.divident, c.dividor)
		})
	}
}

func TestEucRem(t *testing.T) {
	// check TestEucDiv
}

func TestGCD(t *testing.T) {
	cases := []struct {
		a    int
		b    int
		want int
	}{
		{700, 42, 14}, {42, 700, 14}, {42, 42, 42}, {42, 1, 1}, {1, 1, 1},
	}
	for _, c := range cases {
		caseName := fmt.Sprintf("GCD(%d,%d)", c.a, c.b)
		t.Run(caseName, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("unwanted panic: %#v", err)
				}
			}()

			got := GCD(c.a, c.b)
			if got != c.want {
				t.Errorf("want: %d, got: %d", c.want, got)
			}
		})
	}

	panicCases := []struct {
		a    int
		b    int
		want string
	}{
		{1, 0, "invalid value: 0"},
		{1, -1, "invalid value: -1"},
		{0, 1, "invalid value: 0"},
		{-1, 1, "invalid value: -1"},
	}
	for _, c := range panicCases {
		caseName := fmt.Sprintf("GCD(%d,%d)", c.a, c.b)
		t.Run(caseName, func(t *testing.T) {
			defer func() {
				err := recover()
				if err != c.want {
					t.Errorf("wanted panic: %#v, got panic: %#v", c.want, err)
				}
			}()

			GCD(c.a, c.b)
		})
	}
}
