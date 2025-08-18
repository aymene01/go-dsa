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

func(h *Heap) pop() (int, error){
	if h.isEmpty() {
		return 0, errors.New("the heap is empty nothing to pop")
	}

	value := h.values[0]
	h.values[0] = h.values[len(h.values) - 1]
	h.values = h.values[:len(h.values) - 1]

	if !h.isEmpty() {
		h._bubbleDown(0)
	}

	return value, nil

}


