package main

import "fmt"

// Board is a Minesweeper board
type Board struct {
	SideLength int
	Cells      *[][]*Cell
}

// NewBoard returns an initialized Minesweeper board.
func NewBoard(sideLength int) *Board {
	board := Board{SideLength: sideLength}
	cells := [][]*Cell{}
	for i := 0; i < sideLength; i++ {
		row := []*Cell{}
		for j := 0; j < sideLength; j++ {
			row = append(row, &Cell{Board: &board, Row: i + 1, Col: j + 1})
		}
		cells = append(cells, row)
	}
	board.Cells = &cells
	return &board
}

// Print displays the board in stdout.
func (b *Board) Print() {
	edge := " "
	for i := 0; i < b.SideLength; i++ {
		edge += "+"
	}
	println(edge)
	for _, row := range *b.Cells {
		line := "+"
		for _, cell := range row {
			line += cell.Symbol()
		}
		line += "+"
		println(line)
	}
	println(edge)
}

// OutOfBounds returns whether the coordinates are out of bounds.
func (b *Board) OutOfBounds(row, col int) bool {
	return row < 1 || col < 1 || row > b.SideLength || col > b.SideLength
}

// Uncover exposes the given cell.
func (b *Board) Uncover(row, col int) error {
	if b.OutOfBounds(row, col) {
		return fmt.Errorf("%d, %d is out of bounds", row, col)
	}
	cell := b.Cell(row, col)
	if cell.Flagged {
		return fmt.Errorf("%d, %d is flagged and can't be uncovered", row, col)
	}
	cell.Uncovered = true
	if cell.Bomb || cell.IsNumber() {
		return nil
	}
	// Keep uncovering neighbouring blank cells until we reach numbered cells.
	queue := []*Cell{cell}
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		head.Uncovered = true
		head.Flagged = false
		if head.IsBlank() {
			for _, neighbour := range head.Neighbours() {
				if neighbour.Uncovered {
					continue
				}
				if neighbour.IsBlank() || neighbour.IsNumber() {
					queue = append(queue, neighbour)
				}
			}
		}
	}
	return nil
}

// Flag flags the given cell.
func (b *Board) Flag(row, col int) error {
	if b.OutOfBounds(row, col) {
		return fmt.Errorf("%d, %d is out of bounds", row, col)
	}
	cell := b.Cell(row, col)
	if cell.Uncovered {
		return fmt.Errorf("%d, %d is already uncovered", row, col)
	}
	cell.Flagged = !cell.Flagged
	return nil
}

// Cell returns the cell at the given row and column.
func (b *Board) Cell(row, col int) *Cell {
	return (*b.Cells)[row-1][col-1]
}

// NumCells returns the number of cells on the board.
func (b *Board) NumCells() int {
	return b.SideLength * b.SideLength
}

// NumUncovered returns the number of cells that have been uncovered.
func (b *Board) NumUncovered() int {
	numUncovered := 0
	for i := 1; i <= b.SideLength; i++ {
		for j := 1; j <= b.SideLength; j++ {
			cell := b.Cell(i, j)
			if cell.Uncovered {
				numUncovered++
			}
		}
	}
	return numUncovered
}

// BombCells returns the bomb cells
func (b *Board) BombCells() []*Cell {
	bombCells := []*Cell{}
	for i := 1; i <= b.SideLength; i++ {
		for j := 1; j <= b.SideLength; j++ {
			cell := b.Cell(i, j)
			if cell.Bomb {
				bombCells = append(bombCells, cell)
			}
		}
	}
	return bombCells
}

// UncoverAll uncovers all the cells.
func (b *Board) UncoverAll() {
	for i := 1; i <= b.SideLength; i++ {
		for j := 1; j <= b.SideLength; j++ {
			cell := b.Cell(i, j)
			cell.Uncovered = true
		}
	}
}
