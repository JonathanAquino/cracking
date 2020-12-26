package ch4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// checkSubtree returns whether the needle tree is a subtree of haystack tree,
// comparing values.
func checkSubtree(haystack *BinaryTreeNode, needle *BinaryTreeNode) bool {
	if haystack == nil {
		return needle == nil
	}
	return treesEqual(haystack, needle) ||
		checkSubtree(haystack.left, needle) ||
		checkSubtree(haystack.right, needle)
}

// treesEqual returns whether the two trees have the same values.
func treesEqual(a *BinaryTreeNode, b *BinaryTreeNode) bool {
	if a == nil {
		return b == nil
	}
	return a.data == b.data && treesEqual(a.left, b.left) && treesEqual(a.right, b.right)
}

func TestCheckSubtree1(t *testing.T) {
	node1 := &BinaryTreeNode{data: 1}
	node2 := &BinaryTreeNode{data: 2}
	node3 := &BinaryTreeNode{data: 3}

	node4 := &BinaryTreeNode{data: 4}
	node5 := &BinaryTreeNode{data: 5}
	node6 := &BinaryTreeNode{data: 6}

	node1.left = node2
	node1.right = node3
	node4.left = node5
	node4.right = node6
	assert.False(t, checkSubtree(node1, node4))
}

func TestCheckSubtree2(t *testing.T) {
	node1 := &BinaryTreeNode{data: 1}
	node2 := &BinaryTreeNode{data: 2}
	node3 := &BinaryTreeNode{data: 3}

	node4a := &BinaryTreeNode{data: 4}
	node5a := &BinaryTreeNode{data: 5}
	node6a := &BinaryTreeNode{data: 6}

	node4b := &BinaryTreeNode{data: 4}
	node5b := &BinaryTreeNode{data: 5}
	node6b := &BinaryTreeNode{data: 6}

	node1.left = node2
	node1.right = node3
	node4a.left = node5a
	node4a.right = node6a
	node3.right = node4a

	node4b.left = node5b
	node4b.right = node6b
	assert.True(t, checkSubtree(node1, node4b))
}
