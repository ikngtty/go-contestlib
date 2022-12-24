package list

// ListInt is a list of int.
type ListInt struct {
	first *listIntNode
	last  *listIntNode
	len   int
}

type listIntNode struct {
	parent *listIntNode
	child  *listIntNode
	value  int
}

// NewListInt returns a new ListInt.
func NewListInt() *ListInt {
	return &ListInt{nil, nil, 0}
}

// NewListIntFromArray returns a new ListInt having the values `a` has.
func NewListIntFromArray(a []int) *ListInt {
	list := NewListInt()
	for _, elem := range a {
		list.Push(elem)
	}
	return list
}

// Len returns the length of the list.
func (list *ListInt) Len() int {
	return list.len
}

// Push pushes elem to the end of the list.
func (list *ListInt) Push(elem int) {
	node := listIntNode{list.last, nil, elem}
	if list.first == nil {
		list.first = &node
	} else {
		list.last.child = &node
	}
	list.last = &node
	list.len++
}

// PushLeft pushes elem to the beginning of the list.
func (list *ListInt) PushLeft(elem int) {
	node := listIntNode{nil, list.first, elem}
	if list.last == nil {
		list.last = &node
	} else {
		list.first.parent = &node
	}
	list.first = &node
	list.len++
}

// Pop pops elem from the end of the list.
func (list *ListInt) Pop() int {
	if list.last == nil {
		panic("no item")
	}
	value := list.last.value
	list.last = list.last.parent
	if list.last == nil {
		list.first = nil
	} else {
		list.last.child = nil
	}
	list.len--
	return value
}

// PopLeft pops elem from the beginning of the list.
func (list *ListInt) PopLeft() int {
	if list.first == nil {
		panic("no item")
	}
	value := list.first.value
	list.first = list.first.child
	if list.first == nil {
		list.last = nil
	} else {
		list.first.parent = nil
	}
	list.len--
	return value
}

// Concat concatenates the list and the other.
func (list *ListInt) Concat(other *ListInt) {
	if list.first == nil {
		*list = *other
	} else if other.first == nil {
		*other = *list
	} else {
		list.last.child = other.first
		other.first.parent = list.last
		list.last = other.last
		other.first = list.first
		list.len += other.len
		other.len = list.len
	}
}

// Each applies f for every element in the list.
func (list *ListInt) Each(f func(elem int)) {
	cur := list.first
	for cur != nil {
		f(cur.value)
		cur = cur.child
	}
}

// ToA converts the list to an array.
func (list *ListInt) ToA() []int {
	a := make([]int, list.len)
	{
		index := 0
		list.Each(func(elem int) {
			a[index] = elem
			index++
		})
	}
	return a
}
