package decisionstree

import "awesomeProject/app/decisionstree/nodes"

func NewTree(context map[string]string) *Tree{
	headNode := nodes.NewHeadNode()
	hasColorNode := nodes.NewHasColorNode()
	isBlueNode := nodes.NewIsBlueNode()
	isYellowNode := nodes.NewIsYellowNode()
	isWhiteNode := nodes.NewIsWhiteNode()

	withoutColorLeafNode := nodes.NewLeafNode("without_color")
	undeterminedColorLeafNode := nodes.NewLeafNode("undetermined")
	blueLeafNode := nodes.NewLeafNode("blue")
	yellowLeafNode := nodes.NewLeafNode("yellow")
	whiteLeafNode := nodes.NewLeafNode("white")

	isWhiteNode.InsertSons(undeterminedColorLeafNode, whiteLeafNode)
	isYellowNode.InsertSons(isWhiteNode, yellowLeafNode)
	isBlueNode.InsertSons(isYellowNode, blueLeafNode)
	hasColorNode.InsertSons(withoutColorLeafNode, isBlueNode)
	headNode.InsertSons(hasColorNode, hasColorNode)

	return &Tree{
		context: context,
		root: headNode,
	}
}