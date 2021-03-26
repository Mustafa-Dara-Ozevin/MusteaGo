package common

import (
	"fmt"
)

var leafnodes int

func perft(b *Board, depth int) {
	if depth == 0 {
		leafnodes++
	} else {
		var buffer MoveList
		b.GenerateAllMoves(&buffer)
		for _, move := range buffer.GetMoves() {
			if b.MakeMove(move) {
				if depth > 1 {
					perft(b, depth-1)
				} else {
					leafnodes++
				}
				b.TakeMove()
			}
		}
	}
}
func PerftTest(b *Board, depth int) {
	var buffer MoveList
	startT := TimeMs()
	b.PrintBoard()
	leafnodes = 0
	b.GenerateAllMoves(&buffer)
	for idx := range buffer.GetMoves() {
		move := buffer.GetMoves()[idx]
		if b.MakeMove(move) {
			cumnodes := leafnodes
			perft(b, depth-1)
			b.TakeMove()
			longnodes := leafnodes - cumnodes
			fmt.Printf("move %d : %s : %1d\n", idx+1, PrMove(move), longnodes)
		}
	}
	endT := TimeMs()
	fmt.Printf("Test complete: %1d notes visited in %d ms\n", leafnodes, endT-startT)
}
