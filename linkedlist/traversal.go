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

// two pointers (fast - slow)
func getMidleValue(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil  && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}

	return slow
}

func hasCycle(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			slow = head
			for slow != fast {
				slow = slow.next
				fast = fast.next
			}

			return slow
		}
	}
	return nil
}