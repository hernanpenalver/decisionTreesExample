package nodes

func NewHeadNode () *Node {
	return NewNode(headCondition)
}

func headCondition(map[string]string) bool {
	return true
}
