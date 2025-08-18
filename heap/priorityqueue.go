package heap

import "errors"

type PriorityQueue struct {
	heap Heap
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		heap: Heap{},
	}
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.heap.isEmpty()
}

func (pq *PriorityQueue) Enqueue(value int) {
	pq.heap.insert(value)
}

func (pq *PriorityQueue) Dequeue() (int, error) {
	value, err := pq.heap.pop()
	return value, err 
}

func (pq *PriorityQueue) Peek() (int, error) {
	if pq.heap.isEmpty() {
		return 0, errors.New("priority queue is empty")
	}
	return pq.heap.values[0], nil
}

func (pq *PriorityQueue) Size() int {
	return len(pq.heap.values)
}


