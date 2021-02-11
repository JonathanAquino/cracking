package ch8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// permutations returns all permutations of a string of unique characters.
// s must be ASCII.
func permutations(s string) []string {
	if len(s) <= 1 {
		return []string{s}
	}
	results := []string{}
	for i := 0; i < len(s); i++ {
		// s must be ASCII to use s[...].
		subpermutations := permutations(s[0:i] + s[i+1:])
		for _, subpermutation := range subpermutations {
			permutation := string(s[i]) + subpermutation
			results = append(results, permutation)
		}
	}
	return results
}

func TestPermutations(t *testing.T) {
	expected := []string{"abc", "acb", "bac", "bca", "cab", "cba"}
	assert.Equal(t, expected, permutations("abc"))
}
