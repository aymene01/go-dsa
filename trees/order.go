package trees

func postOrder(root *Tree, values *[]int) {
	if root == nil {
		return
	}

	postOrder(root.left, values)
	postOrder(root.right, values)
	(*values) = append((*values), root.value)
}

func inOrder(root *Tree, values *[]int) {
	if root == nil {
		return
	}

	inOrder(root.left, values)
	(*values) = append((*values), root.value)
	inOrder(root.right, values)
}

func preOrder(root *Tree, values *[]int) {
	if root == nil {
		return
	}
	
	(*values) = append((*values), root.value)
	preOrder(root.left, values)
	preOrder(root.right, values)
}
