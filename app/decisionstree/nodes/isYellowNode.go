package nodes

func NewIsYellowNode () *Node {
	return NewNode(isYellow)
}

func isYellow(c map[string]string) bool {
	return c["color"] == "yellow"
}
