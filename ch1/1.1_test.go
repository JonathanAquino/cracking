package ch1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// isUnique returns whether the string has all unique characters.
func isUnique(s string) bool {
	seen := map[rune]bool{}
	for _, c := range s {
		if seen[c] {
			return false
		}
		seen[c] = true
	}
	return true
}

// isUniqueB does the same as isUnique but without using additional data structures.
func isUniqueB(s string) bool {
	for i, c := range s {
		for j, other := range s {
			if j <= i {
				continue
			}
			if c == other {
				return false
			}
		}
	}
	return true
}

func Test1Dot1(t *testing.T) {
	assert.False(t, isUnique("foo"))
	assert.True(t, isUnique("sdfghjkl"))
}

func Test1Dot1B(t *testing.T) {
	assert.False(t, isUnique("foo"))
	assert.True(t, isUniqueB("sdfghjkl"))
}
