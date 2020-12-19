package ch1

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// isPermutation checks if a is a permutation of b.
func isPermutation(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	aLetters := strings.Split(a, "")
	bLetters := strings.Split(b, "")
	sort.Strings(aLetters)
	sort.Strings(bLetters)
	for i, aLetter := range aLetters {
		bLetter := bLetters[i]
		if aLetter != bLetter {
			return false
		}
	}
	return true
}

func Test1Dot2(t *testing.T) {
	assert.False(t, isPermutation("a", "aa"))
	assert.False(t, isPermutation("hello", "hallo"))
	assert.True(t, isPermutation("hello", "olleh"))
}
