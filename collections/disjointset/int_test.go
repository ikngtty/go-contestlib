package disjointset

import (
	"strconv"
	"testing"
)

func TestSame(t *testing.T) {
	const n = 5
	cases := []struct {
		merges [][]int
		x      int
		y      int
		want   bool
	}{
		{[][]int{}, 0, 0, true},
		{[][]int{}, 0, 1, false},
		{[][]int{{0, 1}}, 0, 1, true},
		{[][]int{{0, 1}}, 0, 2, false},
		{[][]int{{0, 1}, {1, 2}}, 0, 2, true},
		{[][]int{{0, 1}, {2, 3}}, 0, 2, false},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			f := NewDisjointSetForest(n)
			for i := 0; i < n; i++ {
				f.Add(i)
			}
			for _, merge := range c.merges {
				f.Merge(merge[0], merge[1])
			}
			got := f.Same(c.x, c.y)
			if got != c.want {
				t.Errorf("want: %v, got: %v", c.want, got)
			}
		})
	}
}

func TestCount(t *testing.T) {
	const n = 10
	const addCount = 5
	cases := []struct {
		merges [][]int
		want   int
	}{
		{[][]int{}, 5},
		{[][]int{{0, 1}}, 4},
		{[][]int{{0, 1}, {1, 2}}, 3},
		{[][]int{{0, 1}, {2, 3}}, 3},
		{[][]int{{0, 1}, {2, 3}, {1, 2}}, 2},
		{[][]int{{0, 1}, {2, 3}, {1, 2}, {1, 3}}, 2},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			f := NewDisjointSetForest(n)
			for i := 0; i < addCount; i++ {
				f.Add(i)
			}
			for _, merge := range c.merges {
				f.Merge(merge[0], merge[1])
			}
			got := f.Count()
			if got != c.want {
				t.Errorf("want: %d, got: %d", c.want, got)
			}
		})
	}
}
