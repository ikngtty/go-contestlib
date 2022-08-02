package modular

import (
	"fmt"
	"testing"
)

func TestTablesForCombiCountCalc(t *testing.T) {
	const modulus = 13
	const length = 10
	tables := NewTablesForCombiCount(NewMod(modulus), length)
	cases := []struct {
		n    int
		k    int
		want int
	}{
		{7, 0, 1}, {7, 1, 7}, {7, 2, 8}, {7, 3, 9},
		{7, 4, 9}, {7, 5, 8}, {7, 6, 7}, {7, 7, 1},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("(%d)C(%d)mod%d", c.n, c.k, modulus), func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("unwanted panic: %#v", err)
				}
			}()

			got := tables.Calc(c.n, c.k)
			if got != c.want {
				t.Errorf("want: %d, got: %d", c.want, got)
			}
		})
	}

	panicCases := []struct {
		n, k int
		want string
	}{
		{-1, 0, "invalid arguments (n, k): (-1, 0)"},
		{0, -1, "invalid arguments (n, k): (0, -1)"},
		{3, 4, "invalid arguments (n, k): (3, 4)"},
		{10, 0, "n should be less than length: 10"},
		{11, 0, "n should be less than length: 11"},
	}
	for _, c := range panicCases {
		t.Run(fmt.Sprintf("(%d)C(%d)mod%d", c.n, c.k, modulus), func(t *testing.T) {
			defer func() {
				err := recover()
				if err != c.want {
					t.Errorf("wanted panic: %#v, got panic: %#v", c.want, err)
				}
			}()

			tables.Calc(c.n, c.k)
		})
	}
}
