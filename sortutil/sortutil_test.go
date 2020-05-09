package sortutil

import (
	"reflect"
	"testing"
)

func TestReverseInts(t *testing.T) {
	a := []int{20, 10, 100, 30}
	ReverseInts(a)

	want := []int{100, 30, 20, 10}
	if !reflect.DeepEqual(a, want) {
		t.Errorf("want: %#v, got: %#v", want, a)
	}
}

func TestReverseStrings(t *testing.T) {
	a := []string{"def", "abc", "abcab", "ghi", "abca"}
	ReverseStrings(a)

	want := []string{"ghi", "def", "abcab", "abca", "abc"}
	if !reflect.DeepEqual(a, want) {
		t.Errorf("want: %#v, got: %#v", want, a)
	}
}
