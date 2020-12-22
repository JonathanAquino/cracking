package ch4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// minimalTree takes a sorted array and returns a BST with minimal height.
func minimalTree(data []int) *BinaryTreeNode {
	if len(data) == 0 {
		return nil
	}
	centerIndex := len(data) / 2
	node := BinaryTreeNode{
		data:  data[centerIndex],
		left:  minimalTree(data[:centerIndex]),
		right: minimalTree(data[centerIndex+1:]),
	}
	return &node
}

func Test4Dot2(t *testing.T) {
	node := minimalTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	assert.Equal(t, 1, node.left.left.left.data)
	assert.Equal(t, 2, node.left.left.data)
	assert.Equal(t, 3, node.left.data)
	assert.Equal(t, 4, node.left.right.left.data)
	assert.Equal(t, 5, node.left.right.data)
	assert.Equal(t, 6, node.data)
	assert.Equal(t, 7, node.right.left.left.data)
	assert.Equal(t, 8, node.right.left.data)
	assert.Equal(t, 9, node.right.data)
	assert.Equal(t, 10, node.right.right.data)
}
