package main
import (
  "log"
  "os/exec"
)

func execCommand() string {
	output, err := exec.Command("ls", "-l").CombinedOutput()
	if err != nil {
		log.Print(err)
	}
	return string(output)
}

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
	// WIP
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
func getMove() {
	// WIP
}
func makeMove() {
	// WIP
}
func prettyPrintBoard() {
	// WIP
}
