package list

type IntList struct {
	first *intListNode
	last  *intListNode
	len   int
}

type intListNode struct {
	child *intListNode
	value int
}

func NewIntList() *IntList {
	return &IntList{nil, nil, 0}
}

func NewIntListFromArray(a []int) *IntList {
	list := NewIntList()
	for _, elem := range a {
		list.Add(elem)
	}
	return list
}

func (list *IntList) Len() int {
	return list.len
}

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

func (list *IntList) Each(f func(elem int)) {
	cur := list.first
	for cur != nil {
		f(cur.value)
		cur = cur.child
	}
}

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
