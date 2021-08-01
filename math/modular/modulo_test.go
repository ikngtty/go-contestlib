package modular

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestModReg(t *testing.T) {
	const modulus = 13

	cases := []struct {
		n    int
		want int
	}{
		{0, 0}, {2, 2}, {16, 3}, {30, 4},
		{-3, 10}, {-17, 9}, {-31, 8},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("(%d)mod%d", c.n, modulus), func(t *testing.T) {
			m := NewMod(modulus)
			got := m.Reg(c.n)
			if got != c.want {
				t.Errorf("want: %d, got: %d", c.want, got)
			}
		})
	}
}

func TestModInv(t *testing.T) {
	cases := []struct {
		modulus int
		n       int
		want    int
	}{
		{13, 1, 1}, {13, 2, 7}, {13, 3, 9}, {13, 4, 10}, {13, 5, 8}, {13, 6, 11},
		{13, 7, 2}, {13, 8, 5}, {13, 9, 3}, {13, 10, 4}, {13, 11, 6}, {13, 12, 12},
		{13, -1, 12}, {13, -2, 6}, {13, 14, 1}, {13, 15, 7},
		{10, 3, 7},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("(%d)^(-1)mod%d", c.n, c.modulus), func(t *testing.T) {
			m := NewMod(c.modulus)
			got := m.Inv(c.n)
			if got != c.want {
				t.Errorf("want: %d, got: %d", c.want, got)
			}
		})
	}
}

func TestModInvs(t *testing.T) {
	const modulus = 13

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
			m := NewMod(modulus)
			got := m.Invs(c.n)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("want: %v, got: %v", c.want, got)
			}
		})
	}
}
