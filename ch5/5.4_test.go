package ch5

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// nextNumber returns the next smallest and next largest number whose binary
// representations have the same number of 1s as the original number.
func nextNumber(x int) (int, int, error) {
	// Assume there are at most 32 bits.
	// Starting from the rightmost bit, look for a bit that we can move over
	// one position to the right.
	nextSmallest := 0
	for i := 1; true; i++ {
		if i == 32 {
			return 0, 0, errors.New("could not find a next smallest number with the same number of 1s")
		}
		if getBit(x, i) == 1 && getBit(x, i-1) == 0 {
			nextSmallest = setBit(setBit(x, i, 0), i-1, 1)
			break
		}
	}
	// Starting from the rightmost bit, look for a bit that we can move over
	// one position to the left.
	nextLargest := 0
	for i := 0; true; i++ {
		if i == 31 {
			return 0, 0, errors.New("could not find a next largest number with the same number of 1s")
		}
		if getBit(x, i) == 1 && getBit(x, i+1) == 0 {
			nextLargest = setBit(setBit(x, i, 0), i+1, 1)
			break
		}
	}
	return nextSmallest, nextLargest, nil
}

// getBit returns whether the ith bit (0-based, starting from the right) is a 1 or a 0.
func getBit(x, i int) int {
	mask := 1 << i
	if x&mask > 0 {
		return 1
	}
	return 0
}

// setBit sets the ith bit (0-based, starting from the right) to a 1 or a 0.
func setBit(x, i, b int) int {
	result := clearBit(x, i)
	if b == 0 {
		return result
	}
	mask := 1 << i
	return result | mask
}

// clearBit sets the ith bit (0-based, starting fromt the right) to 0
func clearBit(x, i int) int {
	mask := 1 << i
	mask = ^mask
	return x & mask
}

func TestGetBit(t *testing.T) {
	assert.Equal(t, 0, getBit(0b11010, 2))
	assert.Equal(t, 1, getBit(0b11010, 1))
}

func TestClearBit(t *testing.T) {
	assert.Equal(t, 0b11000, clearBit(0b11010, 1))
}

func TestSetBit(t *testing.T) {
	assert.Equal(t, 0b11000, setBit(0b11010, 1, 0))
	assert.Equal(t, 0b11011, setBit(0b11010, 0, 1))
}

func TestNextNumber1(t *testing.T) {
	nextSmallest, nextLargest, _ := nextNumber(0b11010)
	assert.Equal(t, 0b11001, nextSmallest)
	assert.Equal(t, 0b11100, nextLargest)
}

func TestNextNumber2(t *testing.T) {
	_, _, err := nextNumber(0)
	assert.NotNil(t, err)
}

func TestNextNumber3(t *testing.T) {
	_, _, err := nextNumber(0b11111)
	assert.NotNil(t, err)
}
