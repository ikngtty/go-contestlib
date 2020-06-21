package list

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewIntListFromArray(t *testing.T) {
	// NOTE: Dependency
	// -   NewIntListFromArray()
	//     -   NewIntList()
	//     -   list.Add()
	// -   list.ToA()
	//     -   list.Len()
	//     -   list.Each()

	cases := []struct {
		a []int
	}{
		{[]int{}},
		{[]int{20}},
		{[]int{20, 30}},
		{[]int{20, 30, 10}},
	}
	for _, c := range cases {
		testName := fmt.Sprintf("from_%d_length_array", len(c.a))
		t.Run(testName, func(t *testing.T) {
			list := NewIntListFromArray(c.a)

			got := list.ToA()
			want := c.a
			if !reflect.DeepEqual(got, want) {
				t.Errorf("want: %#v, got: %#v", want, got)
			}
		})
	}
}

func TestIntListConcat(t *testing.T) {
	// NOTE: Dependency
	// -   list.Concat()
	// -   NewIntListFromArray()
	//     -   NewIntList()
	//     -   list.Add()
	// -   list.ToA()
	//     -   list.Len()
	//     -   list.Each()

	t.Run("the_list_argument_is_unchanged", func(t *testing.T) {
		list := NewIntListFromArray([]int{10, 20, 30})
		other := NewIntListFromArray([]int{40, 50, 60})

		list.Concat(other)

		got := other.ToA()
		want := []int{40, 50, 60}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want: %#v, got: %#v", want, got)
		}
	})

	cases := []struct {
		list  []int
		other []int
		want  []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{22}, []int{22}},
		{[]int{}, []int{22, 44}, []int{22, 44}},
		{[]int{}, []int{22, 44, 33}, []int{22, 44, 33}},
		{[]int{30}, []int{}, []int{30}},
		{[]int{30}, []int{22}, []int{30, 22}},
		{[]int{30}, []int{22, 44}, []int{30, 22, 44}},
		{[]int{30}, []int{22, 44, 33}, []int{30, 22, 44, 33}},
		{[]int{30, 10}, []int{}, []int{30, 10}},
		{[]int{30, 10}, []int{22}, []int{30, 10, 22}},
		{[]int{30, 10}, []int{22, 44}, []int{30, 10, 22, 44}},
		{[]int{30, 10}, []int{22, 44, 33}, []int{30, 10, 22, 44, 33}},
		{[]int{30, 10, 50}, []int{}, []int{30, 10, 50}},
		{[]int{30, 10, 50}, []int{22}, []int{30, 10, 50, 22}},
		{[]int{30, 10, 50}, []int{22, 44}, []int{30, 10, 50, 22, 44}},
		{[]int{30, 10, 50}, []int{22, 44, 33}, []int{30, 10, 50, 22, 44, 33}},
	}
	for _, c := range cases {
		testName := fmt.Sprintf("%v+%v", c.list, c.other)
		t.Run(testName, func(t *testing.T) {
			list := NewIntListFromArray(c.list)
			other := NewIntListFromArray(c.other)

			list.Concat(other)

			got := list.ToA()
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("want: %#v, got: %#v", c.want, got)
			}
		})
	}
}
