package combinatorics

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {
	testCases := []struct {
		n, k int
		want [][]int
	}{
		{n: 0, k: 0, want: [][]int{[]int{}}},
		{n: 3, k: 0, want: [][]int{[]int{}}},
		{n: 3, k: 1, want: [][]int{[]int{0}, []int{1}, []int{2}}},
		{n: 5, k: 3, want: [][]int{
			[]int{0, 1, 2},
			[]int{0, 1, 3},
			[]int{0, 1, 4},
			[]int{0, 2, 1},
			[]int{0, 2, 3},
			[]int{0, 2, 4},
			[]int{0, 3, 1},
			[]int{0, 3, 2},
			[]int{0, 3, 4},
			[]int{0, 4, 1},
			[]int{0, 4, 2},
			[]int{0, 4, 3},
			[]int{1, 0, 2},
			[]int{1, 0, 3},
			[]int{1, 0, 4},
			[]int{1, 2, 0},
			[]int{1, 2, 3},
			[]int{1, 2, 4},
			[]int{1, 3, 0},
			[]int{1, 3, 2},
			[]int{1, 3, 4},
			[]int{1, 4, 0},
			[]int{1, 4, 2},
			[]int{1, 4, 3},
			[]int{2, 0, 1},
			[]int{2, 0, 3},
			[]int{2, 0, 4},
			[]int{2, 1, 0},
			[]int{2, 1, 3},
			[]int{2, 1, 4},
			[]int{2, 3, 0},
			[]int{2, 3, 1},
			[]int{2, 3, 4},
			[]int{2, 4, 0},
			[]int{2, 4, 1},
			[]int{2, 4, 3},
			[]int{3, 0, 1},
			[]int{3, 0, 2},
			[]int{3, 0, 4},
			[]int{3, 1, 0},
			[]int{3, 1, 2},
			[]int{3, 1, 4},
			[]int{3, 2, 0},
			[]int{3, 2, 1},
			[]int{3, 2, 4},
			[]int{3, 4, 0},
			[]int{3, 4, 1},
			[]int{3, 4, 2},
			[]int{4, 0, 1},
			[]int{4, 0, 2},
			[]int{4, 0, 3},
			[]int{4, 1, 0},
			[]int{4, 1, 2},
			[]int{4, 1, 3},
			[]int{4, 2, 0},
			[]int{4, 2, 1},
			[]int{4, 2, 3},
			[]int{4, 3, 0},
			[]int{4, 3, 1},
			[]int{4, 3, 2},
		}},
	}

	for _, testCase := range testCases {
		testCaseName := fmt.Sprintf("f(%d,%d)", testCase.n, testCase.k)
		t.Run(testCaseName, func(t *testing.T) {
			got := make([][]int, 0)
			Permutations(testCase.n, testCase.k, func(pattern []int) {
				patternClone := make([]int, len(pattern))
				copy(patternClone, pattern)
				got = append(got, patternClone)
			})

			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("want: %v, got: %v", testCase.want, got)
			}
		})
	}
}

func TestCombinations(t *testing.T) {
	testCases := []struct {
		n, k int
		want [][]int
	}{
		{n: 0, k: 0, want: [][]int{[]int{}}},
		{n: 3, k: 0, want: [][]int{[]int{}}},
		{n: 3, k: 1, want: [][]int{[]int{0}, []int{1}, []int{2}}},
		{n: 3, k: 3, want: [][]int{[]int{0, 1, 2}}},
		{n: 6, k: 3, want: [][]int{
			[]int{0, 1, 2},
			[]int{0, 1, 3},
			[]int{0, 1, 4},
			[]int{0, 1, 5},
			[]int{0, 2, 3},
			[]int{0, 2, 4},
			[]int{0, 2, 5},
			[]int{0, 3, 4},
			[]int{0, 3, 5},
			[]int{0, 4, 5},
			[]int{1, 2, 3},
			[]int{1, 2, 4},
			[]int{1, 2, 5},
			[]int{1, 3, 4},
			[]int{1, 3, 5},
			[]int{1, 4, 5},
			[]int{2, 3, 4},
			[]int{2, 3, 5},
			[]int{2, 4, 5},
			[]int{3, 4, 5},
		}},
	}

	for _, testCase := range testCases {
		testCaseName := fmt.Sprintf("f(%d,%d)", testCase.n, testCase.k)
		t.Run(testCaseName, func(t *testing.T) {
			got := make([][]int, 0)
			Combinations(testCase.n, testCase.k, func(pattern []int) {
				patternClone := make([]int, len(pattern))
				copy(patternClone, pattern)
				got = append(got, patternClone)
			})

			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("want: %v, got: %v", testCase.want, got)
			}
		})
	}
}

