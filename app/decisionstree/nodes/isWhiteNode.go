package nodes

func NewIsWhiteNode () *Node {
	return NewNode(isWhite)
}

func isWhite(c map[string]string) bool {
	return c["color"] == "white"
}
