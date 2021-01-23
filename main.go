package main

import (
	"errors"
	"fmt"
)
type Game struct {
	board      [9]string
	player     string
	turnNumber int
	winner     string
	gameOver   bool
}

func main() {
	var game Game
	game.player = "O"
	game.gameOver = false
	for game.gameOver != true {
		PrintBoard(game.board)
		oneBasedMovePosition := getInputPlayPosition()
		err := game.play(oneBasedMovePosition)
		if err != nil {
			fmt.Println(err)
			continue
		}

		game.gameOver, game.winner = CheckForWinner(game.board, game.turnNumber)
	}
	//final part
	PrintBoard(game.board)
	if game.winner == "" {
		fmt.Println("Draw! ")
	} else {
		fmt.Printf("Barikallah %s is winner ", game.winner)
	}
}

func CheckForWinner(board [9]string, turnNumber int) (bool, string) {
	k := 0
	for k < 9 {
		if board[k] == board[k+1] && board[k+1] == board[k+2] && board[k] != "" {
			return true, board[k]
		} else {
			k += 3
		}
	}
	k = 0
	for k < 3 {
		if board[k] == board[k+3] && board[k+3] == board[k+6] && board[k] != "" {
			return true, board[k]
		} else {
			k += 1
		}
	}
	diagonal1 := board[2] == board[4] && board[4] == board[6] && board[2] != ""
	diagonal2 := board[0] == board[4] && board[4] == board[8] && board[0] != ""
	if diagonal1 || diagonal2 {
		return true, board[4]
	}
	if turnNumber == 9 {
		return true, ""
	}
	return false, ""
}

func (game *Game) SwitchPlayers() {
	if game.player == "O" {
		game.player = "X"
		return
	}
	game.player = "O"
}

func (game *Game) play(oneBasedPosition int) error {
	if game.board[oneBasedPosition-1] == "" {
		game.board[oneBasedPosition-1] = game.player
		game.SwitchPlayers()
		game.turnNumber += 1
		return nil
	}
	return errors.New("Try another position")
}

func getInputPlayPosition() int {
	var oneBasedMoveInt int
	fmt.Println("Enter One based Position to play: ")
	fmt.Scan(&oneBasedMoveInt)
	return oneBasedMoveInt
}

func PrintBoard(board [9]string) {
	for index, value := range board {
		if value == "" {
			fmt.Printf(" ")
		} else {
			fmt.Printf(value)
		}

		if index > 0 && (index+1)%3 == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf("|")

		}

	}
}
