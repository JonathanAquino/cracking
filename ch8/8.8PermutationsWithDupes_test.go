package ch8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// permutations returns all permutations of a string of characters (not necessarily
// unique). Duplicates are not returned. s must be ASCII.
func permutationsWithDupes(s string) []string {
	permutations := permutations(s)
	uniquePermutations := map[string]bool{}
	for _, permutation := range permutations {
		uniquePermutations[permutation] = true
	}
	permutations = []string{}
	for permutation := range uniquePermutations {
		permutations = append(permutations, permutation)
	}
	return permutations
}

func TestPermutationsWithDupes(t *testing.T) {
	expected := []string{"aac", "aca", "caa"}
	assert.Equal(t, expected, permutationsWithDupes("aac"))
}
