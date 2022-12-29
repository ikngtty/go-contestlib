package avltree

import (
	"math/rand"
	"reflect"
	"testing"
)

func theAVLMapIntInt(size int) *AVLMapIntInt {
	a := make([]int, size)
	for i := 0; i < len(a); i++ {
		a[i] = 2 * i
	}
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})

	m := NewAVLMapIntInt()
	for _, k := range a {
		v := k * 10
		m.Put(k, v)
	}
	return m
}

func testBinarySearchTree(t *testing.T, node *avlMapIntIntNode) {
	if node == nil {
		return
	}

	testBinarySearchTree(t, node.left)
	testBinarySearchTree(t, node.right)

	if node.left != nil && node.compare(node.left.key) > 0 {
		t.Errorf("node key: %d, left key: %d", node.key, node.left.key)
		t.FailNow()
	}
	if node.right != nil && node.compare(node.right.key) < 0 {
		t.Errorf("node key: %d, right key: %d", node.key, node.right.key)
		t.FailNow()
	}
}

func testHeightAdjusted(t *testing.T, node *avlMapIntIntNode) {
	if node == nil {
		return
	}

	testHeightAdjusted(t, node.left)
	testHeightAdjusted(t, node.right)

	height := node.height
	node.adjustHeight()
	if node.height != height {
		t.Errorf("height of node %d is %d, not %d", node.key, node.height, height)
		t.FailNow()
	}
}

func testBalanced(t *testing.T, node *avlMapIntIntNode) {
	if node == nil {
		return
	}

	testBalanced(t, node.left)
	testBalanced(t, node.right)

	if node.bias() < -1 || 1 < node.bias() {
		t.Errorf("bias of node %d is %d", node.key, node.bias())
		t.FailNow()
	}
}

func TestPutSmall(t *testing.T) {
	const size = 20
	m := theAVLMapIntInt(size)
	t.Log(m)

	testBinarySearchTree(t, m.root)
	testHeightAdjusted(t, m.root)
	testBalanced(t, m.root)
	if m.Size() != size {
		t.Errorf("size want: %d, got: %d", size, m.Size())
	}
}

func TestPutLarge(t *testing.T) {
	const size = 100000
	m := theAVLMapIntInt(size)

	testBinarySearchTree(t, m.root)
	testHeightAdjusted(t, m.root)
	testBalanced(t, m.root)
	if m.Size() != size {
		t.Errorf("size want: %d, got: %d", size, m.Size())
	}
}

func TestPutSameKey(t *testing.T) {
	const size = 20
	m := theAVLMapIntInt(size)
	const key = 8
	const value = 42
	m.Put(key, value)
	t.Log(m)

	if m.Size() != size {
		t.Errorf("size want: %d, got: %d", size, m.Size())
	}
	v, _ := m.Get(key)
	if v != value {
		t.Errorf("replaced value want: %d, got: %d", value, v)
	}
}

func TestDeleteSmall(t *testing.T) {
	m := theAVLMapIntInt(10)
	deleteAndTest := func(key int) {
		m.Delete(key)
		value, ok := m.Get(key)
		if ok {
			t.Errorf("key %d is not deleted and got value %d", key, value)
		}
	}
	deleteAndTest(4)
	deleteAndTest(6)
	deleteAndTest(8)
	m.Delete(100)
	t.Log(m)

	testBinarySearchTree(t, m.root)
	testHeightAdjusted(t, m.root)
	testBalanced(t, m.root)
	if m.Size() != 7 {
		t.Errorf("size want: %d, got: %d", 7, m.Size())
	}
}

func TestDeleteLarge(t *testing.T) {
	const size = 100000
	m := theAVLMapIntInt(size)
	const deleteSize = 10000
	for i := 0; i < deleteSize; i++ {
		m.Delete(4 * i)
	}
	const wantedSize = size - deleteSize

	testBinarySearchTree(t, m.root)
	testHeightAdjusted(t, m.root)
	testBalanced(t, m.root)
	if m.Size() != wantedSize {
		t.Errorf("size want: %d, got: %d", wantedSize, m.Size())
	}
}

func TestGet(t *testing.T) {
	m := theAVLMapIntInt(20)

	t.Run("exist key", func(t *testing.T) {
		const key = 4
		const wantedValue = key * 10

		value, ok := m.Get(key)
		if !ok {
			t.Errorf("get key %d is failed", key)
		}
		if value != wantedValue {
			t.Errorf("get key %d want: %d, got %d", key, wantedValue, value)
		}
	})

	t.Run("no exist key", func(t *testing.T) {
		const key = 5

		value, ok := m.Get(key)
		if ok {
			t.Errorf("get key %d should be failed, but got %d", key, value)
		}
	})
}

