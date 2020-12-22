package ch4

import (
	"testing"

	"github.com/adam-lavrik/go-imath/ix"
	"github.com/stretchr/testify/assert"
)

// isBalanced returns true if the tree is balanced, where balanced is defined
// to be the heights of the subtrees of any node differ by no more than one.
func isBalanced(node *BinaryTreeNode) bool {
	if node == nil {
		return true
	}
	return ix.Abs(height(node.left)-height(node.right)) <= 1 && isBalanced(node.left) && isBalanced(node.right)
}

// height returns the height of the given tree.
func height(node *BinaryTreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + ix.Max(height(node.left), height(node.right))
}

func TestIsBalancedReturnsTrue(t *testing.T) {
	node1 := &BinaryTreeNode{data: 1}
	node2 := &BinaryTreeNode{data: 2}
	node3 := &BinaryTreeNode{data: 3}
	node4 := &BinaryTreeNode{data: 4}

	node1.left = node2
	node1.right = node3
	node2.left = node4

	assert.True(t, isBalanced(node1))
}

func TestIsBalancedReturnsFalse(t *testing.T) {
	node1 := &BinaryTreeNode{data: 1}
	node2 := &BinaryTreeNode{data: 2}
	node3 := &BinaryTreeNode{data: 3}
	node4 := &BinaryTreeNode{data: 4}
	node5 := &BinaryTreeNode{data: 5}

	node1.left = node2
	node1.right = node3
	node2.left = node4
	node4.right = node5

	assert.False(t, isBalanced(node1))
}
