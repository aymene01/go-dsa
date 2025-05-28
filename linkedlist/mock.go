package linkedlist

func generateLinkedList() (*ListNode, []int) {
	return &ListNode{
		value: 10,
		next: &ListNode{
			value: 40,
			next: &ListNode{
				value: 50,
			},
		},
	}, []int{10, 40, 50}
}