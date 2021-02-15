package ch8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// EightQueens returns grids showing all ways of arranging 8 queens on an 8x8 grid
// such that none of them share the same row, column, or diagonal. Since each
// queen must be in its own row, each grid is an int array with each index
// representing a row and each int representing a column.
func EightQueens() [][]int {
	basicArrangements := BasicArrangements()
	finalArrangements := [][]int{}
	for _, arrangement := range basicArrangements {
		if !HasDiagonalConflict(arrangement) {
			finalArrangements = append(finalArrangements, arrangement)
		}
	}
	return finalArrangements
}

// HasDiagonalConflict returns whether any of the queens are on the same diagonal.
// Each index of the arrangement represents a row and each int represents a column.
func HasDiagonalConflict(arrangement []int) bool {
	for q1Row, q1Col := range arrangement {
		for q2Row, q2Col := range arrangement {
			if q1Row == q2Row {
				// Same queen
				continue
			}
			slope := (float32(q2Col) - float32(q1Col)) / (float32(q2Row) - float32(q1Row))
			if slope == 1 || slope == -1 {
				return true
			}
		}
	}
	return false
}

// BasicArrangements returns grids showing all ways of arranging 8 queens on an 8x8 grid
// such that none of them share the same row or column. Since each
// queen must be in its own row, each grid is an int array with each index
// representing a row and each int representing a column.
func BasicArrangements() [][]int {
	return basicArrangementsHelper([]int{})
}

// basicArrangementsHelper, given the arrangement so far, returns several
// completions of the arrangement. Each index of the arrangement represents
// a row and each int represents a column.
func basicArrangementsHelper(arrangementSoFar []int) [][]int {
	if len(arrangementSoFar) == 8 {
		return [][]int{}
	}
	takenColumns := map[int]bool{}
	for _, takenColumn := range arrangementSoFar {
		takenColumns[takenColumn] = true
	}
	completions := [][]int{}
	for col := 0; col < 8; col++ {
		if !takenColumns[col] {
			if len(arrangementSoFar) == 7 {
				return [][]int{[]int{col}}
			}
			newArrangementSoFar := append(arrangementSoFar, col)
			subcompletions := basicArrangementsHelper(newArrangementSoFar)
			for _, subcompletion := range subcompletions {
				completion := append([]int{col}, subcompletion...)
				completions = append(completions, completion)
			}
		}
	}
	return completions
}

// Prints the grid represented by the arrangement. Each index of the arrangement
// represents a row and each int represents a column.
func printArrangement(arrangement []int) {
	for _, col := range arrangement {
		s := []rune{'.', '.', '.', '.', '.', '.', '.', '.'}
		s[col] = 'Q'
		println(string(s))
	}
}

func TestEightQueens(t *testing.T) {
	for _, arrangement := range EightQueens() {
		printArrangement(arrangement)
		println()
	}
}

func TestHasDiagonalConflict(t *testing.T) {
	assert.True(t, HasDiagonalConflict([]int{0, 1, 7}))
	assert.False(t, HasDiagonalConflict([]int{0, 4, 7}))
	assert.True(t, HasDiagonalConflict([]int{1, 0, 7}))
	assert.False(t, HasDiagonalConflict([]int{4, 2, 8, 5, 7, 1, 3, 6}))
}
