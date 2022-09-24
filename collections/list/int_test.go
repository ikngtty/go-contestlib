package list

import (
	"reflect"
	"testing"
)

func TestNewIntList(t *testing.T) {
	want := newIntList0()
	got := NewIntList()
	assertEqual(t, want, got)
}

func TestNewIntListFromArray(t *testing.T) {
	cases := []struct {
		name string
		a    []int
		want *IntList
	}{
		{"Length0", []int{}, newIntList0()},
		{"Length1", []int{10}, newIntList1(10)},
		{"Length2", []int{10, 20}, newIntList2(10, 20)},
		{"Length3", []int{10, 20, 30}, newIntList3(10, 20, 30)},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := NewIntListFromArray(c.a)
			assertEqual(t, c.want, got)
		})
	}
}

func TestIntListPush(t *testing.T) {
	got := newIntList0()

	got.Push(10)
	t.Run("1st push", func(t *testing.T) {
		assertEqual(t, newIntList1(10), got)
	})

	got.Push(20)
	t.Run("2nd push", func(t *testing.T) {
		assertEqual(t, newIntList2(10, 20), got)
	})

	got.Push(30)
	t.Run("3rd push", func(t *testing.T) {
		assertEqual(t, newIntList3(10, 20, 30), got)
	})
}

func TestIntListPushLeft(t *testing.T) {
	got := newIntList0()

	got.PushLeft(10)
	t.Run("1st push", func(t *testing.T) {
		assertEqual(t, newIntList1(10), got)
	})

	got.PushLeft(20)
	t.Run("2nd push", func(t *testing.T) {
		assertEqual(t, newIntList2(20, 10), got)
	})

	got.PushLeft(30)
	t.Run("3rd push", func(t *testing.T) {
		assertEqual(t, newIntList3(30, 20, 10), got)
	})
}

func TestIntListPop(t *testing.T) {
	got := newIntList3(10, 20, 30)

	value := got.Pop()
	t.Run("1st pop", func(t *testing.T) {
		assertEqual(t, 30, value)
		assertEqual(t, newIntList2(10, 20), got)
	})

	value = got.Pop()
	t.Run("2nd pop", func(t *testing.T) {
		assertEqual(t, 20, value)
		assertEqual(t, newIntList1(10), got)
	})

	value = got.Pop()
	t.Run("3rd pop", func(t *testing.T) {
		assertEqual(t, 10, value)
		assertEqual(t, newIntList0(), got)
	})
}

func TestIntListPopLeft(t *testing.T) {
	got := newIntList3(10, 20, 30)

	value := got.PopLeft()
	t.Run("1st pop", func(t *testing.T) {
		assertEqual(t, 10, value)
		assertEqual(t, newIntList2(20, 30), got)
	})

	value = got.PopLeft()
	t.Run("2nd pop", func(t *testing.T) {
		assertEqual(t, 20, value)
		assertEqual(t, newIntList1(30), got)
	})

	value = got.PopLeft()
	t.Run("3rd pop", func(t *testing.T) {
		assertEqual(t, 30, value)
		assertEqual(t, newIntList0(), got)
	})
}

func TestIntListConcat(t *testing.T) {
	cases := []struct {
		name  string
		left  *IntList
		right *IntList
		want  *IntList
	}{
		{"Length0+Length0",
			newIntList0(), newIntList0(),
			newIntList0()},
		{"Length0+Length1",
			newIntList0(), newIntList1(100),
			newIntList1(100)},
		{"Length0+Length2",
			newIntList0(), newIntList2(100, 200),
			newIntList2(100, 200)},
		{"Length1+Length0",
			newIntList1(10), newIntList0(),
			newIntList1(10)},
		{"Length1+Length1",
			newIntList1(10), newIntList1(100),
			newIntList2(10, 100)},
		{"Length1+Length2",
			newIntList1(10), newIntList2(100, 200),
			newIntList3(10, 100, 200)},
		{"Length2+Length0",
			newIntList2(10, 20), newIntList0(),
			newIntList2(10, 20)},
		{"Length2+Length1",
			newIntList2(10, 20), newIntList1(100),
			newIntList3(10, 20, 100)},
		{"Length2+Length2",
			newIntList2(10, 20), newIntList2(100, 200),
			newIntList4(10, 20, 100, 200)},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.left.Concat(c.right)
			assertEqual(t, c.want, c.left)
			assertEqual(t, c.want, c.right)
		})
	}
}

func TestIntListEach(t *testing.T) {
	cases := []struct {
		name string
		list *IntList
		want []int
	}{
		{"Length0", newIntList0(), []int{}},
		{"Length1", newIntList1(10), []int{10}},
		{"Length2", newIntList2(10, 20), []int{10, 20}},
		{"Length3", newIntList3(10, 20, 30), []int{10, 20, 30}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := make([]int, 0)
			c.list.Each(func(elem int) {
				got = append(got, elem)
			})
			assertEqual(t, c.want, got)
		})
	}
}

func TestIntListToA(t *testing.T) {
	cases := []struct {
		name string
		list *IntList
		want []int
	}{
		{"Length0", newIntList0(), []int{}},
		{"Length1", newIntList1(10), []int{10}},
		{"Length2", newIntList2(10, 20), []int{10, 20}},
		{"Length3", newIntList3(10, 20, 30), []int{10, 20, 30}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.list.ToA()
			assertEqual(t, c.want, got)
		})
	}
}

func assertEqual(t *testing.T, want, got interface{}) {
	if !reflect.DeepEqual(got, want) {
		// TODO: more detail print (`%#v` can show a pointer value but
		// cannot show the value that a pointer refer to)
		t.Errorf("want: %#v, got: %#v", want, got)
	}
}

func newIntList0() *IntList {
	return &IntList{first: nil, last: nil, len: 0}
}

func newIntList1(value1 int) *IntList {
	node1 := intListNode{parent: nil, child: nil, value: value1}
	return &IntList{first: &node1, last: &node1, len: 1}
}

func newIntList2(value1, value2 int) *IntList {
	node1 := intListNode{parent: nil, child: nil, value: value1}
	node2 := intListNode{parent: &node1, child: nil, value: value2}
	node1.child = &node2
	return &IntList{first: &node1, last: &node2, len: 2}
}

func newIntList3(value1, value2, value3 int) *IntList {
	node1 := intListNode{parent: nil, child: nil, value: value1}
	node2 := intListNode{parent: &node1, child: nil, value: value2}
	node1.child = &node2
	node3 := intListNode{parent: &node2, child: nil, value: value3}
	node2.child = &node3
	return &IntList{first: &node1, last: &node3, len: 3}
}

func newIntList4(value1, value2, value3, value4 int) *IntList {
	node1 := intListNode{parent: nil, child: nil, value: value1}
	node2 := intListNode{parent: &node1, child: nil, value: value2}
	node1.child = &node2
	node3 := intListNode{parent: &node2, child: nil, value: value3}
	node2.child = &node3
	node4 := intListNode{parent: &node3, child: nil, value: value4}
	node3.child = &node4
	return &IntList{first: &node1, last: &node4, len: 4}
}
