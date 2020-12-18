package ch1

import (
	"math"
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func rotate(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	n := len(matrix[0])
	maxJ := int(math.Ceil(float64(n)/2) - 1)
	maxI := n - 1 - 1
	minI := 0
	for j := 0; j <= maxJ; j++ {
		for i := minI; i <= maxI; i++ {
			// The first matrix index is the row; the second, the column.
			// Store lower left value in temp.
			temp := matrix[n-1-i][j]
			// Set lower left value to lower right value.
			matrix[n-1-i][j] = matrix[n-1-j][n-1-i]
			// Set lower right value to upper right value.
			matrix[n-1-j][n-1-i] = matrix[i][n-1-j]
			// Set upper right value to upper left value.
			matrix[i][n-1-j] = matrix[j][i]
			// Set upper left value to temp.
			matrix[j][i] = temp
		}
		minI++
		maxI--
	}
}

func Test1Dot7(t *testing.T) {
	matrix := [][]int{
		[]int{1, 1, 1, 2, 2, 2},
		[]int{1, 1, 1, 2, 2, 2},
		[]int{1, 1, 1, 2, 2, 2},
		[]int{3, 3, 3, 4, 4, 4},
		[]int{3, 3, 3, 4, 4, 4},
		[]int{3, 3, 3, 4, 4, 4},
	}
	expected := [][]int{
		[]int{3, 3, 3, 1, 1, 1},
		[]int{3, 3, 3, 1, 1, 1},
		[]int{3, 3, 3, 1, 1, 1},
		[]int{4, 4, 4, 2, 2, 2},
		[]int{4, 4, 4, 2, 2, 2},
		[]int{4, 4, 4, 2, 2, 2},
	}
	rotate(matrix)
	assert.Equal(t, expected, matrix)
}

func Test1Dot7b(t *testing.T) {
	matrix := [][]int{
		[]int{1, 1, 1, 5, 2, 2, 2},
		[]int{1, 1, 1, 5, 2, 2, 2},
		[]int{1, 1, 1, 5, 2, 2, 2},
		[]int{8, 8, 8, 9, 6, 6, 6},
		[]int{3, 3, 3, 7, 4, 4, 4},
		[]int{3, 3, 3, 7, 4, 4, 4},
		[]int{3, 3, 3, 7, 4, 4, 4},
	}
	expected := [][]int{
		[]int{3, 3, 3, 8, 1, 1, 1},
		[]int{3, 3, 3, 8, 1, 1, 1},
		[]int{3, 3, 3, 8, 1, 1, 1},
		[]int{7, 7, 7, 9, 5, 5, 5},
		[]int{4, 4, 4, 6, 2, 2, 2},
		[]int{4, 4, 4, 6, 2, 2, 2},
		[]int{4, 4, 4, 6, 2, 2, 2},
	}
	rotate(matrix)
	assert.Equal(t, expected, matrix)
}
