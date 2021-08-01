package combinatorics

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {
	cases := []struct {
		n, k int
		want [][]int
	}{
		{n: 0, k: 0, want: [][]int{{}}},
		{n: 3, k: 0, want: [][]int{{}}},
		{n: 3, k: 1, want: [][]int{{0}, {1}, {2}}},
		{n: 3, k: 3, want: [][]int{
			{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0},
		}},
		{n: 5, k: 3, want: [][]int{
			{0, 1, 2}, {0, 1, 3}, {0, 1, 4},
			{0, 2, 1}, {0, 2, 3}, {0, 2, 4},
			{0, 3, 1}, {0, 3, 2}, {0, 3, 4},
			{0, 4, 1}, {0, 4, 2}, {0, 4, 3},
			{1, 0, 2}, {1, 0, 3}, {1, 0, 4},
			{1, 2, 0}, {1, 2, 3}, {1, 2, 4},
			{1, 3, 0}, {1, 3, 2}, {1, 3, 4},
			{1, 4, 0}, {1, 4, 2}, {1, 4, 3},
			{2, 0, 1}, {2, 0, 3}, {2, 0, 4},
			{2, 1, 0}, {2, 1, 3}, {2, 1, 4},
			{2, 3, 0}, {2, 3, 1}, {2, 3, 4},
			{2, 4, 0}, {2, 4, 1}, {2, 4, 3},
			{3, 0, 1}, {3, 0, 2}, {3, 0, 4},
			{3, 1, 0}, {3, 1, 2}, {3, 1, 4},
			{3, 2, 0}, {3, 2, 1}, {3, 2, 4},
			{3, 4, 0}, {3, 4, 1}, {3, 4, 2},
			{4, 0, 1}, {4, 0, 2}, {4, 0, 3},
			{4, 1, 0}, {4, 1, 2}, {4, 1, 3},
			{4, 2, 0}, {4, 2, 1}, {4, 2, 3},
			{4, 3, 0}, {4, 3, 1}, {4, 3, 2},
		}},
	}
	for _, c := range cases {
		caseName := fmt.Sprintf("f(%d,%d)", c.n, c.k)
		t.Run(caseName, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("unwanted panic: %#v", err)
				}
			}()

			got := make([][]int, 0)
			Permutations(c.n, c.k, func(pattern []int) {
				patternClone := make([]int, len(pattern))
				copy(patternClone, pattern)
				got = append(got, patternClone)
			})

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("want: %v, got: %v", c.want, got)
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
	}
	for _, c := range panicCases {
		caseName := fmt.Sprintf("f(%d,%d)", c.n, c.k)
		t.Run(caseName, func(t *testing.T) {
			defer func() {
				err := recover()
				if err != c.want {
					t.Errorf("wanted panic: %#v, got panic: %#v", c.want, err)
				}
			}()

			Permutations(c.n, c.k, func(pattern []int) {})
		})
	}
}

func TestCombinations(t *testing.T) {
	cases := []struct {
		n, k int
		want [][]int
	}{
		{n: 0, k: 0, want: [][]int{{}}},
		{n: 3, k: 0, want: [][]int{{}}},
		{n: 3, k: 1, want: [][]int{{0}, {1}, {2}}},
		{n: 3, k: 3, want: [][]int{{0, 1, 2}}},
		{n: 6, k: 3, want: [][]int{
			{0, 1, 2}, {0, 1, 3}, {0, 1, 4}, {0, 1, 5},
			{0, 2, 3}, {0, 2, 4}, {0, 2, 5},
			{0, 3, 4}, {0, 3, 5},
			{0, 4, 5},
			{1, 2, 3}, {1, 2, 4}, {1, 2, 5},
			{1, 3, 4}, {1, 3, 5},
			{1, 4, 5},
			{2, 3, 4}, {2, 3, 5},
			{2, 4, 5},
			{3, 4, 5},
		}},
	}
	for _, c := range cases {
		caseName := fmt.Sprintf("f(%d,%d)", c.n, c.k)
		t.Run(caseName, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("unwanted panic: %#v", err)
				}
			}()

			got := make([][]int, 0)
			Combinations(c.n, c.k, func(pattern []int) {
				patternClone := make([]int, len(pattern))
				copy(patternClone, pattern)
				got = append(got, patternClone)
			})

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("want: %v, got: %v", c.want, got)
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
	}
	for _, c := range panicCases {
		caseName := fmt.Sprintf("f(%d,%d)", c.n, c.k)
		t.Run(caseName, func(t *testing.T) {
			defer func() {
				err := recover()
				if err != c.want {
					t.Errorf("wanted panic: %#v, got panic: %#v", c.want, err)
				}
			}()

			Combinations(c.n, c.k, func(pattern []int) {})
		})
	}
}

