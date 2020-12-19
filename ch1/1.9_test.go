package ch1

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// isRotation checks if a is a rotation of b using a single call to isSubstring.
func isRotation(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	return isSubstring(a, b+b)
}

func isSubstring(needle string, haystack string) bool {
	return strings.Contains(haystack, needle)
}

func TestIsSubstring(t *testing.T) {
	assert.True(t, isSubstring("ell", "hello"))
	assert.False(t, isSubstring("hello", "ell"))
}

func Test1Dot9(t *testing.T) {
	assert.True(t, isRotation("waterbottle", "erbottlewat"))
	assert.True(t, isRotation("erbottlewat", "waterbottle"))
	assert.False(t, isRotation("waterbottlX", "erbottlewat"))
	assert.False(t, isRotation("erbottlewat", "waterbottlX"))
	assert.False(t, isRotation("waterbottl", "erbottlewat"))
	assert.False(t, isRotation("erbottlewat", "waterbottl"))
}
