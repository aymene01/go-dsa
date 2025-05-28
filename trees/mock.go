package trees

func generateTree() *Tree {
	return &Tree{
		value: 1,
		left: &Tree{
			value: 2,
		},
		right: &Tree{
			value: 3,
		},
	}
}