package avltree

import "fmt"

type AVLMapIntInt struct {
	root *avlMapIntIntNode
	size int
}

type avlMapIntIntNode struct {
	key    int
	value  int
	left   *avlMapIntIntNode
	right  *avlMapIntIntNode
	height int
}

func NewAVLMapIntInt() *AVLMapIntInt {
	return &AVLMapIntInt{
		root: nil,
		size: 0,
	}
}

func (m *AVLMapIntInt) Size() int {
	return m.size
}

func (m *AVLMapIntInt) Empty() bool {
	return m.root == nil
}

func (m *AVLMapIntInt) Put(key int, value int) {
	m.root = m.put(m.root, key, value)
}

func (m *AVLMapIntInt) Delete(key int) {
	m.root = m.delete(m.root, key)
}

func (m *AVLMapIntInt) Clear() {
	m.root = nil
	m.size = 0
}

func (m *AVLMapIntInt) Get(key int) (value int, found bool) {
	node := m.root
	for node != nil {
		cmp := node.compare(key)
		switch {
		case cmp == 0:
			value = node.value
			found = true
			return

		case cmp < 0:
			node = node.left
		case cmp > 0:
			node = node.right
		}
	}
	return
}

func (m *AVLMapIntInt) MostLeft() (int, int) {
	if m.root == nil {
		panic("no item")
	}

	node := m.root
	for node.left != nil {
		node = node.left
	}
	return node.key, node.value
}

func (m *AVLMapIntInt) MostRight() (int, int) {
	if m.root == nil {
		panic("no item")
	}

	node := m.root
	for node.right != nil {
		node = node.right
	}
	return node.key, node.value
}

func (m *AVLMapIntInt) Floor(key int) (k int, v int, found bool) {
	node := m.root
	for node != nil {
		cmp := node.compare(key)
		switch {
		case cmp == 0:
			k, v = node.key, node.value
			found = true
			return

		case cmp < 0:
			node = node.left
		case cmp > 0:
			k, v = node.key, node.value
			found = true
			node = node.right
		}
	}
	return
}

func (m *AVLMapIntInt) Ceiling(key int) (k int, v int, found bool) {
	node := m.root
	for node != nil {
		cmp := node.compare(key)
		switch {
		case cmp == 0:
			k, v = node.key, node.value
			found = true
			return

		case cmp > 0:
			node = node.right
		case cmp < 0:
			k, v = node.key, node.value
			found = true
			node = node.left
		}
	}
	return
}

func (m *AVLMapIntInt) Iterator() *AVLMapIntIntIterator {
	iter := AVLMapIntIntIterator{}
	iter.iterate(m)
	return &iter
}

func (m *AVLMapIntInt) IteratorReverse() *AVLMapIntIntIterator {
	iter := AVLMapIntIntIterator{}
	iter.iterateReverse(m)
	return &iter
}

func (m *AVLMapIntInt) IteratorFrom(key int) *AVLMapIntIntIterator {
	iter := AVLMapIntIntIterator{}
	iter.iterateFrom(m, key)
	return &iter
}

func (m *AVLMapIntInt) IteratorReverseFrom(key int) *AVLMapIntIntIterator {
	iter := AVLMapIntIntIterator{}
	iter.iterateReverseFrom(m, key)
	return &iter
}

func (m *AVLMapIntInt) String() string {
	str := "AVLTree\n"
	if !m.Empty() {
		m.string(m.root, "", true, &str)
	}
	return str
}

func (m *AVLMapIntInt) put(node *avlMapIntIntNode, key int, value int) *avlMapIntIntNode {
	if node == nil {
		m.size++
		return &avlMapIntIntNode{
			key:    key,
			value:  value,
			height: 1,
		}
	}

	cmp := node.compare(key)
	switch {
	case cmp == 0:
		node.value = value
		return node

	case cmp < 0:
		node.left = m.put(node.left, key, value)
	case cmp > 0:
		node.right = m.put(node.right, key, value)
	}
	node.adjustHeight()
	return m.balance(node)
}

func (m *AVLMapIntInt) delete(node *avlMapIntIntNode, key int) *avlMapIntIntNode {
	if node == nil {
		return nil
	}

	cmp := node.compare(key)
	switch {
	case cmp == 0:
		m.size--

		if node.left == nil {
			return node.right
		}

		var max *avlMapIntIntNode
		node.left, max = m.deleteMax(node.left)
		node.key, node.value = max.key, max.value
	case cmp < 0:
		node.left = m.delete(node.left, key)
	case cmp > 0:
		node.right = m.delete(node.right, key)
	}
	node.adjustHeight()
	return m.balance(node)
}

func (m *AVLMapIntInt) deleteMax(node *avlMapIntIntNode) (*avlMapIntIntNode, *avlMapIntIntNode) {
	if node == nil {
		panic("node should not be nil")
	}

	if node.right == nil {
		return node.left, node
	}

	var max *avlMapIntIntNode
	node.right, max = m.deleteMax(node.right)
	node.adjustHeight()
	return m.balance(node), max
}

func (m *AVLMapIntInt) balance(node *avlMapIntIntNode) *avlMapIntIntNode {
	if node == nil {
		return nil
	}

	if node.bias() < -1 {
		if node.right.bias() > 0 {
			node.right = m.rotateR(node.right)
		}
		return m.rotateL(node)
	}
	if node.bias() > 1 {
		if node.left.bias() < 0 {
			node.left = m.rotateL(node.left)
		}
		return m.rotateR(node)
	}
	return node
}

