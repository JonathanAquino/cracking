package ch1

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// isPalindromePermutation checks if the string is a permutation of a palindrome.
func isPalindromePermutation(s string) bool {
	s = strings.ToLower(s)
	charToCount := map[string]int{}
	for _, char := range strings.Split(s, "") {
		if char == " " {
			continue
		}
		charToCount[char]++
	}
	numCharsWithOddCounts := 0
	for _, count := range charToCount {
		if count%2 != 0 {
			numCharsWithOddCounts++
		}
	}
	return numCharsWithOddCounts <= 1
}

func Test1Dot4(t *testing.T) {
	assert.True(t, isPalindromePermutation("Tact Coa"))
	assert.True(t, isPalindromePermutation("Ab a"))
	assert.True(t, isPalindromePermutation("Abb a"))
	assert.False(t, isPalindromePermutation("Ab ad"))
	assert.True(t, isPalindromePermutation("Abb ad"))
}