func TestDupPermutations(t *testing.T) {
	testCases := []struct {
		n, k int
		want [][]int
	}{
		{n: 0, k: 0, want: [][]int{[]int{}}},
		{n: 3, k: 0, want: [][]int{[]int{}}},
		{n: 3, k: 1, want: [][]int{[]int{0}, []int{1}, []int{2}}},
		{n: 4, k: 3, want: [][]int{
			[]int{0, 0, 0},
			[]int{0, 0, 1},
			[]int{0, 0, 2},
			[]int{0, 0, 3},
			[]int{0, 1, 0},
			[]int{0, 1, 1},
			[]int{0, 1, 2},
			[]int{0, 1, 3},
			[]int{0, 2, 0},
			[]int{0, 2, 1},
			[]int{0, 2, 2},
			[]int{0, 2, 3},
			[]int{0, 3, 0},
			[]int{0, 3, 1},
			[]int{0, 3, 2},
			[]int{0, 3, 3},
			[]int{1, 0, 0},
			[]int{1, 0, 1},
			[]int{1, 0, 2},
			[]int{1, 0, 3},
			[]int{1, 1, 0},
			[]int{1, 1, 1},
			[]int{1, 1, 2},
			[]int{1, 1, 3},
			[]int{1, 2, 0},
			[]int{1, 2, 1},
			[]int{1, 2, 2},
			[]int{1, 2, 3},
			[]int{1, 3, 0},
			[]int{1, 3, 1},
			[]int{1, 3, 2},
			[]int{1, 3, 3},
			[]int{2, 0, 0},
			[]int{2, 0, 1},
			[]int{2, 0, 2},
			[]int{2, 0, 3},
			[]int{2, 1, 0},
			[]int{2, 1, 1},
			[]int{2, 1, 2},
			[]int{2, 1, 3},
			[]int{2, 2, 0},
			[]int{2, 2, 1},
			[]int{2, 2, 2},
			[]int{2, 2, 3},
			[]int{2, 3, 0},
			[]int{2, 3, 1},
			[]int{2, 3, 2},
			[]int{2, 3, 3},
			[]int{3, 0, 0},
			[]int{3, 0, 1},
			[]int{3, 0, 2},
			[]int{3, 0, 3},
			[]int{3, 1, 0},
			[]int{3, 1, 1},
			[]int{3, 1, 2},
			[]int{3, 1, 3},
			[]int{3, 2, 0},
			[]int{3, 2, 1},
			[]int{3, 2, 2},
			[]int{3, 2, 3},
			[]int{3, 3, 0},
			[]int{3, 3, 1},
			[]int{3, 3, 2},
			[]int{3, 3, 3},
		}},
	}

	for _, testCase := range testCases {
		testCaseName := fmt.Sprintf("f(%d,%d)", testCase.n, testCase.k)
		t.Run(testCaseName, func(t *testing.T) {
			got := make([][]int, 0)
			DupPermutations(testCase.n, testCase.k, func(pattern []int) {
				patternClone := make([]int, len(pattern))
				copy(patternClone, pattern)
				got = append(got, patternClone)
			})

			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("want: %v, got: %v", testCase.want, got)
			}
		})
	}
}

