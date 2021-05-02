package nodes

type Node struct {
	left      *Node
	right     *Node
	result    string
	condition func(map[string]string) bool
}

func NewNode(condition func(map[string]string) bool) *Node {
	return &Node{
		condition: condition,
	}
}

func NewLeafNode(result string) *Node {
	return &Node{
		result:    result,
	}
}

func (n *Node) InsertSons(left, right *Node) {
	n.left = left
	n.right = right
}

func (n *Node) Condition(c map[string]string) bool {
	return n.condition(c)
}

func (n *Node) Result() string {
	return n.result
}

func (n *Node) SetResult(result string) {
	n.result = result
}

func (n *Node) Right() *Node {
	return n.right
}

func (n *Node) SetRight(right *Node) {
	n.right = right
}

func (n *Node) Left() *Node {
	return n.left
}

func (n *Node) SetLeft(left *Node) {
	n.left = left
}
