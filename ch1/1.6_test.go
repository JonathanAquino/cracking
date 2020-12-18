package ch1

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func compress(s string) string {
	compressed := ""
	currentLetter := ""
	currentLetterCount := 0
	for _, letter := range strings.Split(s, "") {
		if currentLetter != letter {
			if currentLetter != "" {
				compressed += fmt.Sprintf("%s%d", currentLetter, currentLetterCount)
			}
			currentLetter = letter
			currentLetterCount = 1
		} else {
			currentLetterCount++
		}
	}
	if currentLetter != "" {
		compressed += fmt.Sprintf("%s%d", currentLetter, currentLetterCount)
	}
	if len(compressed) >= len(s) {
		return s
	}
	return compressed
}

func Test1Dot6(t *testing.T) {
	assert.Equal(t, "a2b1c5a3", compress("aabcccccaaa"))
	assert.Equal(t, "", compress(""))
	assert.Equal(t, "abcde", compress("abcde"))
	assert.Equal(t, "aaAA", compress("aaAA"))
	assert.Equal(t, "a3A3", compress("aaaAAA"))
}
