package intmath

import (
	"fmt"
	"testing"
)

func TestAbs64(t *testing.T) {
	cases := []struct {
		name string
		x    int64
		want int64
	}{
		{"positive", 5, 5},
		{"zero", 0, 0},
		{"negative", -5, 5},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Abs64(c.x)
			if got != c.want {
				t.Errorf("want: %d, got: %d", c.want, got)
			}
		})
	}
}

func TestMin64(t *testing.T) {
	cases := []struct {
		name   string
		values []int64
		want   int64
	}{
		{"1 item", []int64{10}, 10},
		{"2 items 1st", []int64{10, 20}, 10},
		{"2 items 2nd", []int64{20, 10}, 10},
		{"3 items 1st", []int64{10, 20, 30}, 10},
		{"3 items 2nd", []int64{30, 10, 20}, 10},
		{"3 items 3rd", []int64{20, 30, 10}, 10},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Min64(c.values...)
			if got != c.want {
				t.Errorf("want: %d, got: %d", c.want, got)
			}
		})
	}
}

func TestMax64(t *testing.T) {
	cases := []struct {
		name   string
		values []int64
		want   int64
	}{
		{"1 item", []int64{10}, 10},
		{"2 items 1st", []int64{20, 10}, 20},
		{"2 items 2nd", []int64{10, 20}, 20},
		{"3 items 1st", []int64{30, 20, 10}, 30},
		{"3 items 2nd", []int64{10, 30, 20}, 30},
		{"3 items 3rd", []int64{20, 10, 30}, 30},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Max64(c.values...)
			if got != c.want {
				t.Errorf("want: %d, got: %d", c.want, got)
			}
		})
	}
}

func TestPow64(t *testing.T) {
	cases := []struct {
		base     int64
		exponent uint
		want     int64
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
			got := Pow64(c.base, c.exponent)
			if got != c.want {
				t.Errorf("want: %d, got: %d", c.want, got)
			}
		})
	}
}
