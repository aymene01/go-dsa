package trees

import (
	"reflect"
	"testing"
)

func TestTreeOrder(t *testing.T) {
	t.Run("test preorder",func(t *testing.T) {
		values := make([]int, 0)
		root := &Tree{
			value: 1,
			left: &Tree{
				value: 2,
			},
			right: &Tree{
				value: 3,
			},
		}

		preOrder(root, &values)
		expected := []int{1, 2, 3}

		if !reflect.DeepEqual(expected, values) {
			t.Errorf("expected: %v - got: %v", expected, values)
		}
	})

	t.Run("test postorder",func(t *testing.T) {
		values := make([]int, 0)
		root := &Tree{
			value: 1,
			left: &Tree{
				value: 2,
			},
			right: &Tree{
				value: 3,
			},
		}

		postOrder(root, &values)
		expected := []int{2, 3, 1}

		if !reflect.DeepEqual(expected, values) {
			t.Errorf("expected: %v - got: %v", expected, values)
		}
	})

	t.Run("test inorder",func(t *testing.T) {
		values := make([]int, 0)
		root := &Tree{
			value: 1,
			left: &Tree{
				value: 2,
			},
			right: &Tree{
				value: 3,
			},
		}

		inOrder(root, &values)
		expected := []int{2, 1, 3}

		if !reflect.DeepEqual(expected, values) {
			t.Errorf("expected: %v - got: %v", expected, values)
		}
	})
}