func TestMostLeft(t *testing.T) {
	m := theAVLMapIntInt(20)
	m.Delete(0)

	const wantedKey = 2
	const wantedValue = 20
	key, value := m.MostLeft()
	if key != wantedKey {
		t.Errorf("key want: %d, got: %d", wantedKey, key)
	}
	if value != wantedValue {
		t.Errorf("value want: %d, got: %d", wantedValue, value)
	}
}

func TestMostRight(t *testing.T) {
	m := theAVLMapIntInt(20)

	const wantedKey = 38
	const wantedValue = 380
	key, value := m.MostRight()
	if key != wantedKey {
		t.Errorf("key want: %d, got: %d", wantedKey, key)
	}
	if value != wantedValue {
		t.Errorf("value want: %d, got: %d", wantedValue, value)
	}
}

func TestFloor(t *testing.T) {
	m := theAVLMapIntInt(20)

	t.Run("exist key", func(t *testing.T) {
		const key = 4
		const wantedK = 4
		const wantedV = wantedK * 10
		k, v, ok := m.Floor(key)
		if !ok {
			t.Errorf("get key %d is failed", key)
		}
		if k != wantedK {
			t.Errorf("key want: %d, got: %d", wantedK, k)
		}
		if v != wantedV {
			t.Errorf("value want: %d, got: %d", wantedV, v)
		}
	})

	t.Run("no exist key", func(t *testing.T) {
		const key = 7
		const wantedK = 6
		const wantedV = wantedK * 10
		k, v, ok := m.Floor(key)
		if !ok {
			t.Errorf("get key %d is failed", key)
		}
		if k != wantedK {
			t.Errorf("key want: %d, got: %d", wantedK, k)
		}
		if v != wantedV {
			t.Errorf("value want: %d, got: %d", wantedV, v)
		}
	})

	t.Run("no floor", func(t *testing.T) {
		const key = -1
		k, v, ok := m.Floor(key)
		if ok {
			t.Errorf("get key %d should be failed, but got key %d and value %d", key, k, v)
		}
	})
}

func TestCeiling(t *testing.T) {
	m := theAVLMapIntInt(20)

	t.Run("exist key", func(t *testing.T) {
		const key = 12
		const wantedK = 12
		const wantedV = wantedK * 10
		k, v, ok := m.Ceiling(key)
		if !ok {
			t.Errorf("get key %d is failed", key)
		}
		if k != wantedK {
			t.Errorf("key want: %d, got: %d", wantedK, k)
		}
		if v != wantedV {
			t.Errorf("value want: %d, got: %d", wantedV, v)
		}
	})

	t.Run("no exist key", func(t *testing.T) {
		const key = 15
		const wantedK = 16
		const wantedV = wantedK * 10
		k, v, ok := m.Ceiling(key)
		if !ok {
			t.Errorf("get key %d is failed", key)
		}
		if k != wantedK {
			t.Errorf("key want: %d, got: %d", wantedK, k)
		}
		if v != wantedV {
			t.Errorf("value want: %d, got: %d", wantedV, v)
		}
	})

	t.Run("no ceiling", func(t *testing.T) {
		const key = 40
		k, v, ok := m.Ceiling(key)
		if ok {
			t.Errorf("get key %d should be failed, but got key %d and value %d", key, k, v)
		}
	})
}

func TestIterator(t *testing.T) {
	const size = 10
	m := theAVLMapIntInt(size)

	wantedIteratedKeys := make([]int, size)
	wantedIteratedValues := make([]int, size)
	for i := 0; i < size; i++ {
		wantedIteratedKeys[i] = 2 * i
		wantedIteratedValues[i] = 20 * i
	}

	iter := m.Iterator()
	iteratedKeys := make([]int, 0)
	iteratedValues := make([]int, 0)
	for i := 0; i < size && iter.Next(); i++ {
		iteratedKeys = append(iteratedKeys, iter.Key())
		iteratedValues = append(iteratedValues, iter.Value())
	}

	if iter.Next() {
		t.Errorf("iterator should stop but does not")
	}
	if !reflect.DeepEqual(iteratedKeys, wantedIteratedKeys) {
		t.Errorf("iterated keys want: %v, got: %v", wantedIteratedKeys, iteratedKeys)
	}
	if !reflect.DeepEqual(iteratedValues, wantedIteratedValues) {
		t.Errorf("iterated values want: %v, got: %v", wantedIteratedValues, iteratedValues)
	}
}

