package common

import "fmt"

func PrSq(sq int) string {

	file := FileBrd[sq]
	rank := RankBrd[sq]

	return fmt.Sprintf("%c%c", ('a' + file), ('1' + rank))

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

	for index := 0; index < list.count; index++ {

		move := list.moves[index].move
		score := list.moves[index].score

		fmt.Printf("Move:%d > %s (score:%d)\n", index+1, PrMove(move), score)
	}
	fmt.Printf("MoveList Total %d Moves:\n\n", list.count)
}
