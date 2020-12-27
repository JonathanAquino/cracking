package ch4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// pathsWithSum returns the number of paths that sum to the given value,
// given a binary tree.
func pathsWithSum(node *BinaryTreeNode, desiredSum int) int {
	if node == nil {
		return 0
	}
	pathsWithSum := 0
	// BFS to iterate over all nodes
	queue := []*BinaryTreeNode{node}
	for len(queue) > 0 {
		curr := queue[0]
		pathsWithSum += pathsWithSumStartingAt(curr, desiredSum)
		if curr.left != nil {
			queue = append(queue, curr.left)
		}
		if curr.right != nil {
			queue = append(queue, curr.right)
		}
		queue = queue[1:]
	}
	return pathsWithSum
}

// pathsWithSumStartingAt returns the number of paths that sum to the given value,
// starting at the given node.
func pathsWithSumStartingAt(node *BinaryTreeNode, desiredSum int) int {
	pathsWithSum := 0
	if node == nil {
		return 0
	}
	if node.data == desiredSum {
		pathsWithSum++
	}
	pathsWithSum += pathsWithSumStartingAt(node.left, desiredSum-node.data) +
		pathsWithSumStartingAt(node.right, desiredSum-node.data)
	return pathsWithSum
}

func TestPathsWithSum(t *testing.T) {
	//         1
	//        1 1
	//       1   9
	//      5     -6
	//    -2 -2
	nodeA := &BinaryTreeNode{data: 1}
	nodeB := &BinaryTreeNode{data: 1}
	nodeC := &BinaryTreeNode{data: 1}
	nodeD := &BinaryTreeNode{data: 1}
	nodeE := &BinaryTreeNode{data: 9}
	nodeF := &BinaryTreeNode{data: 5}
	nodeG := &BinaryTreeNode{data: -6}
	nodeH := &BinaryTreeNode{data: -2}
	nodeI := &BinaryTreeNode{data: -2}

	nodeA.left = nodeB
	nodeA.right = nodeC
	nodeB.left = nodeD
	nodeD.left = nodeF
	nodeF.left = nodeH
	nodeF.right = nodeI
	nodeC.right = nodeE
	nodeE.right = nodeG

	assert.Equal(t, 4, pathsWithSum(nodeA, 3))
}
