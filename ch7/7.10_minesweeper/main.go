package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

// A minesweeper game. Execute "go build" to build the executable.

func main() {
	seed := time.Now().UTC().UnixNano()
	println("MINESWEEPER")
	println(fmt.Sprintf("Enter help for help. Seed = %d.", seed))
	println("")
	rand.Seed(seed)
	game := NewGame(8, 3)
	coordinatePattern := regexp.MustCompile(`^([A-Z])\s*([a-z])$`)
	flagPattern := regexp.MustCompile(`^f\s*([A-Z])\s*([a-z])$`)
	for true {
		game.Board.Print()
		println("")
		if game.Won() {
			println("YOU WON! :-)")
			break
		}
		if game.Lost() {
			println("You lost :-(")
			println("")
			game.Board.UncoverAll()
			game.Board.Print()
			break
		}
		print("> ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "help" {
			println("Enter Ce to uncover row C, column e.")
			println("Enter fCe to flag/unflag row C, column e.")
			println("Enter quit to stop.")
		} else if text == "quit" {
			println("Goodbye!")
			break
		} else if matches := coordinatePattern.FindStringSubmatch(text); matches != nil {
			row := ToRowNumber(matches[1])
			col := ToColNumber(matches[2])
			err := game.Board.Uncover(row, col)
			if err != nil {
				println(err.Error())
			}
		} else if matches := flagPattern.FindStringSubmatch(text); matches != nil {
			row := ToRowNumber(matches[1])
			col := ToColNumber(matches[2])
			err := game.Board.Flag(row, col)
			if err != nil {
				println(err.Error())
			}
		} else {
			println("I didn't understand that. Enter help for help.")
		}
		println("")
	}
}