func TestIteratorReverse(t *testing.T) {
	const size = 10
	m := theAVLMapIntInt(size)

	wantedIteratedKeys := make([]int, size)
	wantedIteratedValues := make([]int, size)
	for i := 0; i < size; i++ {
		wantedIteratedKeys[i] = 2 * (size - 1 - i)
		wantedIteratedValues[i] = 20 * (size - 1 - i)
	}

	iter := m.IteratorReverse()
	iteratedKeys := make([]int, 0)
	iteratedValues := make([]int, 0)
	for i := 0; i < size && iter.Next(); i++ {
		iteratedKeys = append(iteratedKeys, iter.Key())
		iteratedValues = append(iteratedValues, iter.Value())
	}

	if iter.Next() {
		t.Errorf("iterator should stop but does not")
	}
	if !reflect.DeepEqual(iteratedKeys, wantedIteratedKeys) {
		t.Errorf("iterated keys want: %v, got: %v", wantedIteratedKeys, iteratedKeys)
	}
	if !reflect.DeepEqual(iteratedValues, wantedIteratedValues) {
		t.Errorf("iterated values want: %v, got: %v", wantedIteratedValues, iteratedValues)
	}
}

func TestIteratorFrom(t *testing.T) {
	const size = 20
	const skipCount = 5
	m := theAVLMapIntInt(size)

	wantedIteratedKeys := make([]int, size-skipCount)
	wantedIteratedValues := make([]int, size-skipCount)
	for i := 0; i < size-skipCount; i++ {
		wantedIteratedKeys[i] = 2 * (i + skipCount)
		wantedIteratedValues[i] = 20 * (i + skipCount)
	}

	iter := m.IteratorFrom(2 * skipCount)
	iteratedKeys := make([]int, 0)
	iteratedValues := make([]int, 0)
	for i := 0; i < size-skipCount && iter.Next(); i++ {
		iteratedKeys = append(iteratedKeys, iter.Key())
		iteratedValues = append(iteratedValues, iter.Value())
	}

	if iter.Next() {
		t.Errorf("iterator should stop but does not")
	}
	if !reflect.DeepEqual(iteratedKeys, wantedIteratedKeys) {
		t.Errorf("iterated keys want: %v, got: %v", wantedIteratedKeys, iteratedKeys)
	}
	if !reflect.DeepEqual(iteratedValues, wantedIteratedValues) {
		t.Errorf("iterated values want: %v, got: %v", wantedIteratedValues, iteratedValues)
	}
}

func TestIteratorReverseFrom(t *testing.T) {
	const size = 20
	const skipCount = 5
	m := theAVLMapIntInt(size)

	wantedIteratedKeys := make([]int, size-skipCount)
	wantedIteratedValues := make([]int, size-skipCount)
	for i := 0; i < size-skipCount; i++ {
		wantedIteratedKeys[i] = 2 * (size - 1 - i - skipCount)
		wantedIteratedValues[i] = 20 * (size - 1 - i - skipCount)
	}

	iter := m.IteratorReverseFrom(2 * (size - 1 - skipCount))
	iteratedKeys := make([]int, 0)
	iteratedValues := make([]int, 0)
	for i := 0; i < size-skipCount && iter.Next(); i++ {
		iteratedKeys = append(iteratedKeys, iter.Key())
		iteratedValues = append(iteratedValues, iter.Value())
	}

	if iter.Next() {
		t.Errorf("iterator should stop but does not")
	}
	if !reflect.DeepEqual(iteratedKeys, wantedIteratedKeys) {
		t.Errorf("iterated keys want: %v, got: %v", wantedIteratedKeys, iteratedKeys)
	}
	if !reflect.DeepEqual(iteratedValues, wantedIteratedValues) {
		t.Errorf("iterated values want: %v, got: %v", wantedIteratedValues, iteratedValues)
	}
}

func TestString(t *testing.T) {
	m := NewAVLMapIntInt()
	for i := 0; i < 6; i++ {
		m.Put(i, i*10)
	}

	str := m.String()
	const wantedStr = `AVLTree
│       ┌── [5] 50
│   ┌── [4] 40
└── [3] 30
    │   ┌── [2] 20
    └── [1] 10
        └── [0] 0
`

	if str != wantedStr {
		t.Errorf("want:\n%s\ngot:\n%s", wantedStr, str)
	}
}
