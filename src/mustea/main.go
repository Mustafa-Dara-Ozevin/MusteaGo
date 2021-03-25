package main

import (
	"github.com/Mustafa-Dara-Ozevin/MusteaGo.git/src/common"
)

func main() {
	common.AllInit()
	var board common.Board
	a := "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1"
	board.ParseFen(a)
	board.PrintBoard()
	var list common.MoveList
	board.GenerateAllMoves(&list)
	common.PrintMoveList(&list)

}
