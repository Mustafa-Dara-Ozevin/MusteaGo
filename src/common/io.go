package common

import "fmt"

func PrSq(sq int) string {

	file := FileBrd[sq]
	rank := RankBrd[sq]

	return fmt.Sprintf("%c%c", ('a' + file), ('1' + rank))

}

func (b *Board) ParseMove(fen string) int {
	from := Fr2Sq(int(byte(fen[0])-'a'), int(byte(fen[1])-'1'))
	to := Fr2Sq(int(byte(fen[2])-'a'), int(byte(fen[3])-'1'))

	var list MoveList
	b.GenerateAllMoves(&list)
	PromPce := Empty

	for MoveNum := 0; MoveNum < list.Count; MoveNum++ {
		Move := list.Moves[MoveNum].Move
		if FromSq(Move) == from && ToSq(Move) == to {
			PromPce = Promoted(Move)
			if PromPce != Empty {
				if IsRq(PromPce) && !IsBq(PromPce) && fen[4] == 'r' {
					return Move
				} else if !IsRq(PromPce) && IsBq(PromPce) && fen[4] == 'b' {
					return Move
				} else if IsRq(PromPce) && IsBq(PromPce) && fen[4] == 'q' {
					return Move
				} else if IsKn(PromPce) && fen[4] == 'n' {
					return Move
				}
				continue
			}
			return Move
		}
	}

	return NoMove
}

func PrMove(move int) string {

	var MvStr string

	ff := FileBrd[FromSq(move)]
	rf := RankBrd[FromSq(move)]
	ft := FileBrd[ToSq(move)]
	rt := RankBrd[ToSq(move)]

	promoted := Promoted(move)

	if promoted != 0 {
		pchar := 'q'
		if IsKn(promoted) {
			pchar = 'n'
		} else if IsRq(promoted) && !IsBq(promoted) {
			pchar = 'r'
		} else if !IsRq(promoted) && IsBq(promoted) {
			pchar = 'b'
		}
		MvStr = fmt.Sprintf("%c%c%c%c%c", ('a' + ff), ('1' + rf), ('a' + ft), ('1' + rt), pchar)
	} else {
		MvStr = fmt.Sprintf("%c%c%c%c", ('a' + ff), ('1' + rf), ('a' + ft), ('1' + rt))
	}

	return MvStr
}

func PrintMoveList(list *MoveList) {
	fmt.Printf("MoveList:\n")

	for index := 0; index < list.Count; index++ {

		move := list.Moves[index].Move
		score := list.Moves[index].Score

		fmt.Printf("Move:%d > %s (score:%d)\n", index+1, PrMove(move), score)
	}
	fmt.Printf("MoveList Total %d Moves:\n\n", list.Count)
}