func TestDupCombinations(t *testing.T) {
	testCases := []struct {
		n, k int
		want [][]int
	}{
		{n: 0, k: 0, want: [][]int{[]int{}}},
		{n: 3, k: 0, want: [][]int{[]int{}}},
		{n: 3, k: 1, want: [][]int{[]int{0}, []int{1}, []int{2}}},
		{n: 3, k: 3, want: [][]int{
			[]int{0, 0, 0},
			[]int{0, 0, 1},
			[]int{0, 0, 2},
			[]int{0, 1, 1},
			[]int{0, 1, 2},
			[]int{0, 2, 2},
			[]int{1, 1, 1},
			[]int{1, 1, 2},
			[]int{1, 2, 2},
			[]int{2, 2, 2},
		}},
		{n: 5, k: 3, want: [][]int{
			[]int{0, 0, 0},
			[]int{0, 0, 1},
			[]int{0, 0, 2},
			[]int{0, 0, 3},
			[]int{0, 0, 4},
			[]int{0, 1, 1},
			[]int{0, 1, 2},
			[]int{0, 1, 3},
			[]int{0, 1, 4},
			[]int{0, 2, 2},
			[]int{0, 2, 3},
			[]int{0, 2, 4},
			[]int{0, 3, 3},
			[]int{0, 3, 4},
			[]int{0, 4, 4},
			[]int{1, 1, 1},
			[]int{1, 1, 2},
			[]int{1, 1, 3},
			[]int{1, 1, 4},
			[]int{1, 2, 2},
			[]int{1, 2, 3},
			[]int{1, 2, 4},
			[]int{1, 3, 3},
			[]int{1, 3, 4},
			[]int{1, 4, 4},
			[]int{2, 2, 2},
			[]int{2, 2, 3},
			[]int{2, 2, 4},
			[]int{2, 3, 3},
			[]int{2, 3, 4},
			[]int{2, 4, 4},
			[]int{3, 3, 3},
			[]int{3, 3, 4},
			[]int{3, 4, 4},
			[]int{4, 4, 4},
		}},
	}

	for _, testCase := range testCases {
		testCaseName := fmt.Sprintf("f(%d,%d)", testCase.n, testCase.k)
		t.Run(testCaseName, func(t *testing.T) {
			got := make([][]int, 0)
			DupCombinations(testCase.n, testCase.k, func(pattern []int) {
				patternClone := make([]int, len(pattern))
				copy(patternClone, pattern)
				got = append(got, patternClone)
			})

			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("want: %v, got: %v", testCase.want, got)
			}
		})
	}
}

func TestBitPatterns(t *testing.T) {
	testCases := []struct {
		bitsLen int
		want    []string
	}{
		{0, []string{""}},
		{1, []string{"0", "1"}},
		{2, []string{"00", "01", "10", "11"}},
		{3, []string{"000", "001", "010", "011", "100", "101", "110", "111"}},
	}
	for _, testCase := range testCases {
		testCaseName := fmt.Sprintf("f(%d)", testCase.bitsLen)
		t.Run(testCaseName, func(t *testing.T) {
			want := make([][]bool, len(testCase.want))
			for iBits, strBits := range testCase.want {
				bits := make([]bool, len(strBits))
				for iChar, charBit := range strBits {
					bits[iChar] = charBit == '1'
				}
				want[iBits] = bits
			}

			got := make([][]bool, 0)
			BitPatterns(testCase.bitsLen, func(bits []bool) {
				bitsClone := make([]bool, len(bits))
				copy(bitsClone, bits)
				got = append(got, bitsClone)
			})

			if !reflect.DeepEqual(got, want) {
				t.Errorf("want: %v, got: %v", want, got)
			}
		})
	}
}
