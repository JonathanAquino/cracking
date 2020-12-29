package ch5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// overwriteBits overwrites bits i through j in n with m. Assumes that
// m fits in bits i through j.
func overwriteBits(n, m, i, j int) int {
	// Clear the bits from i through j
	mask := 0
	for x := i; x < j+1; x++ {
		mask = (mask << 1) | 1
	}
	mask = mask << i
	mask = ^mask
	n = n & mask
	// Shift m over by i
	m = m << i
	return n | m
}

func TestOverwriteBits(t *testing.T) {
	assert.Equal(t, 0b10001001100, overwriteBits(0b10000000000, 0b10011, 2, 6))
}
