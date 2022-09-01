package disjointset

import (
	"strconv"
	"testing"
)

func Test(t *testing.T) {
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
