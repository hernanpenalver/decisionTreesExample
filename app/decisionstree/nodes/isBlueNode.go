package nodes

func NewIsBlueNode () *Node {
	return NewNode(isBlue)
}

func isBlue(c map[string]string) bool {
	return c["color"] == "blue"
}
