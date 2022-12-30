package avltree

type AVLMultiSetInt struct {
	m    *AVLMapIntInt
	size int
}

func NewAVLMultiSetInt() *AVLMultiSetInt {
	return &AVLMultiSetInt{
		m:    NewAVLMapIntInt(),
		size: 0,
	}
}

func (ms *AVLMultiSetInt) Size() int {
	return ms.size
}

func (ms *AVLMultiSetInt) Empty() bool {
	return ms.m.Empty()
}

func (ms *AVLMultiSetInt) Add(item int) {
	count, _ := ms.m.Get(item)
	ms.m.Put(item, count+1)
	ms.size++
}

func (ms *AVLMultiSetInt) RemoveOne(item int) bool {
	count, ok := ms.m.Get(item)
	if !ok {
		return false
	}

	if count == 1 {
		ms.m.Delete(item)
	} else {
		ms.m.Put(item, count-1)
	}
	ms.size--
	return true
}

func (ms *AVLMultiSetInt) RemoveAll(item int) int {
	count, ok := ms.m.Get(item)
	if !ok {
		return 0
	}

	ms.m.Delete(item)
	ms.size -= count
	return count
}

func (ms *AVLMultiSetInt) Clear() {
	ms.m.Clear()
	ms.size = 0
}

func (ms *AVLMultiSetInt) Contains(item int) (found bool) {
	_, found = ms.m.Get(item)
	return
}

func (ms *AVLMultiSetInt) Count(item int) (count int) {
	count, _ = ms.m.Get(item)
	return
}

func (ms *AVLMultiSetInt) MostLeft() (item int) {
	item, _ = ms.m.MostLeft()
	return
}

func (ms *AVLMultiSetInt) MostRight() (item int) {
	item, _ = ms.m.MostRight()
	return
}

func (ms *AVLMultiSetInt) Floor(key int) (item int, found bool) {
	item, _, found = ms.m.Floor(key)
	return
}

func (ms *AVLMultiSetInt) Ceiling(key int) (item int, found bool) {
	item, _, found = ms.m.Ceiling(key)
	return
}

func (ms *AVLMultiSetInt) Iterator() *AVLMultiSetIntIterator {
	iter := AVLMultiSetIntIterator{}
	iter.iterate(ms.m.Iterator())
	return &iter
}

func (ms *AVLMultiSetInt) IteratorReverse() *AVLMultiSetIntIterator {
	iter := AVLMultiSetIntIterator{}
	iter.iterate(ms.m.IteratorReverse())
	return &iter
}

func (ms *AVLMultiSetInt) IteratorFrom(key int) *AVLMultiSetIntIterator {
	iter := AVLMultiSetIntIterator{}
	iter.iterate(ms.m.IteratorFrom(key))
	return &iter
}

func (ms *AVLMultiSetInt) IteratorReverseFrom(key int) *AVLMultiSetIntIterator {
	iter := AVLMultiSetIntIterator{}
	iter.iterate(ms.m.IteratorReverseFrom(key))
	return &iter
}

type AVLMultiSetIntIterator struct {
	cur *int
	c   chan *int
}

func (iter *AVLMultiSetIntIterator) Next() bool {
	if iter.c == nil {
		panic("iteration has not begun")
	}

	var ok bool
	iter.cur, ok = <-iter.c
	return ok
}

func (iter *AVLMultiSetIntIterator) Value() int {
	if iter.cur == nil {
		panic("no item")
	}

	return *iter.cur
}

func (msIter *AVLMultiSetIntIterator) iterate(mIter *AVLMapIntIntIterator) {
	msIter.c = make(chan *int)
	go func() {
		defer close(msIter.c)
		for mIter.Next() {
			item, count := mIter.Key(), mIter.Value()
			for i := 0; i < count; i++ {
				msIter.c <- &item
			}
		}
	}()
}
