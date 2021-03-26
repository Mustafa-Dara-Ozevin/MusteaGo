package main

import (
	"github.com/Mustafa-Dara-Ozevin/MusteaGo.git/src/common"
)

func main() {
	common.AllInit()

	a := "r3k2r/p1ppqpb1/bn2pnN1/3P4/1p2P3/2N2Q1p/PPPBBPPP/R3K2R b KQkq - 0 1"
	board := common.NewBoardFromFEN(a)
	common.PerftTest(&board, 3)

}
