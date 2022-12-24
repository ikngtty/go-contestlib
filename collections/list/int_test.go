package list

import (
	"reflect"
	"testing"
)

func TestNewListInt(t *testing.T) {
	want := newListInt0()
	got := NewListInt()
	assertEqual(t, want, got)
}

func TestNewListIntFromArray(t *testing.T) {
	cases := []struct {
		name string
		a    []int
		want *ListInt
	}{
		{"Length0", []int{}, newListInt0()},
		{"Length1", []int{10}, newListInt1(10)},
		{"Length2", []int{10, 20}, newListInt2(10, 20)},
		{"Length3", []int{10, 20, 30}, newListInt3(10, 20, 30)},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := NewListIntFromArray(c.a)
			assertEqual(t, c.want, got)
		})
	}
}

func TestListIntPush(t *testing.T) {
	got := newListInt0()

	got.Push(10)
	t.Run("1st push", func(t *testing.T) {
		assertEqual(t, newListInt1(10), got)
	})

	got.Push(20)
	t.Run("2nd push", func(t *testing.T) {
		assertEqual(t, newListInt2(10, 20), got)
	})

	got.Push(30)
	t.Run("3rd push", func(t *testing.T) {
		assertEqual(t, newListInt3(10, 20, 30), got)
	})
}

func TestListIntPushLeft(t *testing.T) {
	got := newListInt0()

	got.PushLeft(10)
	t.Run("1st push", func(t *testing.T) {
		assertEqual(t, newListInt1(10), got)
	})

	got.PushLeft(20)
	t.Run("2nd push", func(t *testing.T) {
		assertEqual(t, newListInt2(20, 10), got)
	})

	got.PushLeft(30)
	t.Run("3rd push", func(t *testing.T) {
		assertEqual(t, newListInt3(30, 20, 10), got)
	})
}

func TestListIntPop(t *testing.T) {
	got := newListInt3(10, 20, 30)

	value := got.Pop()
	t.Run("1st pop", func(t *testing.T) {
		assertEqual(t, 30, value)
		assertEqual(t, newListInt2(10, 20), got)
	})

	value = got.Pop()
	t.Run("2nd pop", func(t *testing.T) {
		assertEqual(t, 20, value)
		assertEqual(t, newListInt1(10), got)
	})

	value = got.Pop()
	t.Run("3rd pop", func(t *testing.T) {
		assertEqual(t, 10, value)
		assertEqual(t, newListInt0(), got)
	})
}

func TestListIntPopLeft(t *testing.T) {
	got := newListInt3(10, 20, 30)

	value := got.PopLeft()
	t.Run("1st pop", func(t *testing.T) {
		assertEqual(t, 10, value)
		assertEqual(t, newListInt2(20, 30), got)
	})

	value = got.PopLeft()
	t.Run("2nd pop", func(t *testing.T) {
		assertEqual(t, 20, value)
		assertEqual(t, newListInt1(30), got)
	})

	value = got.PopLeft()
	t.Run("3rd pop", func(t *testing.T) {
		assertEqual(t, 30, value)
		assertEqual(t, newListInt0(), got)
	})
}

func TestListIntConcat(t *testing.T) {
	cases := []struct {
		name  string
		left  *ListInt
		right *ListInt
		want  *ListInt
	}{
		{"Length0+Length0",
			newListInt0(), newListInt0(),
			newListInt0()},
		{"Length0+Length1",
			newListInt0(), newListInt1(100),
			newListInt1(100)},
		{"Length0+Length2",
			newListInt0(), newListInt2(100, 200),
			newListInt2(100, 200)},
		{"Length1+Length0",
			newListInt1(10), newListInt0(),
			newListInt1(10)},
		{"Length1+Length1",
			newListInt1(10), newListInt1(100),
			newListInt2(10, 100)},
		{"Length1+Length2",
			newListInt1(10), newListInt2(100, 200),
			newListInt3(10, 100, 200)},
		{"Length2+Length0",
			newListInt2(10, 20), newListInt0(),
			newListInt2(10, 20)},
		{"Length2+Length1",
			newListInt2(10, 20), newListInt1(100),
			newListInt3(10, 20, 100)},
		{"Length2+Length2",
			newListInt2(10, 20), newListInt2(100, 200),
			newListInt4(10, 20, 100, 200)},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.left.Concat(c.right)
			assertEqual(t, c.want, c.left)
			assertEqual(t, c.want, c.right)
		})
	}
}

func TestListIntEach(t *testing.T) {
	cases := []struct {
		name string
		list *ListInt
		want []int
	}{
		{"Length0", newListInt0(), []int{}},
		{"Length1", newListInt1(10), []int{10}},
		{"Length2", newListInt2(10, 20), []int{10, 20}},
		{"Length3", newListInt3(10, 20, 30), []int{10, 20, 30}},
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

func TestListIntToA(t *testing.T) {
	cases := []struct {
		name string
		list *ListInt
		want []int
	}{
		{"Length0", newListInt0(), []int{}},
		{"Length1", newListInt1(10), []int{10}},
		{"Length2", newListInt2(10, 20), []int{10, 20}},
		{"Length3", newListInt3(10, 20, 30), []int{10, 20, 30}},
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

func newListInt0() *ListInt {
	return &ListInt{first: nil, last: nil, len: 0}
}

func newListInt1(value1 int) *ListInt {
	node1 := listIntNode{parent: nil, child: nil, value: value1}
	return &ListInt{first: &node1, last: &node1, len: 1}
}

func newListInt2(value1, value2 int) *ListInt {
	node1 := listIntNode{parent: nil, child: nil, value: value1}
	node2 := listIntNode{parent: &node1, child: nil, value: value2}
	node1.child = &node2
	return &ListInt{first: &node1, last: &node2, len: 2}
}

func newListInt3(value1, value2, value3 int) *ListInt {
	node1 := listIntNode{parent: nil, child: nil, value: value1}
	node2 := listIntNode{parent: &node1, child: nil, value: value2}
	node1.child = &node2
	node3 := listIntNode{parent: &node2, child: nil, value: value3}
	node2.child = &node3
	return &ListInt{first: &node1, last: &node3, len: 3}
}

func newListInt4(value1, value2, value3, value4 int) *ListInt {
	node1 := listIntNode{parent: nil, child: nil, value: value1}
	node2 := listIntNode{parent: &node1, child: nil, value: value2}
	node1.child = &node2
	node3 := listIntNode{parent: &node2, child: nil, value: value3}
	node2.child = &node3
	node4 := listIntNode{parent: &node3, child: nil, value: value4}
	node3.child = &node4
	return &ListInt{first: &node1, last: &node4, len: 4}
}
