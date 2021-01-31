package ch8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func steps(n int) int {
	if n <= 1 {
		return 1
	}
	ways := 0
	// Takes 1 step
	ways += steps(n - 1)
	// Takes 2 steps
	if n >= 2 {
		ways += steps(n - 2)
	}
	// Takes 3 steps
	if n >= 3 {
		ways += steps(n - 3)
	}
	return ways
}

func Test(t *testing.T) {
	assert.Equal(t, 13, steps(5))
}
