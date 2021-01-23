package main

import (
	"bufio"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// A minesweeper game. Execute "go build" to build the executable.

func main() {
	println("MINESWEEPER")
	println("Enter help for help.")
	println("")
	rand.Seed(time.Now().UTC().UnixNano())
	game := NewGame(8, 3)
	coordinatePattern := regexp.MustCompile(`^(\d+)\s*,\s*(\d+)$`)
	flagPattern := regexp.MustCompile(`^f\s*(\d+)\s*,\s*(\d+)$`)
	for true {
		game.Board.Print()
		println("")
		if game.Won() {
			println("YOU WON! :-)")
			break
		}
		if game.Lost() {
			println("You lost :-(")
			game.Board.UncoverAll()
			game.Board.Print()
			break
		}
		print("> ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.ToLower(strings.TrimSpace(text))
		if text == "help" {
			println("Enter 3,5 to uncover row 3, column 5.")
			println("Enter F3,5 to flag/unflag row 3, column 5.")
			println("Enter quit to stop.")
		} else if text == "quit" {
			println("Goodbye!")
			break
		} else if matches := coordinatePattern.FindStringSubmatch(text); matches != nil {
			row, _ := strconv.Atoi(matches[1])
			col, _ := strconv.Atoi(matches[2])
			err := game.Board.Uncover(row, col)
			if err != nil {
				println(err.Error())
			}
		} else if matches := flagPattern.FindStringSubmatch(text); matches != nil {
			row, _ := strconv.Atoi(matches[1])
			col, _ := strconv.Atoi(matches[2])
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
