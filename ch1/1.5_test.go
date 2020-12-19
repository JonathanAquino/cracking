package ch1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// isOneAway checks if a and b are 0 or 1 edit away (edit being insertion, deletion, or
// replacement of a character). Assumes a and b are ASCII.
func isOneAway(a, b string) bool {
	aLen := len(a)
	bLen := len(b)
	if aLen <= bLen {
		// Ensure that a is longer
		temp := b
		b = a
		a = temp
		aLen = len(a)
		bLen = len(b)
	}
	editSeen := false
	j := 0
	for i := 0; i < aLen; i++ {
		if j == bLen || a[i] != b[j] {
			if editSeen {
				return false
			}
			editSeen = true
			// If there is a difference and the lengths are not the same,
			// don't move the index on the smaller string. Allow the longer
			// string to advance one character while staying where we are
			// on the smaller string.
			if aLen != bLen {
				continue
			}
		}
		j++
	}
	return true
}

func Test1Dot5(t *testing.T) {
	assert.True(t, isOneAway("pale", "ple"))
	assert.True(t, isOneAway("ple", "pale"))
	assert.True(t, isOneAway("pales", "pale"))
	assert.True(t, isOneAway("pale", "pales"))
	assert.True(t, isOneAway("pale", "bale"))
	assert.False(t, isOneAway("pale", "bake"))
	assert.False(t, isOneAway("palest", "pale"))
}
