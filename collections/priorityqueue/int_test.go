package priorityqueue

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestInitAndPop(t *testing.T) {
	pq := make(PriorityQueueInt, 5)
	pq[0] = &PriorityQueueIntItem{Value: 0, Priority: 20, Index: 0}
	pq[1] = &PriorityQueueIntItem{Value: 10, Priority: 40, Index: 1}
	pq[2] = &PriorityQueueIntItem{Value: 20, Priority: 10, Index: 2}
	pq[3] = &PriorityQueueIntItem{Value: 30, Priority: 50, Index: 3}
	pq[4] = &PriorityQueueIntItem{Value: 40, Priority: 30, Index: 4}

	heap.Init(&pq)

	results := make([]int, 5)
	for i := 0; i < 5; i++ {
		results[i] = heap.Pop(&pq).(*PriorityQueueIntItem).Value
	}

	assertEqual(t, []int{30, 10, 40, 0, 20}, results)
}

func TestPushAndPop(t *testing.T) {
	pq := make(PriorityQueueInt, 3)
	pq[0] = &PriorityQueueIntItem{Value: 0, Priority: 20, Index: 0}
	pq[1] = &PriorityQueueIntItem{Value: 10, Priority: 40, Index: 1}
	pq[2] = &PriorityQueueIntItem{Value: 20, Priority: 10, Index: 2}

	heap.Init(&pq)

	heap.Push(&pq, &PriorityQueueIntItem{Value: 30, Priority: 50, Index: 3})
	heap.Push(&pq, &PriorityQueueIntItem{Value: 40, Priority: 30, Index: 4})

	results := make([]int, 5)
	for i := 0; i < 5; i++ {
		results[i] = heap.Pop(&pq).(*PriorityQueueIntItem).Value
	}

	assertEqual(t, []int{30, 10, 40, 0, 20}, results)
}

func TestUpdate(t *testing.T) {
	items := make([]*PriorityQueueIntItem, 5)
	items[0] = &PriorityQueueIntItem{Value: 0, Priority: 20, Index: 0}
	items[1] = &PriorityQueueIntItem{Value: 10, Priority: 40, Index: 1}
	items[2] = &PriorityQueueIntItem{Value: 20, Priority: 10, Index: 2}
	items[3] = &PriorityQueueIntItem{Value: 30, Priority: 50, Index: 3}
	items[4] = &PriorityQueueIntItem{Value: 40, Priority: 30, Index: 4}

	pq := make(PriorityQueueInt, 5)
	for i := 0; i < 5; i++ {
		pq[i] = items[i]
	}

	heap.Init(&pq)

	pq.Update(items[2], 50, 100)

	results := make([]int, 5)
	for i := 0; i < 5; i++ {
		results[i] = heap.Pop(&pq).(*PriorityQueueIntItem).Value
	}

	assertEqual(t, []int{50, 30, 10, 40, 0}, results)
}

func assertEqual(t *testing.T, want, got interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want: %#v, got: %#v\n", want, got)
	}
}