func TestDupPermutations(t *testing.T) {
	cases := []struct {
		n, k int
		want [][]int
	}{
		{n: 0, k: 0, want: [][]int{{}}},
		{n: 3, k: 0, want: [][]int{{}}},
		{n: 3, k: 1, want: [][]int{{0}, {1}, {2}}},
		{n: 3, k: 3, want: [][]int{
			{0, 0, 0}, {0, 0, 1}, {0, 0, 2},
			{0, 1, 0}, {0, 1, 1}, {0, 1, 2},
			{0, 2, 0}, {0, 2, 1}, {0, 2, 2},
			{1, 0, 0}, {1, 0, 1}, {1, 0, 2},
			{1, 1, 0}, {1, 1, 1}, {1, 1, 2},
			{1, 2, 0}, {1, 2, 1}, {1, 2, 2},
			{2, 0, 0}, {2, 0, 1}, {2, 0, 2},
			{2, 1, 0}, {2, 1, 1}, {2, 1, 2},
			{2, 2, 0}, {2, 2, 1}, {2, 2, 2},
		}},
		{n: 4, k: 3, want: [][]int{
			{0, 0, 0}, {0, 0, 1}, {0, 0, 2}, {0, 0, 3},
			{0, 1, 0}, {0, 1, 1}, {0, 1, 2}, {0, 1, 3},
			{0, 2, 0}, {0, 2, 1}, {0, 2, 2}, {0, 2, 3},
			{0, 3, 0}, {0, 3, 1}, {0, 3, 2}, {0, 3, 3},
			{1, 0, 0}, {1, 0, 1}, {1, 0, 2}, {1, 0, 3},
			{1, 1, 0}, {1, 1, 1}, {1, 1, 2}, {1, 1, 3},
			{1, 2, 0}, {1, 2, 1}, {1, 2, 2}, {1, 2, 3},
			{1, 3, 0}, {1, 3, 1}, {1, 3, 2}, {1, 3, 3},
			{2, 0, 0}, {2, 0, 1}, {2, 0, 2}, {2, 0, 3},
			{2, 1, 0}, {2, 1, 1}, {2, 1, 2}, {2, 1, 3},
			{2, 2, 0}, {2, 2, 1}, {2, 2, 2}, {2, 2, 3},
			{2, 3, 0}, {2, 3, 1}, {2, 3, 2}, {2, 3, 3},
			{3, 0, 0}, {3, 0, 1}, {3, 0, 2}, {3, 0, 3},
			{3, 1, 0}, {3, 1, 1}, {3, 1, 2}, {3, 1, 3},
			{3, 2, 0}, {3, 2, 1}, {3, 2, 2}, {3, 2, 3},
			{3, 3, 0}, {3, 3, 1}, {3, 3, 2}, {3, 3, 3},
		}},
	}
	for _, c := range cases {
		caseName := fmt.Sprintf("f(%d,%d)", c.n, c.k)
		t.Run(caseName, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("unwanted panic: %#v", err)
				}
			}()

			got := make([][]int, 0)
			DupPermutations(c.n, c.k, func(pattern []int) {
				patternClone := make([]int, len(pattern))
				copy(patternClone, pattern)
				got = append(got, patternClone)
			})

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("want: %v, got: %v", c.want, got)
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
	}
	for _, c := range panicCases {
		caseName := fmt.Sprintf("f(%d,%d)", c.n, c.k)
		t.Run(caseName, func(t *testing.T) {
			defer func() {
				err := recover()
				if err != c.want {
					t.Errorf("wanted panic: %#v, got panic: %#v", c.want, err)
				}
			}()

			DupPermutations(c.n, c.k, func(pattern []int) {})
		})
	}
}

