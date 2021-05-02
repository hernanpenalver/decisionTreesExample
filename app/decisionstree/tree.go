package decisionstree

import (
	"awesomeProject/app/decisionstree/nodes"
)

type Tree struct {
	context map[string]string
	root *nodes.Node
}

func (t Tree) Run()  string {
	node := printPreOrder(&t)

	return node.Result()
}

func printPreOrder(t *Tree) *nodes.Node {
	if t == nil {
		panic("error")
	}

	if t.root.Left() == nil || t.root.Right() == nil {
		return t.root
	}

	if t.root.Condition(t.context) {
		t.root = t.root.Right()
		return printPreOrder(t)
	}

	t.root = t.root.Left()
	return printPreOrder(t)
}
