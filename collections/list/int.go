package list

// IntList is a list of int.
type IntList struct {
	first *intListNode
	last  *intListNode
	len   int
}

type intListNode struct {
	child *intListNode
	value int
}

// NewIntList returns a new IntList.
func NewIntList() *IntList {
	return &IntList{nil, nil, 0}
}

// NewIntListFromArray returns a new IntList having the values "a" has.
func NewIntListFromArray(a []int) *IntList {
	list := NewIntList()
	for _, elem := range a {
		list.Add(elem)
	}
	return list
}

// Len returns length of the list.
func (list *IntList) Len() int {
	return list.len
}

// Add adds elem to the list.
func (list *IntList) Add(elem int) {
	node := intListNode{nil, elem}
	if list.first == nil {
		list.first = &node
		list.last = &node
	} else {
		list.last.child = &node
		list.last = &node
	}
	list.len++
}

// Concat concatenates the list and the other.
func (list *IntList) Concat(other *IntList) {
	if list.first == nil {
		*list = *other
	} else if other.first == nil {
		// Do nothing
	} else {
		list.last.child = other.first
		list.last = other.last
		list.len += other.len
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
	a := make([]int, list.Len())
	{
		index := 0
		list.Each(func(elem int) {
			a[index] = elem
			index++
		})
	}
	return a
}
