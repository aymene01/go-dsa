package heap

import (
	"errors"
)

type Heap struct {
	values []int
}

func(h *Heap) _bubbleUp(idx int){
	parent := (idx - 1) / 2

	for idx > 0 && h.values[idx] > h.values[parent] {
		tmp := h.values[parent]
		h.values[parent] = h.values[idx]
		h.values[idx] = tmp
		idx = parent
		parent = (idx - 1) / 2
	}
}

func(h *Heap) _bubbleDown(idx int){
	left, right := idx * 2 + 1, idx * 2 + 2
	largest := idx 

	if left < len(h.values) && h.values[left] > h.values[largest] {
		largest = left
	}

	if right < len(h.values) && h.values[right] > h.values[largest] {
		largest = right
	}

	if largest != idx {
		tmp := h.values[largest]
		h.values[largest] = h.values[idx]
		h.values[idx] = tmp
		h._bubbleDown(largest)
	}

}

func(h *Heap) isEmpty() bool {
	return len(h.values) == 0
}

func(h *Heap) insert(value int){
	h.values = append(h.values, value)
	h._bubbleUp(len(h.values) - 1)
}

func(h *Heap) pop() (error, int){
	if h.isEmpty() {
		return errors.New("the heap is empty nothing to pop"), 0
	}

	value := h.values[0]
	h.values[0] = h.values[len(h.values) - 1]
	h.values = h.values[:len(h.values) - 1]

	if !h.isEmpty() {
		h._bubbleDown(0)
	}

	return nil, value

}


type PriorityQueue struct {
	heap Heap
}

func (p *PriorityQueue) isEmpty() bool {
	return p.heap.isEmpty()
}

func (p *PriorityQueue) append(value int) {
	p.heap.insert(value)
}

func (p *PriorityQueue) pop() (error, int) {
	return p.heap.pop()
}


func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		heap: Heap{},
	}
}


