package linkedlist

import (
	"reflect"
	"testing"
)

func TestTraversal(t *testing.T) {
	list, expected := generateLinkedList()
	result := linkedlistValues(list)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected: %v - got: %v", expected, result)
	}
}