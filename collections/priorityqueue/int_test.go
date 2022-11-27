package priorityqueue

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestInitAndPop(t *testing.T) {
	pq := make(IntPriorityQueue, 5)
	pq[0] = &IntPriorityQueueItem{Value: 0, Priority: 20, Index: 0}
	pq[1] = &IntPriorityQueueItem{Value: 10, Priority: 40, Index: 1}
	pq[2] = &IntPriorityQueueItem{Value: 20, Priority: 10, Index: 2}
	pq[3] = &IntPriorityQueueItem{Value: 30, Priority: 50, Index: 3}
	pq[4] = &IntPriorityQueueItem{Value: 40, Priority: 30, Index: 4}

	heap.Init(&pq)

	results := make([]int, 5)
	for i := 0; i < 5; i++ {
		results[i] = heap.Pop(&pq).(*IntPriorityQueueItem).Value
	}

	assertEqual(t, []int{30, 10, 40, 0, 20}, results)
}

func TestPushAndPop(t *testing.T) {
	pq := make(IntPriorityQueue, 3)
	pq[0] = &IntPriorityQueueItem{Value: 0, Priority: 20, Index: 0}
	pq[1] = &IntPriorityQueueItem{Value: 10, Priority: 40, Index: 1}
	pq[2] = &IntPriorityQueueItem{Value: 20, Priority: 10, Index: 2}

	heap.Init(&pq)

	heap.Push(&pq, &IntPriorityQueueItem{Value: 30, Priority: 50, Index: 3})
	heap.Push(&pq, &IntPriorityQueueItem{Value: 40, Priority: 30, Index: 4})

	results := make([]int, 5)
	for i := 0; i < 5; i++ {
		results[i] = heap.Pop(&pq).(*IntPriorityQueueItem).Value
	}

	assertEqual(t, []int{30, 10, 40, 0, 20}, results)
}

func TestUpdate(t *testing.T) {
	items := make([]*IntPriorityQueueItem, 5)
	items[0] = &IntPriorityQueueItem{Value: 0, Priority: 20, Index: 0}
	items[1] = &IntPriorityQueueItem{Value: 10, Priority: 40, Index: 1}
	items[2] = &IntPriorityQueueItem{Value: 20, Priority: 10, Index: 2}
	items[3] = &IntPriorityQueueItem{Value: 30, Priority: 50, Index: 3}
	items[4] = &IntPriorityQueueItem{Value: 40, Priority: 30, Index: 4}

	pq := make(IntPriorityQueue, 5)
	for i := 0; i < 5; i++ {
		pq[i] = items[i]
	}

	heap.Init(&pq)

	pq.Update(items[2], 50, 100)

	results := make([]int, 5)
	for i := 0; i < 5; i++ {
		results[i] = heap.Pop(&pq).(*IntPriorityQueueItem).Value
	}

	assertEqual(t, []int{50, 30, 10, 40, 0}, results)
}

func assertEqual(t *testing.T, want, got interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want: %#v, got: %#v\n", want, got)
	}
}
