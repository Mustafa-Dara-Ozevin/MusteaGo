package main

import (
	"github.com/Mustafa-Dara-Ozevin/MusteaGo.git/src/common"
)

func main() {
	common.AllInit()
	var board common.Board
	a := "r1bq2k1/ppp1n1pp/2nb4/3pp3/8/PP1PPrPP/1BP1NP2/RN1QK2R w KQ - 0 11"
	board.ParseFen(a)
	board.PrintBoard()

}
