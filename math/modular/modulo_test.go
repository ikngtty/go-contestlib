package modular

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestReg(t *testing.T) {
	const modulus = 13
	m := NewMod(modulus)

	cases := []struct {
		n    int
		want int
	}{
		{0, 0}, {2, 2}, {16, 3}, {30, 4},
		{-3, 10}, {-17, 9}, {-31, 8},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("(%d)mod%d", c.n, modulus), func(t *testing.T) {
			got := m.Reg(c.n)
			if got != c.want {
				t.Errorf("want: %d, got: %d", c.want, got)
			}
		})
	}
}

func TestInv(t *testing.T) {
	const modulus = 13
	m := NewMod(modulus)

	cases := []struct {
		n    int
		want int
	}{
		{1, 1}, {2, 7}, {3, 9}, {4, 10}, {5, 8}, {6, 11},
		{7, 2}, {8, 5}, {9, 3}, {10, 4}, {11, 6}, {12, 12},
		{-1, 12}, {-2, 6}, {14, 1}, {15, 7},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("(%d)^(-1)mod%d", c.n, modulus), func(t *testing.T) {
			got := m.Inv(c.n)
			if got != c.want {
				t.Errorf("want: %d, got: %d", c.want, got)
			}
		})
	}
}

func TestInvs(t *testing.T) {
	const modulus = 13
	m := NewMod(modulus)

	cases := []struct {
		n    int
		want []int
	}{
		{0, []int{}},
		{1, []int{0}},
		{2, []int{0, 1}},
		{4, []int{0, 1, 7, 9}},
		{13, []int{0, 1, 7, 9, 10, 8, 11, 2, 5, 3, 4, 6, 12}},
	}
	for _, c := range cases {
		t.Run(strconv.Itoa(c.n), func(t *testing.T) {
			got := m.Invs(c.n)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("want: %v, got: %v", c.want, got)
			}
		})
	}
}
