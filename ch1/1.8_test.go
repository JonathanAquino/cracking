package ch1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func zeroMatrix(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	rowsToZero := []int{}
	colsToZero := []int{}
	rowCount := len(matrix)
	colCount := len(matrix[0])
	for j := 0; j < rowCount; j++ {
		for i := 0; i < colCount; i++ {
			if matrix[j][i] == 0 {
				rowsToZero = append(rowsToZero, j)
				colsToZero = append(colsToZero, i)
			}
		}
	}
	for _, y := range rowsToZero {
		for i := 0; i < colCount; i++ {
			matrix[y][i] = 0
		}
	}
	for _, x := range colsToZero {
		for j := 0; j < rowCount; j++ {
			matrix[j][x] = 0
		}
	}
}

func Test1Dot8(t *testing.T) {
	matrix := [][]int{
		[]int{1, 1, 1, 2, 2, 2},
		[]int{1, 1, 1, 2, 0, 2},
		[]int{1, 1, 1, 2, 2, 2},
		[]int{3, 0, 3, 4, 4, 4},
	}
	expected := [][]int{
		[]int{1, 0, 1, 2, 0, 2},
		[]int{0, 0, 0, 0, 0, 0},
		[]int{1, 0, 1, 2, 0, 2},
		[]int{0, 0, 0, 0, 0, 0},
	}
	zeroMatrix(matrix)
	assert.Equal(t, expected, matrix)
}
