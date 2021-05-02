package nodes

func NewHasColorNode () *Node {
	return NewNode(hasColor)
}

func hasColor(c map[string]string) bool {
	return c["color"] != ""
}