func (m *AVLMapIntInt) rotateL(node *avlMapIntIntNode) *avlMapIntIntNode {
	if node == nil {
		panic("node should not be nil")
	}
	if node.right == nil {
		panic("the right of the node should not be nil")
	}

	oldRoot, newRoot := node, node.right
	oldRoot.right, newRoot.left = newRoot.left, oldRoot
	oldRoot.adjustHeight()
	newRoot.adjustHeight()
	return newRoot
}

func (m *AVLMapIntInt) rotateR(node *avlMapIntIntNode) *avlMapIntIntNode {
	if node == nil {
		panic("node should not be nil")
	}
	if node.left == nil {
		panic("the left of the node should not be nil")
	}

	oldRoot, newRoot := node, node.left
	oldRoot.left, newRoot.right = newRoot.right, oldRoot
	oldRoot.adjustHeight()
	newRoot.adjustHeight()
	return newRoot
}

func (m *AVLMapIntInt) string(node *avlMapIntIntNode, prefix string, isTail bool, str *string) {
	if node.right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		m.string(node.right, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if node.left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		m.string(node.left, newPrefix, true, str)
	}
}

func (n *avlMapIntIntNode) compare(key int) int {
	return key - n.key
}

func (n *avlMapIntIntNode) leftHeight() int {
	if n.left == nil {
		return 0
	}
	return n.left.height
}

func (n *avlMapIntIntNode) rightHeight() int {
	if n.right == nil {
		return 0
	}
	return n.right.height
}

func (n *avlMapIntIntNode) adjustHeight() {
	l := n.leftHeight()
	r := n.rightHeight()
	if l < r {
		n.height = 1 + r
	} else {
		n.height = 1 + l
	}
}

func (n *avlMapIntIntNode) bias() int {
	return n.leftHeight() - n.rightHeight()
}

func (n *avlMapIntIntNode) String() string {
	return fmt.Sprintf("[%v] %v", n.key, n.value)
}

type AVLMapIntIntIterator struct {
	cur *avlMapIntIntNode
	c   chan *avlMapIntIntNode
}

func (iter *AVLMapIntIntIterator) Next() bool {
	if iter.c == nil {
		panic("iteration has not begun")
	}

	var ok bool
	iter.cur, ok = <-iter.c
	return ok
}

func (iter *AVLMapIntIntIterator) Key() int {
	if iter.cur == nil {
		panic("no item")
	}

	return iter.cur.key
}

func (iter *AVLMapIntIntIterator) Value() int {
	if iter.cur == nil {
		panic("no item")
	}

	return iter.cur.value
}

func (iter *AVLMapIntIntIterator) iterate(m *AVLMapIntInt) {
	iter.c = make(chan *avlMapIntIntNode)
	go func() {
		defer close(iter.c)
		iter.iterateNode(m.root)
	}()
}

func (iter *AVLMapIntIntIterator) iterateReverse(m *AVLMapIntInt) {
	iter.c = make(chan *avlMapIntIntNode)
	go func() {
		defer close(iter.c)
		iter.iterateNodeReverse(m.root)
	}()
}

func (iter *AVLMapIntIntIterator) iterateFrom(m *AVLMapIntInt, key int) {
	iter.c = make(chan *avlMapIntIntNode)
	go func() {
		defer close(iter.c)
		iter.iterateNodeFrom(m.root, key)
	}()
}

func (iter *AVLMapIntIntIterator) iterateReverseFrom(m *AVLMapIntInt, key int) {
	iter.c = make(chan *avlMapIntIntNode)
	go func() {
		defer close(iter.c)
		iter.iterateNodeReverseFrom(m.root, key)
	}()
}

func (iter *AVLMapIntIntIterator) iterateNode(node *avlMapIntIntNode) {
	if node == nil {
		return
	}

	iter.iterateNode(node.left)
	iter.c <- node
	iter.iterateNode(node.right)
}

func (iter *AVLMapIntIntIterator) iterateNodeReverse(node *avlMapIntIntNode) {
	if node == nil {
		return
	}

	iter.iterateNodeReverse(node.right)
	iter.c <- node
	iter.iterateNodeReverse(node.left)
}

func (iter *AVLMapIntIntIterator) iterateNodeFrom(node *avlMapIntIntNode, key int) {
	if node == nil {
		return
	}

	cmp := node.compare(key)
	if cmp <= 0 {
		if cmp < 0 {
			iter.iterateNodeFrom(node.left, key)
		}
		iter.c <- node
		iter.iterateNode(node.right)
	} else {
		iter.iterateNodeFrom(node.right, key)
	}
}

func (iter *AVLMapIntIntIterator) iterateNodeReverseFrom(node *avlMapIntIntNode, key int) {
	if node == nil {
		return
	}

	cmp := node.compare(key)
	if cmp >= 0 {
		if cmp > 0 {
			iter.iterateNodeReverseFrom(node.right, key)
		}
		iter.c <- node
		iter.iterateNodeReverse(node.left)
	} else {
		iter.iterateNodeReverseFrom(node.left, key)
	}
}
