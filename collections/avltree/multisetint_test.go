package avltree

import (
	"math/rand"
	"reflect"
	"strconv"
	"testing"
)

func TestMultiSetAdd(t *testing.T) {
	ms := NewAVLMultiSetInt()
	ms.Add(20)
	ms.Add(42)
	ms.Add(42)

	const wantedSize = 3
	if ms.Size() != wantedSize {
		t.Errorf("size want: %d, got: %d", wantedSize, ms.Size())
	}

	cases := []struct {
		item     int
		count    int
		contains bool
	}{
		{10, 0, false},
		{20, 1, true},
		{42, 2, true},
	}

	for _, c := range cases {
		t.Run(strconv.Itoa(c.item), func(t *testing.T) {
			count := ms.Count(c.item)
			if count != c.count {
				t.Errorf("count want: %d, got: %d", c.count, count)
			}

			contains := ms.Contains(c.item)
			if contains != c.contains {
				t.Errorf("contains want: %v, got: %v", c.contains, contains)
			}
		})
	}
}

func TestMultiSetRemoveOne(t *testing.T) {
	ms := NewAVLMultiSetInt()
	ms.Add(15)
	ms.Add(20)
	ms.Add(42)
	ms.Add(42)
	ms.Add(42)

	removeAndTest := func(item int, wantedOK bool) {
		ok := ms.RemoveOne(item)
		if ok != wantedOK {
			t.Errorf("ok want: %v, got: %v", wantedOK, ok)
		}
	}
	removeAndTest(10, false)
	removeAndTest(20, true)
	removeAndTest(42, true)

	const wantedSize = 3
	if ms.Size() != wantedSize {
		t.Errorf("size want: %d, got: %d", wantedSize, ms.Size())
	}

	cases := []struct {
		item     int
		count    int
		contains bool
	}{
		{20, 0, false},
		{42, 2, true},
	}

	for _, c := range cases {
		t.Run(strconv.Itoa(c.item), func(t *testing.T) {
			count := ms.Count(c.item)
			if count != c.count {
				t.Errorf("count want: %d, got: %d", c.count, count)
			}

			contains := ms.Contains(c.item)
			if contains != c.contains {
				t.Errorf("contains want: %v, got: %v", c.contains, contains)
			}
		})
	}
}

func TestMultiSetRemoveAll(t *testing.T) {
	ms := NewAVLMultiSetInt()
	ms.Add(15)
	ms.Add(20)
	ms.Add(42)
	ms.Add(42)
	ms.Add(42)

	removeAndTest := func(item int, wantedCount int) {
		count := ms.RemoveAll(item)
		if count != wantedCount {
			t.Errorf("count want: %d, got: %d", wantedCount, count)
		}
	}
	removeAndTest(10, 0)
	removeAndTest(42, 3)

	const wantedSize = 2
	if ms.Size() != wantedSize {
		t.Errorf("size want: %d, got: %d", wantedSize, ms.Size())
	}

	cases := []struct {
		item     int
		count    int
		contains bool
	}{
		{42, 0, false},
	}

	for _, c := range cases {
		t.Run(strconv.Itoa(c.item), func(t *testing.T) {
			count := ms.Count(c.item)
			if count != c.count {
				t.Errorf("count want: %d, got: %d", c.count, count)
			}

			contains := ms.Contains(c.item)
			if contains != c.contains {
				t.Errorf("contains want: %v, got: %v", c.contains, contains)
			}
		})
	}
}

func TestMultiSetIterator(t *testing.T) {
	ms := NewAVLMultiSetInt()
	source := []int{10, 10, 10, 20, 20, 30, 40, 40, 50, 50, 50}
	rand.Shuffle(len(source), func(i, j int) {
		source[i], source[j] = source[j], source[i]
	})
	for _, item := range source {
		ms.Add(item)
	}

	wantedIterated := []int{10, 10, 10, 20, 20, 30, 40, 40, 50, 50, 50}

	iter := ms.Iterator()
	iterated := make([]int, 0)
	for i := 0; i < len(wantedIterated) && iter.Next(); i++ {
		iterated = append(iterated, iter.Value())
	}

	if iter.Next() {
		t.Errorf("iterator should stop but does not")
	}
	if !reflect.DeepEqual(iterated, wantedIterated) {
		t.Errorf("iterated keys want: %v, got: %v", wantedIterated, iterated)
	}
}

func TestMultiSetIteratorReverse(t *testing.T) {
	ms := NewAVLMultiSetInt()
	source := []int{10, 10, 10, 20, 20, 30, 40, 40, 50, 50, 50}
	rand.Shuffle(len(source), func(i, j int) {
		source[i], source[j] = source[j], source[i]
	})
	for _, item := range source {
		ms.Add(item)
	}

	wantedIterated := []int{50, 50, 50, 40, 40, 30, 20, 20, 10, 10, 10}

	iter := ms.IteratorReverse()
	iterated := make([]int, 0)
	for i := 0; i < len(wantedIterated) && iter.Next(); i++ {
		iterated = append(iterated, iter.Value())
	}

	if iter.Next() {
		t.Errorf("iterator should stop but does not")
	}
	if !reflect.DeepEqual(iterated, wantedIterated) {
		t.Errorf("iterated keys want: %v, got: %v", wantedIterated, iterated)
	}
}

func TestMultiSetIteratorFrom(t *testing.T) {
	ms := NewAVLMultiSetInt()
	source := []int{10, 10, 10, 20, 20, 30, 40, 40, 50, 50, 50}
	rand.Shuffle(len(source), func(i, j int) {
		source[i], source[j] = source[j], source[i]
	})
	for _, item := range source {
		ms.Add(item)
	}

	wantedIterated := []int{20, 20, 30, 40, 40, 50, 50, 50}

	iter := ms.IteratorFrom(20)
	iterated := make([]int, 0)
	for i := 0; i < len(wantedIterated) && iter.Next(); i++ {
		iterated = append(iterated, iter.Value())
	}

	if iter.Next() {
		t.Errorf("iterator should stop but does not")
	}
	if !reflect.DeepEqual(iterated, wantedIterated) {
		t.Errorf("iterated keys want: %v, got: %v", wantedIterated, iterated)
	}
}

func TestMultiSetIteratorReverseFrom(t *testing.T) {
	ms := NewAVLMultiSetInt()
	source := []int{10, 10, 10, 20, 20, 30, 40, 40, 50, 50, 50}
	rand.Shuffle(len(source), func(i, j int) {
		source[i], source[j] = source[j], source[i]
	})
	for _, item := range source {
		ms.Add(item)
	}

	wantedIterated := []int{40, 40, 30, 20, 20, 10, 10, 10}

	iter := ms.IteratorReverseFrom(40)
	iterated := make([]int, 0)
	for i := 0; i < len(wantedIterated) && iter.Next(); i++ {
		iterated = append(iterated, iter.Value())
	}

	if iter.Next() {
		t.Errorf("iterator should stop but does not")
	}
	if !reflect.DeepEqual(iterated, wantedIterated) {
		t.Errorf("iterated keys want: %v, got: %v", wantedIterated, iterated)
	}
}
