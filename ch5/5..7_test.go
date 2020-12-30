package ch5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// pairwiseSwap swaps odd and even bits with as few operations as possible.
// For example, bits 0 and 1 are swapped, bits 2 and 3 are swapped, etc.
func pairwiseSwap(x int) int {
	// Assume these are 32-bit integers
	odds := x & 0b01010101010101010101010101010101
	evens := x & 0b10101010101010101010101010101010
	return (odds << 1) | (evens >> 1)
}

func TestPairwiseSwap(t *testing.T) {
	assert.Equal(t, 0b100101, pairwiseSwap(0b11010))
}
