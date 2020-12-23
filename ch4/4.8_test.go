package ch4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// firstCommonAncestor returns the lowest ancestor common to both nodes.
func firstCommonAncestor(a *BinaryTreeNode, b *BinaryTreeNode) *BinaryTreeNode {
	// Find root node.
	curr := a
	for curr.parent != nil {
		curr = curr.parent
	}
	root := curr
	// Clear visited flag using BFS.
	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		queue[0].visited = false
		if queue[0].left != nil {
			queue = append(queue, queue[0].left)
		}
		if queue[0].right != nil {
			queue = append(queue, queue[0].right)
		}
		queue = queue[1:]
	}
	// Mark a's ancestors as visited.
	curr = a
	for curr != nil {
		curr.visited = true
		curr = curr.parent
	}
	// Walk up b's ancestors until we find a visited node.
	curr = b
	for true {
		if curr.visited {
			return curr
		}
		curr = curr.parent
	}
	panic("shouldn't get here")
}

func TestFirstCommonAncestor(t *testing.T) {
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

	assert.Equal(t, node4, firstCommonAncestor(node1, node7))
	assert.Equal(t, node2, firstCommonAncestor(node1, node3))
	assert.Equal(t, node4, firstCommonAncestor(node1, node6))
}
