package ch5

import (
	"testing"

	"github.com/adam-lavrik/go-imath/ix"
	"github.com/stretchr/testify/assert"
)

// flipBitToWin returns the length of the longest sequence of 1s that occurs
// if we can flip one bit from 0 to 1. For example, given 0b11011101111 we
// return 8
func flipBitToWin(x int) int {
	maxSequenceOf1s := 0
	// Assume int has 32 bits. Try flipping each bit to 1 and count the longest
	// sequence of 1s.
	for i := 0; i < 32; i++ {
		mask := 1 << i
		candidate := x | mask
		maxSequenceOf1s = ix.Max(maxSequenceOf1s, findMaxSequenceOf1s(candidate))
	}
	return maxSequenceOf1s
}

// findMaxSequenceOf1s returns the length of the longest sequence of 1s in the
// binary representation of x.
func findMaxSequenceOf1s(x int) int {
	maxSequenceOf1s := 0
	currentSequenceOf1s := 0
	// Assume int has 32 bits.
	for i := 0; i < 32; i++ {
		mask := 1 << i
		if x&mask > 0 {
			currentSequenceOf1s++
		} else {
			maxSequenceOf1s = ix.Max(maxSequenceOf1s, currentSequenceOf1s)
			currentSequenceOf1s = 0
		}
	}
	maxSequenceOf1s = ix.Max(maxSequenceOf1s, currentSequenceOf1s)
	return maxSequenceOf1s
}

func TestFlipBitToWin(t *testing.T) {
	assert.Equal(t, 1, flipBitToWin(0))
	assert.Equal(t, 6, flipBitToWin(0b11111))
	assert.Equal(t, 8, flipBitToWin(0b11011101111))
}
