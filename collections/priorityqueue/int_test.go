package priorityqueue

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestInitAndPop(t *testing.T) {
	pq := HeapInt{20, 40, 10, 50, 30}

	heap.Init(&pq)

	results := make([]int, 5)
	for i := 0; i < 5; i++ {
		results[i] = heap.Pop(&pq).(int)
	}

	assertEqual(t, []int{10, 20, 30, 40, 50}, results)
}

func TestPushAndPop(t *testing.T) {
	pq := HeapInt{20, 40, 10}

	heap.Init(&pq)

	heap.Push(&pq, 50)
	heap.Push(&pq, 30)

	results := make([]int, 5)
	for i := 0; i < 5; i++ {
		results[i] = heap.Pop(&pq).(int)
	}

	assertEqual(t, []int{10, 20, 30, 40, 50}, results)
}

func assertEqual(t *testing.T, want, got interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want: %#v, got: %#v\n", want, got)
	}
}
