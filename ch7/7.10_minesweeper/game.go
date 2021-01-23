package main

import "math/rand"

// Game is a Minesweeper game.
type Game struct {
	Board     *Board
	BombCount int
}

// NewGame returns an initialized Minesweeper game.
func NewGame(sideLength, bombCount int) *Game {
	game := &Game{
		Board:     NewBoard(sideLength),
		BombCount: bombCount,
	}
	for i := 0; i < bombCount; i++ {
		for true {
			row := rand.Intn(sideLength) + 1
			col := rand.Intn(sideLength) + 1
			cell := game.Board.Cell(row, col)
			if !cell.Bomb {
				cell.Bomb = true
				for _, neighbour := range cell.Neighbours() {
					neighbour.AdjacentBombs++
				}
				break
			}
		}
	}
	return game
}

// Won returns true if you won the game.
func (g *Game) Won() bool {
	return g.Board.NumUncovered() == g.Board.NumCells()-g.BombCount && !g.Lost()
}

// Lost returns true if you lost the game.
func (g *Game) Lost() bool {
	for _, bombCell := range g.Board.BombCells() {
		if bombCell.Uncovered {
			return true
		}
	}
	return false
}
