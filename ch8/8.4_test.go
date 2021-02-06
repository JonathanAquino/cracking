package ch8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// allSubsets returns all subsets of the given set.
func allSubsets(a []int) [][]int {
	// Start with the empty set.
	subsets := [][]int{{}}
	for i := 0; i < len(a); i++ {
		childSubsets := allSubsets(a[i+1:])
		for _, childSubset := range childSubsets {
			// Prepend the element to each child subset.
			subset := append([]int{a[i]}, childSubset...)
			subsets = append(subsets, subset)
		}
	}
	return subsets
}

func Test8Dot4(t *testing.T) {
	actual := allSubsets([]int{1, 2, 3})
	expected := [][]int{
		{},
		{1},
		{1, 2},
		{1, 2, 3},
		{1, 3},
		{2},
		{2, 3},
		{3},
	}
	assert.Equal(t, expected, actual)
}
