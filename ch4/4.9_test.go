package ch4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// bstSequences returns all possible arrays of data that could have been used to
// construct the given binary search tree
func bstSequences(node *BinaryTreeNode) [][]int {
	if node == nil {
		return [][]int{}
	}
	if node.left == nil && node.right == nil {
		return [][]int{[]int{node.data}}
	}
	// Get the sequences for the left node.
	leftBstSequences := bstSequences(node.left)
	// Get the sequences for the right node.
	rightBstSequences := bstSequences(node.right)
	// We will get all possible pairs of the left and right sequences, and for each
	// pair we will find all possible ways to combine the items while preserving
	// the order in each sequence. This is because when building the tree, the order
	// of each of the two sequences is important to ensure that the nodes go down
	// in order, but the order of choosing the left or right sequence is not important
	// (you can freely switch between building the left subtree vs. the right).
	combinedSequences := combine(leftBstSequences, rightBstSequences)
	// Prepend the node to each of the combined sequences.
	combinedSequencesWithNode := [][]int{}
	for _, combinedSequence := range combinedSequences {
		combinedSequenceWithNode := append([]int{node.data}, combinedSequence...)
		combinedSequencesWithNode = append(combinedSequencesWithNode, combinedSequenceWithNode)
	}
	return combinedSequencesWithNode
}

// combine gets all possible pairs of the left and right sequences, and for each
// pair it will find all possible ways to combine the items while preserving
// the order in each sequence.
func combine(leftBstSequences, rightBstSequences [][]int) [][]int {
	combinedSequences := [][]int{}
	for _, leftBstSequence := range leftBstSequences {
		for _, rightBstSequence := range rightBstSequences {
			combinedSequences = append(combinedSequences, combineTwoSequences(leftBstSequence, rightBstSequence)...)
		}
	}
	return combinedSequences
}

// combineTwoSequences finds all possible ways to combine the items while preserving
// the order in each sequence.
func combineTwoSequences(leftBstSequence, rightBstSequence []int) [][]int {
	if len(leftBstSequence) == 0 {
		return [][]int{rightBstSequence}
	}
	if len(rightBstSequence) == 0 {
		return [][]int{leftBstSequence}
	}
	sequences := [][]int{}
	// Insert the first right item at each point along the left sequence, then recurse.
	// Go to i <= len rather than i < len to include placing the first right item
	// at the end of the left sequence.
	for i := 0; i <= len(leftBstSequence); i++ {
		suffixes := combineTwoSequences(leftBstSequence[i:], rightBstSequence[1:])
		prefix := append(copyIntSlice(leftBstSequence[:i]), rightBstSequence[0])
		for _, suffix := range suffixes {
			sequences = append(sequences, append(copyIntSlice(prefix), suffix...))
		}
	}
	return sequences
}

// copyIntSlice makes a copy of an int slice to ensure that the original slice will not
// be changed by operations on the new slice.
func copyIntSlice(s []int) []int {
	copy := []int{}
	for _, item := range s {
		copy = append(copy, item)
	}
	return copy
}

func TestBSTSequences(t *testing.T) {
	node1 := &BinaryTreeNode{data: 1}
	node2 := &BinaryTreeNode{data: 2}
	node3 := &BinaryTreeNode{data: 3}

	node2.left = node1
	node2.right = node3
	expected := [][]int{
		[]int{2, 3, 1},
		[]int{2, 1, 3},
	}
	assert.Equal(t, expected, bstSequences(node2))
}

func TestBSTSequences2(t *testing.T) {
	node1 := &BinaryTreeNode{data: 1}
	node2 := &BinaryTreeNode{data: 2}
	node3 := &BinaryTreeNode{data: 3}
	node4 := &BinaryTreeNode{data: 4}
	node5 := &BinaryTreeNode{data: 5}

	node4.left = node2
	node4.right = node5
	node2.left = node1
	node2.right = node3
	expected := [][]int{
		[]int{4, 5, 2, 3, 1},
		[]int{4, 2, 5, 3, 1},
		[]int{4, 2, 3, 5, 1},
		[]int{4, 2, 3, 1, 5},
		[]int{4, 5, 2, 1, 3},
		[]int{4, 2, 5, 1, 3},
		[]int{4, 2, 1, 5, 3},
		[]int{4, 2, 1, 3, 5},
	}
	assert.Equal(t, expected, bstSequences(node4))
}
