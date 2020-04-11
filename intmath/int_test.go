package intmath

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
		name   string
		values []int
		want   int
	}{
		{"1 item", []int{10}, 10},
		{"2 items 1st", []int{10, 20}, 10},
		{"2 items 2nd", []int{20, 10}, 10},
		{"3 items 1st", []int{10, 20, 30}, 10},
		{"3 items 2nd", []int{30, 10, 20}, 10},
		{"3 items 3rd", []int{20, 30, 10}, 10},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Min(c.values...)
			if got != c.want {
				t.Errorf("want: %d, got: %d", c.want, got)
			}
		})
	}
}

func TestMax(t *testing.T) {
	cases := []struct {
		name   string
		values []int
		want   int
	}{
		{"1 item", []int{10}, 10},
		{"2 items 1st", []int{20, 10}, 20},
		{"2 items 2nd", []int{10, 20}, 20},
		{"3 items 1st", []int{30, 20, 10}, 30},
		{"3 items 2nd", []int{10, 30, 20}, 30},
		{"3 items 3rd", []int{20, 10, 30}, 30},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Max(c.values...)
			if got != c.want {
				t.Errorf("want: %d, got: %d", c.want, got)
			}
		})
	}
}

func TestPow(t *testing.T) {
	cases := []struct {
		base     int
		exponent uint
		want     int
	}{
		{5, 0, 1},
		{5, 1, 5},
		{5, 2, 25},
		{5, 3, 125},
		{-5, 0, 1},
		{-5, 1, -5},
		{-5, 2, 25},
		{-5, 3, -125},
	}

	for _, c := range cases {
		caseName := fmt.Sprintf("(%d)^%d", c.base, c.exponent)
		t.Run(caseName, func(t *testing.T) {
			got := Pow(c.base, c.exponent)
			if got != c.want {
				t.Errorf("want: %d, got: %d", c.want, got)
			}
		})
	}
}
