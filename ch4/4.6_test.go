package ch4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// isBST returns the next highest node, given a BST.
func successor(node *BinaryTreeNode) *BinaryTreeNode {
	if node == nil {
		return nil
	}
	if node.right != nil {
		curr := node.right
		for curr.left != nil {
			curr = curr.left
		}
		return curr
	}
	// Walk up ancestors.
	curr := node
	for true {
		if curr.parent == nil {
			return nil
		}
		if curr == curr.parent.left {
			return curr.parent
		}
		curr = curr.parent
	}
	return nil
}

func TestSuccessor(t *testing.T) {
	node1 := &BinaryTreeNode{data: 1}
	node2 := &BinaryTreeNode{data: 2}
	node3 := &BinaryTreeNode{data: 3}
	node4 := &BinaryTreeNode{data: 4}
	node5 := &BinaryTreeNode{data: 5}
	node6 := &BinaryTreeNode{data: 6}
	node7 := &BinaryTreeNode{data: 7}

	node4.left = node2
	node2.parent = node4
	node4.right = node6
	node6.parent = node4
	node2.left = node1
	node1.parent = node2
	node2.right = node3
	node3.parent = node2
	node6.left = node5
	node5.parent = node6
	node6.right = node7
	node7.parent = node6

	assert.Equal(t, node2, successor(node1))
	assert.Equal(t, node3, successor(node2))
	assert.Equal(t, node4, successor(node3))
	assert.Equal(t, node5, successor(node4))
	assert.Equal(t, node6, successor(node5))
	assert.Equal(t, node7, successor(node6))
	assert.Nil(t, successor(node7))
}
