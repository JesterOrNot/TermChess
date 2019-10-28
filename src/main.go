package main
import (
	"log"
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
	var theBoard [][]string = newBoard()
	fmt.Print(theBoard)
}
func mainMenu() {
	// WIP
}
func newBoard() [][]string {
	var board = [][]string{{"BR","BN","BB","BQ","BK","BB","BN","BR"},
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
func prettyPrintBoard(theBoard string[8][8]) {
	fmt.Println("---------")
	for i:=1; i<=8; i++ {
		for j:=1; j<=8; j++ {
			fmt.Print("|")
		}
	}
}
