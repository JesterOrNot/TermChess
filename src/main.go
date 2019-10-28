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
			}
			if j==7 {
				fmt.Println()
			}
		}
	}
}
