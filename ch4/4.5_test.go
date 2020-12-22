package ch4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// isBST returns whether the tree is a binary search tree
func isBST(node *BinaryTreeNode) bool {
	if node == nil {
		return true
	}
	if node.left != nil && node.left.data > node.data {
		return false
	}
	if node.right != nil && node.right.data < node.data {
		return false
	}
	return isBST(node.left) && isBST(node.right)
}

func TestIsBSTReturnsTrue(t *testing.T) {
	node1 := &BinaryTreeNode{data: 1}
	node2 := &BinaryTreeNode{data: 2}
	node3 := &BinaryTreeNode{data: 3}
	node4 := &BinaryTreeNode{data: 4}
	node5 := &BinaryTreeNode{data: 5}
	node6 := &BinaryTreeNode{data: 6}
	node7 := &BinaryTreeNode{data: 7}

	node4.left = node2
	node4.right = node6
	node2.left = node1
	node2.right = node3
	node6.left = node5
	node6.right = node7

	assert.True(t, isBST(node4))
}

func TestIsBSTReturnsFalse(t *testing.T) {
	node1 := &BinaryTreeNode{data: 1}
	node2 := &BinaryTreeNode{data: 2}
	node3 := &BinaryTreeNode{data: 3}
	node4 := &BinaryTreeNode{data: 4}
	node5 := &BinaryTreeNode{data: 5}
	node6 := &BinaryTreeNode{data: 6}
	node7 := &BinaryTreeNode{data: 7}

	node4.left = node2
	node4.right = node6
	node2.left = node1
	node2.right = node3
	node6.left = node7
	node6.right = node5

	assert.False(t, isBST(node4))
}
