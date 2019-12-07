package main

import (
	"strings"
	"github.com/gookit/color"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	var theBoard [8][8]string = newBoard()
	for true {
		prettyPrintBoard(theBoard)
		theBoard = makeMove(theBoard)
	}
}
func mainMenu() {
	// WIP
}
func newBoard() [8][8]string {
	var board = [8][8]string{{"BR", "BN", "BB", "BQ", "BK", "BB", "BN", "BR"},
		{"BP", "BP", "BP", "BP", "BP", "BP", "BP", "BP"},
		{"XX", "XX", "XX", "XX", "XX", "XX", "XX", "XX"},
		{"XX", "XX", "XX", "XX", "XX", "XX", "XX", "XX"},
		{"XX", "XX", "XX", "XX", "XX", "XX", "XX", "XX"},
		{"XX", "XX", "XX", "XX", "XX", "XX", "XX", "XX"},
		{"WP", "WP", "WP", "WP", "WP", "WP", "WP", "WP"},
		{"WR", "WN", "WB", "WQ", "WK", "WB", "WN", "WR"}}
	return board
}
func executeCommand(theCommand string) {
	cmd := exec.Command(theCommand)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Print(err)
	}
}
func getMove() [4]int {
	var currentX, currentY, targetX, targetY int
	fmt.Print("What is the current X position?: ")
	fmt.Scan(&currentX)
	fmt.Print("What is the current Y position?: ")
	fmt.Scan(&currentY)
	fmt.Print("What is the target X position?: ")
	fmt.Scan(&targetX)
	fmt.Print("What is the target Y position?: ")
	fmt.Scan(&targetY)
	var theArray = [4]int{currentY - 1, currentX - 1, targetY - 1, targetX - 1}
	return theArray
}
func makeMove(board [8][8]string) [8][8]string {
	var data = getMove()
	var target = board[data[2]][data[3]]
	var temp = board[data[0]][data[1]]
	if target == "XX" && temp == "WP" {
		if data[2] == data[0]-1 && data[3] == data[1] {
			board[data[2]][data[3]] = temp
			board[data[0]][data[1]] = "XX"
		} else {
			fmt.Println("Invalid move")
		}
	} else if target == "XX" && temp == "BP" {
		if data[2] == data[0]+1 && data[3] == data[1] {
			board[data[2]][data[3]] = temp
			board[data[0]][data[1]] = "XX"
		}
	} else if target == "XX" && (temp == "BK" || temp == "WK") {
		if (data[2] == data[0]+1) || (data[2] == data[0]+1 && data[3] == data[1]+1) || (data[3] == data[1]+1) || (data[3] == data[1]-1) || (data[2] == data[0]-1) || (data[2] == data[0]-1 && data[3] == data[1]+1) || (data[2] == data[0]+1 && data[3] == data[1]-1) || (data[2] == data[0]-1 && data[3] == data[1]-1) {
			board[data[2]][data[3]] = temp
			board[data[0]][data[1]] = "XX"
		} else {
			fmt.Println("Invalid move")
		}
	} else if (target == "XX" || strings.HasPrefix(target,"B")) && (temp == "BN" || temp == "WN") {
		if (data[2] == data[0]+2 && data[3] == data[1]+1) || (data[2] == data[0]+2 && data[3] == data[1]-1) || (data[2] == data[0]-2 && data[3] == data[1]-1) || (data[2] == data[0]-2 && data[3] == data[1]+1) || (data[2] == data[0]+1 && data[3] == data[1]-2) || (data[2] == data[0]-1 && data[3] == data[1]-2) || (data[2] == data[0]+1 && data[3] == data[1]+2) || (data[2] == data[0]-1 && data[3] == data[1]+2) {
			board[data[2]][data[3]] = temp
			board[data[0]][data[1]] = "XX"
		} else {
			fmt.Println("Invalid move")
		}
	} else {
		fmt.Println("Invalid move")
	}
	return board
}
func prettyPrintBoard(theBoard [8][8]string) {
	executeCommand("clear")
	fmt.Println("   1  2  3  4  5  6  7  8")
	fmt.Println("  ________________________")
	var count = 1
	for i := 0; i <= 7; i++ {
		fmt.Print(count, "|")
		count++
		for j := 0; j <= 7; j++ {
			if theBoard[i][j] == "XX" {
				fmt.Print(" x ")
			} else if theBoard[i][j] == "BR" {
				color.FgRed.Print(" R ")
			} else if theBoard[i][j] == "BN" {
				color.FgRed.Print(" N ")
			} else if theBoard[i][j] == "BB" {
				color.FgRed.Print(" B ")
			} else if theBoard[i][j] == "BK" {
				color.FgRed.Print(" K ")
			} else if theBoard[i][j] == "BQ" {
				color.FgRed.Print(" Q ")
			} else if theBoard[i][j] == "BP" {
				color.FgRed.Print(" P ")
			} else if theBoard[i][j] == "WP" {
				color.FgCyan.Print(" P ")
			} else if theBoard[i][j] == "WR" {
				color.FgCyan.Print(" R ")
			} else if theBoard[i][j] == "WN" {
				color.FgCyan.Print(" N ")
			} else if theBoard[i][j] == "WB" {
				color.FgCyan.Print(" B ")
			} else if theBoard[i][j] == "WK" {
				color.FgCyan.Print(" K ")
			} else if theBoard[i][j] == "WQ" {
				color.FgCyan.Print(" Q ")
			}
			if j == 7 {
				fmt.Print("|")
				fmt.Println()
			}
		}
	}
}
