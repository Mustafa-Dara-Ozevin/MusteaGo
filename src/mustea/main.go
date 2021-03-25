package main

import (
	"github.com/Mustafa-Dara-Ozevin/MusteaGo.git/src/common"
)

func main() {
	common.AllInit()
	var board common.Board
	a := common.StartingFen
	board.ParseFen(a)
	board.PrintBoard()
}
