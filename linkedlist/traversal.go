package linkedlist

func linkedlistValues(head *ListNode) []int {
	result := make([]int, 0)
	curr := head

	for curr != nil {
		result = append(result, curr.value)
		curr = curr.next
	}

	return result
}