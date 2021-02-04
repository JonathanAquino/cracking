package ch8

import (
	"testing"

	"github.com/adam-lavrik/go-imath/ix"

	"github.com/stretchr/testify/assert"
)

// findPath marks a path from start (S) to finish (F) with asterisks.
// In the grid, X indicates barriers, while blank indicates no barriers.
func findPath(grid [][]string) [][]string {
	path := getPath(grid, Point{row: 0, col: 0})
	grid = clearVisitedMarkers(grid)
	for _, point := range path {
		grid[point.row][point.col] = "*"
	}
	rowCount := len(grid)
	colCount := len(grid[0])
	grid[rowCount-1][colCount-1] = "F"
	return grid
}

// Prints the grid to stdout.
func printGrid(grid [][]string) {
	for _, row := range grid {
		for _, col := range row {
			print(col)
		}
		print("\n")
	}
	println("\n")
}

// clearVisitedMarkers blanks out the "." markers indicating that a cell was visited.
func clearVisitedMarkers(grid [][]string) [][]string {
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == "." {
				grid[i][j] = " "
			}
		}
	}
	return grid
}

// Point represents a location on the grid.
type Point struct {
	row, col int
}

// getPath returns a path of points from point to the finish marker ("F"),
// or an empty array if no path was found.
func getPath(grid [][]string, point Point) []Point {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			// Disallow (0, 0), (1, 1), (1, -1), (-1, 1), (-1, -1)
			if ix.Abs(i) == ix.Abs(j) {
				continue
			}
			newPoint := Point{point.row + i, point.col + j}
			if cell, ok := getCell(grid, newPoint); ok {
				switch cell {
				case "F":
					return []Point{newPoint}
				case " ":
					grid[newPoint.row][newPoint.col] = "."
					path := getPath(grid, newPoint)
					if len(path) == 0 {
						continue
					}
					return append(path, newPoint)
				case "X":
					continue
				case ".":
					continue
				case "S":
					continue
				default:
					panic("unrecognized cell: " + cell)
				}
			}
		}
	}
	return []Point{}
}

// getCell returns the value of the cell at the given point.
// Returns false if point is out of bounds.
func getCell(grid [][]string, point Point) (string, bool) {
	rowCount := len(grid)
	colCount := len(grid[0])
	if point.row < 0 || point.row >= rowCount {
		return "", false
	}
	if point.col < 0 || point.col >= colCount {
		return "", false
	}
	return grid[point.row][point.col], true
}

func Test8Dot2(t *testing.T) {
	// S = start
	// F = finish
	// X = barrier
	// . = visited
	grid := [][]string{
		[]string{"S", " ", " ", " ", " ", " ", " ", "X"},
		[]string{"X", " ", "X", "X", "X", "X", "X", "X"},
		[]string{"X", " ", " ", " ", " ", " ", " ", " "},
		[]string{" ", " ", "X", "X", " ", "X", "X", "X"},
		[]string{" ", "X", " ", " ", " ", " ", " ", "F"},
	}
	expected := [][]string{
		[]string{"S", "*", " ", " ", " ", " ", " ", "X"},
		[]string{"X", "*", "X", "X", "X", "X", "X", "X"},
		[]string{"X", "*", "*", "*", "*", " ", " ", " "},
		[]string{" ", " ", "X", "X", "*", "X", "X", "X"},
		[]string{" ", "X", " ", " ", "*", "*", "*", "F"},
	}
	printGrid(grid)
	actual := findPath(grid)
	printGrid(actual)
	assert.Equal(t, expected, actual)
}
