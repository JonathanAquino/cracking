package ch8

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// permutations returns all permutations of a string of characters (not necessarily
// unique). Duplicates are not returned. s must be ASCII.
func permutationsWithDupes(s string) []string {
	results := permutations(s)
	sort.Strings(results)
	if len(results) < 2 {
		return results
	}
	uniqueResults := []string{results[0]}
	for i := 1; i < len(results); i++ {
		if results[i] != results[i-1] {
			uniqueResults = append(uniqueResults, results[i])
		}
	}
	return uniqueResults
}

func TestPermutationsWithDupes(t *testing.T) {
	expected := []string{"aac", "aca", "caa"}
	assert.Equal(t, expected, permutationsWithDupes("aac"))
}
