package ch8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// PaintFill fills the point and the surrounding area with newColor
func PaintFill(image [][]rune, row, col int, newColor rune) {
	paintFillHelper(image, row, col, image[row][col], newColor)
}

func paintFillHelper(image [][]rune, row, col int, oldColor, newColor rune) {
	if image[row][col] != oldColor {
		return
	}
	image[row][col] = newColor
	printImage(image)
	if row > 0 {
		paintFillHelper(image, row-1, col, oldColor, newColor)
	}
	if col > 0 {
		paintFillHelper(image, row, col-1, oldColor, newColor)
	}
	if row < len(image)-1 {
		paintFillHelper(image, row+1, col, oldColor, newColor)
	}
	if col < len(image[0])-1 {
		paintFillHelper(image, row, col+1, oldColor, newColor)
	}
}

func printImage(image [][]rune) {
	for _, row := range image {
		for _, col := range row {
			print(string(col))
		}
		println()
	}
	println()
}

func TestPaintFill(t *testing.T) {
	image := [][]rune{
		{'x', 'x', ' ', ' ', ' '},
		{'x', ' ', ' ', ' ', 'x'},
		{' ', ' ', ' ', 'x', ' '},
		{' ', ' ', 'x', ' ', ' '},
	}
	expected := [][]rune{
		{'x', 'x', 'E', 'E', 'E'},
		{'x', 'E', 'E', 'E', 'x'},
		{'E', 'E', 'E', 'x', ' '},
		{'E', 'E', 'x', ' ', ' '},
	}
	printImage(image)
	PaintFill(image, 1, 2, 'E')
	assert.Equal(t, expected, image)
}