func TestDupCombinations(t *testing.T) {
	cases := []struct {
		n, k int
		want [][]int
	}{
		{n: 0, k: 0, want: [][]int{{}}},
		{n: 3, k: 0, want: [][]int{{}}},
		{n: 3, k: 1, want: [][]int{{0}, {1}, {2}}},
		{n: 3, k: 3, want: [][]int{
			{0, 0, 0}, {0, 0, 1}, {0, 0, 2},
			{0, 1, 1}, {0, 1, 2},
			{0, 2, 2},
			{1, 1, 1}, {1, 1, 2},
			{1, 2, 2},
			{2, 2, 2},
		}},
		{n: 5, k: 3, want: [][]int{
			{0, 0, 0}, {0, 0, 1}, {0, 0, 2}, {0, 0, 3}, {0, 0, 4},
			{0, 1, 1}, {0, 1, 2}, {0, 1, 3}, {0, 1, 4},
			{0, 2, 2}, {0, 2, 3}, {0, 2, 4},
			{0, 3, 3}, {0, 3, 4},
			{0, 4, 4},
			{1, 1, 1}, {1, 1, 2}, {1, 1, 3}, {1, 1, 4},
			{1, 2, 2}, {1, 2, 3}, {1, 2, 4},
			{1, 3, 3}, {1, 3, 4},
			{1, 4, 4},
			{2, 2, 2}, {2, 2, 3}, {2, 2, 4},
			{2, 3, 3}, {2, 3, 4},
			{2, 4, 4},
			{3, 3, 3}, {3, 3, 4},
			{3, 4, 4},
			{4, 4, 4},
		}},
	}
	for _, c := range cases {
		caseName := fmt.Sprintf("f(%d,%d)", c.n, c.k)
		t.Run(caseName, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("unwanted panic: %#v", err)
				}
			}()

			got := make([][]int, 0)
			DupCombinations(c.n, c.k, func(pattern []int) {
				patternClone := make([]int, len(pattern))
				copy(patternClone, pattern)
				got = append(got, patternClone)
			})

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("want: %v, got: %v", c.want, got)
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
	}
	for _, c := range panicCases {
		caseName := fmt.Sprintf("f(%d,%d)", c.n, c.k)
		t.Run(caseName, func(t *testing.T) {
			defer func() {
				err := recover()
				if err != c.want {
					t.Errorf("wanted panic: %#v, got panic: %#v", c.want, err)
				}
			}()

			DupCombinations(c.n, c.k, func(pattern []int) {})
		})
	}
}

func TestBitPatterns(t *testing.T) {
	cases := []struct {
		bitsLen int
		want    []string
	}{
		{0, []string{""}},
		{1, []string{"0", "1"}},
		{2, []string{"00", "01", "10", "11"}},
		{3, []string{
			"000", "001", "010", "011",
			"100", "101", "110", "111",
		}},
	}
	for _, c := range cases {
		caseName := fmt.Sprintf("f(%d)", c.bitsLen)
		t.Run(caseName, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("unwanted panic: %#v", err)
				}
			}()

			want := make([][]bool, len(c.want))
			for iBits, strBits := range c.want {
				bits := make([]bool, len(strBits))
				for iChar, charBit := range strBits {
					bits[iChar] = charBit == '1'
				}
				want[iBits] = bits
			}

			got := make([][]bool, 0)
			BitPatterns(c.bitsLen, func(bits []bool) {
				bitsClone := make([]bool, len(bits))
				copy(bitsClone, bits)
				got = append(got, bitsClone)
			})

			if !reflect.DeepEqual(got, want) {
				t.Errorf("want: %v, got: %v", want, got)
			}
		})
	}

	panicCases := []struct {
		bitsLen int
		want    string
	}{
		{-1, "invalid bitsLen: -1"},
	}
	for _, c := range panicCases {
		caseName := fmt.Sprintf("f(%d)", c.bitsLen)
		t.Run(caseName, func(t *testing.T) {
			defer func() {
				err := recover()
				if err != c.want {
					t.Errorf("wanted panic: %#v, got panic: %#v", c.want, err)
				}
			}()

			BitPatterns(c.bitsLen, func(bits []bool) {})
		})
	}
}
