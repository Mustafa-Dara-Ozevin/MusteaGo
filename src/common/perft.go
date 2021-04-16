package common

import (
	"fmt"
)

var leafnodes int
var castles int
var captures int
var enpas int
var promotions int

func perft(b *Board, depth int) {
	if depth == 0 {
		leafnodes++
	} else {
		var buffer MoveList
		b.GenerateAllMoves(&buffer)
		for _, move := range buffer.GetMoves() {
			if IsCa(move.Move) {
				castles++
			}
			if IsCp(move.Move) {
				captures++
			}
			if IsPr(move.Move) {
				promotions++
			}
			if IsEnPas(move.Move) {
				enpas++
			}
			if b.MakeMove(move.Move) {
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
	castles = 0
	promotions = 0
	enpas = 0
	captures = 0
	b.GenerateAllMoves(&buffer)
	for idx := range buffer.GetMoves() {
		move := buffer.GetMoves()[idx]
		if b.MakeMove(move.Move) {
			if IsCa(move.Move) {
				castles++
			}
			if IsCp(move.Move) {
				captures++
			}
			if IsPr(move.Move) {
				promotions++
			}
			if IsEnPas(move.Move) {
				enpas++
			}
			cumnodes := leafnodes
			perft(b, depth-1)
			b.TakeMove()
			longnodes := leafnodes - cumnodes
			fmt.Printf("move %d : %s : %1d\n", idx+1, PrMove(move.Move), longnodes)
		}
	}
	endT := TimeMs()
	fmt.Printf("Test complete: %1d castles : %1d captures : %1d Enpas : %1d promotions : %1d notes visited in %d ms\n", leafnodes, castles, captures, enpas, promotions, endT-startT)
}
