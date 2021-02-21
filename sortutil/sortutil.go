package sortutil

import (
	"sort"
)

// ReverseInts sorts a reversely.
func ReverseInts(a []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
}

// ReverseStrings sorts a reversely.
func ReverseStrings(a []string) {
	sort.Sort(sort.Reverse(sort.StringSlice(a)))
}
