package ch5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// conversion returns the number of bits to flip to convert a to b.
func conversion(a, b int) int {
	n := 0
	// Assume these are 32-bit integers
	for i := 0; i < 32; i++ {
		if getBit(a, i) != getBit(b, i) {
			n++
		}
	}
	return n
}

func TestConversion(t *testing.T) {
	assert.Equal(t, 2, conversion(0b11101, 0b01111))
}
