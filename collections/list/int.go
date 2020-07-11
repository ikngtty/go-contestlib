package list

type IntList struct {
	root *intListNode
	last *intListNode
	len  int
}

type intListNode struct {
	child *intListNode
	value int
}

func NewIntList() *IntList {
	root := intListNode{nil, 0}
	return &IntList{&root, &root, 0}
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
	list.last.child = &node
	list.last = &node
	list.len++
}

func (list *IntList) Concat(other *IntList) {
	if other.Len() == 0 {
		return
	}
	list.last.child = other.root.child
	list.last = other.last
	list.len += other.Len()
}

func (list *IntList) Each(f func(elem int)) {
	cur := list.root
	for cur.child != nil {
		cur = cur.child
		f(cur.value)
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
