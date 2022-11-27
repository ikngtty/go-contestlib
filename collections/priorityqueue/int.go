package priorityqueue

import "container/heap"

// An IntPriorityQueueItem is something we manage in a priority queue.
type IntPriorityQueueItem struct {
	Value    int // The value of the item; arbitrary.
	Priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	Index int // The index of the item in the heap.
}

// A IntPriorityQueue implements heap.Interface and holds Items.
type IntPriorityQueue []*IntPriorityQueueItem

func (pq IntPriorityQueue) Len() int { return len(pq) }

func (pq IntPriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Priority > pq[j].Priority
}

func (pq IntPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *IntPriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*IntPriorityQueueItem)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *IntPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Update modifies the priority and value of an Item in the queue.
func (pq *IntPriorityQueue) Update(item *IntPriorityQueueItem, value int, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}
