package heap

import (
	"testing"
)

func TestHeap_Insert_And_Pop(t *testing.T) {
	h := &Heap{}

	// Test inserting single element
	h.insert(10)
	if h.isEmpty() {
		t.Error("Heap should not be empty after inserting element")
	}

	// Test popping single element
	err, value := h.pop()
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if value != 10 {
		t.Errorf("Expected 10, got %d", value)
	}
	if !h.isEmpty() {
		t.Error("Heap should be empty after popping last element")
	}
}

func TestHeap_MaxHeap_Property(t *testing.T) {
	h := &Heap{}
	
	// Insert elements in random order
	elements := []int{5, 10, 3, 8, 1, 15, 7}
	for _, elem := range elements {
		h.insert(elem)
	}

	// Pop all elements - should come out in descending order (max-heap)
	expected := []int{15, 10, 8, 7, 5, 3, 1}
	for i, expectedValue := range expected {
		if h.isEmpty() {
			t.Errorf("Heap should not be empty at iteration %d", i)
			break
		}
		
		err, value := h.pop()
		if err != nil {
			t.Errorf("Expected no error at iteration %d, got: %v", i, err)
		}
		if value != expectedValue {
			t.Errorf("At iteration %d: expected %d, got %d", i, expectedValue, value)
		}
	}

	if !h.isEmpty() {
		t.Error("Heap should be empty after popping all elements")
	}
}

func TestHeap_Pop_Empty_Heap(t *testing.T) {
	h := &Heap{}

	err, value := h.pop()
	if err == nil {
		t.Error("Expected error when popping from empty heap")
	}
	if value != 0 {
		t.Errorf("Expected 0 value when popping from empty heap, got %d", value)
	}
}

func TestHeap_IsEmpty(t *testing.T) {
	h := &Heap{}

	// Test empty heap
	if !h.isEmpty() {
		t.Error("New heap should be empty")
	}

	// Test after inserting
	h.insert(5)
	if h.isEmpty() {
		t.Error("Heap should not be empty after inserting")
	}

	// Test after popping
	h.pop()
	if !h.isEmpty() {
		t.Error("Heap should be empty after popping last element")
	}
}

func TestHeap_Large_Dataset(t *testing.T) {
	h := &Heap{}
	
	// Insert 100 elements
	for i := 1; i <= 100; i++ {
		h.insert(i)
	}

	// Pop all elements - should be in descending order
	lastValue := 101 // Start with value larger than any inserted
	for i := 0; i < 100; i++ {
		err, value := h.pop()
		if err != nil {
			t.Errorf("Unexpected error at iteration %d: %v", i, err)
		}
		if value >= lastValue {
			t.Errorf("Max-heap property violated: got %d after %d", value, lastValue)
		}
		lastValue = value
	}
}

func TestHeap_Duplicate_Values(t *testing.T) {
	h := &Heap{}
	
	// Insert duplicate values
	duplicates := []int{5, 5, 3, 8, 3, 8, 8}
	for _, elem := range duplicates {
		h.insert(elem)
	}

	// Pop all elements
	var result []int
	for !h.isEmpty() {
		err, value := h.pop()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		result = append(result, value)
	}

	// Check that we got all elements in non-ascending order
	expected := []int{8, 8, 8, 5, 5, 3, 3}
	if len(result) != len(expected) {
		t.Errorf("Expected %d elements, got %d", len(expected), len(result))
	}

	for i := 0; i < len(result) && i < len(expected); i++ {
		if result[i] != expected[i] {
			t.Errorf("At position %d: expected %d, got %d", i, expected[i], result[i])
		}
	}
}

func TestHeap_Single_Element_Operations(t *testing.T) {
	h := &Heap{}
	
	// Test multiple insert/pop cycles
	for i := 0; i < 5; i++ {
		h.insert(i * 10)
		err, value := h.pop()
		if err != nil {
			t.Errorf("Unexpected error at cycle %d: %v", i, err)
		}
		if value != i*10 {
			t.Errorf("At cycle %d: expected %d, got %d", i, i*10, value)
		}
		if !h.isEmpty() {
			t.Errorf("Heap should be empty after cycle %d", i)
		}
	}
}