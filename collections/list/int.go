package list

// IntList is a list of int.
type IntList struct {
	first *intListNode
	last  *intListNode
	len   int
}

type intListNode struct {
	parent *intListNode
	child  *intListNode
	value  int
}

// NewIntList returns a new IntList.
func NewIntList() *IntList {
	return &IntList{nil, nil, 0}
}

// NewIntListFromArray returns a new IntList having the values `a` has.
func NewIntListFromArray(a []int) *IntList {
	list := NewIntList()
	for _, elem := range a {
		list.Push(elem)
	}
	return list
}

// Len returns the length of the list.
func (list *IntList) Len() int {
	return list.len
}

// Push pushes elem to the end of the list.
func (list *IntList) Push(elem int) {
	node := intListNode{list.last, nil, elem}
	if list.first == nil {
		list.first = &node
	} else {
		list.last.child = &node
	}
	list.last = &node
	list.len++
}

// PushLeft pushes elem to the beginning of the list.
func (list *IntList) PushLeft(elem int) {
	node := intListNode{nil, list.first, elem}
	if list.last == nil {
		list.last = &node
	} else {
		list.first.parent = &node
	}
	list.first = &node
	list.len++
}

// Pop pops elem from the end of the list.
func (list *IntList) Pop() int {
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
func (list *IntList) PopLeft() int {
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
func (list *IntList) Concat(other *IntList) {
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
func (list *IntList) Each(f func(elem int)) {
	cur := list.first
	for cur != nil {
		f(cur.value)
		cur = cur.child
	}
}

// ToA converts the list to an array.
func (list *IntList) ToA() []int {
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
