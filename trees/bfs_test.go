package trees

import (
	"reflect"
	"testing"
)

func TestTreeBFS(t *testing.T){
	t.Run("should return the right order", func(t *testing.T) {
		root := generateTree()
		expected := []int{1, 2, 3}
		result := bfs(root)

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("expected: %v - got: %v", expected, result)
		}
	})
}