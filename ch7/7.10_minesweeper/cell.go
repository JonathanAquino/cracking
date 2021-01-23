package main

import "strconv"

// Cell is a spot on the Minesweeper board.
type Cell struct {
	// Uncovered is whether the Cell has been exposed.
	Uncovered bool

	// Flagged is whether the Cell has been flagged.
	Flagged bool

	// Bomb is whether the cell has a bomb
	Bomb bool

	// AdjacentBombs is the number of surrounding cells that have a bomb
	// (up to 8).
	AdjacentBombs int

	Board *Board
	Row   int
	Col   int
}

// Symbol returns a 1-character string for rendering the cell.
func (c *Cell) Symbol() string {
	if c.Flagged {
		return "F"
	}
	if !c.Uncovered {
		return "?"
	}
	if c.Bomb {
		return "*"
	}
	if c.AdjacentBombs > 0 {
		return strconv.Itoa(c.AdjacentBombs)
	}
	return " "
}

// IsBlank returns whether this is a blank cell.
func (c *Cell) IsBlank() bool {
	if c.Bomb {
		return false
	}
	return c.AdjacentBombs == 0
}

// IsNumber returns whether this is a numbered cell.
func (c *Cell) IsNumber() bool {
	if c.Bomb {
		return false
	}
	return c.AdjacentBombs > 0
}

// Neighbours returns the (up to) 8 cells around the given cell.
func (c *Cell) Neighbours() []*Cell {
	neighbours := []*Cell{}
	for i := max(c.Row-1, 1); i <= min(c.Row+1, c.Board.SideLength); i++ {
		for j := max(c.Col-1, 1); j <= min(c.Col+1, c.Board.SideLength); j++ {
			if i == c.Row && j == c.Col {
				continue
			}
			neighbours = append(neighbours, c.Board.Cell(i, j))
		}
	}
	return neighbours
}

// max returns the maximum of the two ints.
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// min returns the minimum of the two ints.
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
