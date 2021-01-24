package main

import "fmt"

// maxSideLength is the max for sideLength. It is 26 to allow the letters a-z.
const maxSideLength = 26

// Board is a Minesweeper board
type Board struct {
	SideLength int
	Cells      *[][]*Cell
}

// NewBoard returns an initialized Minesweeper board.
func NewBoard(sideLength int) *Board {
	if sideLength > maxSideLength {
		panic(fmt.Sprintf("sideLength must not exceed %d", maxSideLength))
	}
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
		edge += ToColLetter(i + 1)
	}
	println(edge)
	for i, row := range *b.Cells {
		line := ToRowLetter(i + 1)
		for _, cell := range row {
			line += cell.Symbol()
		}
		line += ToRowLetter(i + 1)
		println(line)
	}
	println(edge)
}

// OutOfBounds returns whether the coordinates are out of bounds.
func (b *Board) OutOfBounds(row, col int) bool {
	return row < 1 || col < 1 || row > b.SideLength || col > b.SideLength
}

// ToRowLetter converts 1-26 to A-Z.
func ToRowLetter(row int) string {
	return string(int('A') + row - 1)
}

// ToColLetter converts 1-26 to a-z.
func ToColLetter(col int) string {
	return string(int('a') + col - 1)
}

// ToRowNumber converts A-Z to 1-26.
func ToRowNumber(row string) int {
	return int(row[0]) - int('A') + 1
}

// ToColNumber converts a-z to 1-26.
func ToColNumber(row string) int {
	return int(row[0]) - int('a') + 1
}

// Uncover exposes the given cell.
func (b *Board) Uncover(row, col int) error {
	if b.OutOfBounds(row, col) {
		return fmt.Errorf("%s%s is out of bounds", ToRowLetter(row), ToColLetter(col))
	}
	cell := b.Cell(row, col)
	if cell.Flagged {
		return fmt.Errorf("%s%s is flagged and can't be uncovered", ToRowLetter(row), ToColLetter(col))
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
		return fmt.Errorf("%s%s is out of bounds", ToRowLetter(row), ToColLetter(col))
	}
	cell := b.Cell(row, col)
	if cell.Uncovered {
		return fmt.Errorf("%s%s is already uncovered", ToRowLetter(row), ToColLetter(col))
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
			cell.Flagged = false
		}
	}
}
