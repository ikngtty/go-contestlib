package modular

import (
	"fmt"
	"testing"
)

func TestExtendedEuclidean(t *testing.T) {
	cases := []struct {
		a   int
		b   int
		gcd int
	}{
		{0, 0, 0},
		{0, 3, 3}, {3, 0, 3},
		{1, 1, 1},
		{3, 1, 1}, {1, 3, 1},
		{20, 20, 20},
		{15, 28, 1}, {28, 15, 1},
		{28, 100, 4}, {100, 28, 4},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%dx+%dy", c.a, c.b), func(t *testing.T) {
			gcd, x, y := ExtendedEuclidean(c.a, c.b)
			if gcd != c.gcd {
				t.Errorf("GCD want: %d, got: %d", c.gcd, gcd)
			}
			d := c.a*x + c.b*y
			if d != gcd {
				t.Errorf("%d*%d+%d*%d=%d (not %d)", c.a, x, c.b, y, d, gcd)
			}
		})
	}
}
