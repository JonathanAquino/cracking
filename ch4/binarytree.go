package ch4

// A BinaryTreeNode is a node in a binary tree.
type BinaryTreeNode struct {
	left    *BinaryTreeNode
	right   *BinaryTreeNode
	data    int
	parent  *BinaryTreeNode
	visited bool
}
