package ch8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// isMagic returns true if it has an element such that a[i] == i.
// a is assumed to be a sorted array of distinct numbers.
func isMagic(a []int) bool {
	// Use binary search.
	return isMagicHelper(a, 0, len(a))
}

// isMagicHelper returns whether the array from lo (inclusive) to hi (exclusive)
// contains a magic index (a[i] == i).
func isMagicHelper(a []int, lo, hi int) bool {
	if hi <= lo {
		return false
	}
	mid := (hi + lo) / 2
	if a[mid] == mid {
		return true
	}
	if isMagicHelper(a, lo, mid) {
		return true
	}
	if lo == mid {
		return false
	}
	return isMagicHelper(a, mid, hi)
}

func Test8Dot3(t *testing.T) {
	assert.False(t, isMagic([]int{3, 4, 5, 6, 7}))
	assert.True(t, isMagic([]int{-10, -5, 0, 3, 7}))
}
