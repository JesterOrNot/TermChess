package main
import (
	"log"
	"os"
	"os/exec"
	"fmt"
)

type King struct {
	// WIP
}
type Queen struct {
	// WIP
}
type Pawn struct {
	// WIP
}
type Rook struct {
	// WIP
}
type Knight struct {
	// WIP
}
type Bishop struct {
	// WIP
}

func main() {
	var theBoard [8][8]string = newBoard()
	prettyPrintBoard(theBoard)
}
func mainMenu() {
	// WIP
}
func newBoard() [8][8]string {
	var board = [8][8]string{{"BR","BN","BB","BQ","BK","BB","BN","BR"},
						   {"BP","BP","BP","BP","BP","BP","BP","BP"},
						   {"XX","XX","XX","XX","XX","XX","XX","XX"},
						   {"XX","XX","XX","XX","XX","XX","XX","XX"},
						   {"XX","XX","XX","XX","XX","XX","XX","XX"},
						   {"XX","XX","XX","XX","XX","XX","XX","XX"},
						   {"WP","WP","WP","WP","WP","WP","WP","WP"},
						   {"WR","WN","WB","WQ","WK","WB","WN","WR"}}
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
func getMove() {
	// WIP
}
func makeMove() {
	// WIP
}
func prettyPrintBoard(theBoard [8][8]string) {
	fmt.Println("---------")
	for i:=0; i<=7; i++ {
		for j:=0; j<=7; j++ {
			if theBoard[i][j] == "XX" {
				fmt.Print(" x ")
			} else if theBoard[i][j] == "BR" {
				executeCommand("./src/lib/bin/changeToRed")
				fmt.Print(" R ")
				executeCommand("./src/lib/bin/resetColor")
			} else if theBoard[i][j] == "BN" {
				executeCommand("./src/lib/bin/changeToRed")
				fmt.Print(" N ")
				executeCommand("./src/lib/bin/resetColor")
			} else if theBoard[i][j] == "BB" {
				executeCommand("./src/lib/bin/changeToRed")
				fmt.Print(" B ")
				executeCommand("./src/lib/bin/resetColor")
			} else if theBoard[i][j] == "BK" {
				executeCommand("./src/lib/bin/changeToRed")
				fmt.Print(" K ")
				executeCommand("./src/lib/bin/resetColor")
			} else if theBoard[i][j] == "BQ" {
				executeCommand("./src/lib/bin/changeToRed")
				fmt.Print(" Q ")
				executeCommand("./src/lib/bin/resetColor")
			} else if theBoard[i][j] == "BP" {
				executeCommand("./src/lib/bin/changeToRed")
				fmt.Print(" P ")
				executeCommand("./src/lib/bin/resetColor")
			} else if theBoard[i][j] == "WP" {
				executeCommand("./src/lib/bin/changeToBlue")
				fmt.Print(" P ")
				executeCommand("./src/lib/bin/resetColor")
			} else if theBoard[i][j] == "WR" {
				executeCommand("./src/lib/bin/changeToBlue")
				fmt.Print(" R ")
				executeCommand("./src/lib/bin/resetColor")
			} else if theBoard[i][j] == "WN" {
				executeCommand("./src/lib/bin/changeToBlue")
				fmt.Print(" N ")
				executeCommand("./src/lib/bin/resetColor")
			} else if theBoard[i][j] == "WB" {
				executeCommand("./src/lib/bin/changeToBlue")
				fmt.Print(" B ")
				executeCommand("./src/lib/bin/resetColor")
			} else if theBoard[i][j] == "WK" {
				executeCommand("./src/lib/bin/changeToBlue")
				fmt.Print(" K ")
				executeCommand("./src/lib/bin/resetColor")
			} else if theBoard[i][j] == "WQ" {
				executeCommand("./src/lib/bin/changeToBlue")
				fmt.Print(" Q ")
				executeCommand("./src/lib/bin/resetColor")
			}
			if j==7 {
				fmt.Println()
			}
		}
	}
}
