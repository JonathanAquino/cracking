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
	// Combine the sequences by getting all pairs of left and right sequences
	// and finding all permutations that preserve the order within the left and
	// right sequences. We want to preserve the order because the array order
	// is important, but elements can be chosen from either array in any order.
	combinedSequences := combine(leftBstSequences, rightBstSequences)
	// Prepend the node to each of the combined sequences.
	combinedSequencesWithNode := [][]int{}
	for _, combinedSequence := range combinedSequences {
		combinedSequenceWithNode := append([]int{node.data}, combinedSequence...)
		combinedSequencesWithNode = append(combinedSequencesWithNode, combinedSequenceWithNode)
	}
	return combinedSequencesWithNode
}

// combine combines the sequences by getting all pairs of left and right sequences
// and finding all permutations that preserve the order within the left and
// right sequences.
func combine(leftBstSequences, rightBstSequences [][]int) [][]int {
	combinedSequences := [][]int{}
	for _, leftBstSequence := range leftBstSequences {
		for _, rightBstSequence := range rightBstSequences {
			combinedSequences = append(combinedSequences, combineTwoSequences(leftBstSequence, rightBstSequence)...)
		}
	}
	return combinedSequences
}

// combineTwoSequences combines the sequences by finding all permutations of
// the elements from the two sequences that preserve the order within each of
// the two sequences.
func combineTwoSequences(leftBstSequence, rightBstSequence []int) [][]int {
	sequences := [][]int{}
	// Use bit sequences to capture all possible left-right decisions.
	i := 0
	// Stop when the number of bits is greater than the length of the left sequence.
	// Because then we will start from the beginning: it will be like the all-zero case.
	max := 1 << len(leftBstSequence)
	for i < max {
		currSequence := []int{}
		bitIndex := 0
		leftIndex := 0
		rightIndex := 0
		for true {
			bit := (i >> bitIndex) & 1
			if bit == 0 && leftIndex < len(leftBstSequence) {
				currSequence = append(currSequence, leftBstSequence[leftIndex])
				leftIndex++
				continue
			}
			if bit == 1 && rightIndex < len(rightBstSequence) {
				currSequence = append(currSequence, rightBstSequence[rightIndex])
				rightIndex++
				continue
			}
			if bit == 0 && leftIndex == len(leftBstSequence) {
				currSequence = append(currSequence, rightBstSequence[rightIndex:]...)
				break
			}
			if bit == 1 && rightIndex == len(rightBstSequence) {
				currSequence = append(currSequence, leftBstSequence[leftIndex:]...)
				break
			}
			bitIndex++
		}
		sequences = append(sequences, currSequence)
		i++
	}
	return sequences
}

// Returns whether the two slices have the same values in the same order.
func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestBSTSequences(t *testing.T) {
	node1 := &BinaryTreeNode{data: 1}
	node2 := &BinaryTreeNode{data: 2}
	node3 := &BinaryTreeNode{data: 3}

	node2.left = node1
	node2.right = node3
	expected := [][]int{
		[]int{2, 1, 3},
		[]int{2, 3, 1},
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
		[]int{4, 2, 1, 3, 5},
		[]int{4, 2, 3, 1, 5},
		[]int{4, 2, 1, 5, 3},
		[]int{4, 2, 3, 5, 1},
		[]int{4, 2, 5, 1, 3},
		[]int{4, 2, 5, 3, 1},
		[]int{4, 2, 3, 5, 1},
		[]int{4, 5, 2, 1, 3},
		[]int{4, 5, 2, 3, 1},
	}
	assert.Equal(t, expected, bstSequences(node4))
}
