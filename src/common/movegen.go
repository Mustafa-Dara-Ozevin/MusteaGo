package common

var LoopSlidePc = [...]int{WB, WR, WQ, 0, BB, BR, BQ, 0}
var LoopSlideIndex = [...]int{0, 4}

var LoopNonSlidePc = [...]int{WN, WK, 0, BN, BK, 0}
var LoopNonSlideIndex = [...]int{0, 3}

var PceDir = [13][8]int{
	{0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0},
	{-8, -19, -21, -12, 8, 19, 21, 12},
	{-9, -11, 11, 9, 0, 0, 0, 0},
	{-1, -10, 1, 10, 0, 0, 0, 0},
	{-1, -10, 1, 10, -9, -11, 11, 9},
	{-1, -10, 1, 10, -9, -11, 11, 9},
	{0, 0, 0, 0, 0, 0, 0},
	{-8, -19, -21, -12, 8, 19, 21, 12},
	{-9, -11, 11, 9, 0, 0, 0, 0},
	{-1, -10, 1, 10, 0, 0, 0, 0},
	{-1, -10, 1, 10, -9, -11, 11, 9},
	{-1, -10, 1, 10, -9, -11, 11, 9},
}

var NumDir = [13]int{
	0, 0, 8, 4, 4, 8, 8, 0, 8, 4, 4, 8, 8,
}

func MOVE(f, t, ca, pro, fl int) int {
	return (f | t<<7) | (ca << 14) | (pro << 20) | fl
}
func isSqOffboard(sq int) bool {
	return FileBrd[sq] == OffBoard
}

func (b *Board) addQuietMove(move int, list *MoveList) {
	list.Moves[list.Count].Move = move
	list.Moves[list.Count].Score = 0
	list.Count++

}
func (b *Board) addCaptureMove(move int, list *MoveList) {
	list.Moves[list.Count].Move = move
	list.Moves[list.Count].Score = 0
	list.Count++

}
func (b *Board) addEnPassantMove(move int, list *MoveList) {
	list.Moves[list.Count].Move = move
	list.Moves[list.Count].Score = 0
	list.Count++

}

func (b *Board) addWhitePawnCapMove(from, to, cap int, list *MoveList) {

	if RankBrd[from] == Rank7 {
		b.addQuietMove(MOVE(from, to, cap, WQ, 0), list)
		b.addQuietMove(MOVE(from, to, cap, WR, 0), list)
		b.addQuietMove(MOVE(from, to, cap, WB, 0), list)
		b.addQuietMove(MOVE(from, to, cap, WN, 0), list)
	} else {
		b.addQuietMove(MOVE(from, to, cap, Empty, 0), list)
	}
}

func (b *Board) addWhitePawnMove(from, to int, list *MoveList) {

	if RankBrd[from] == Rank7 {
		b.addQuietMove(MOVE(from, to, Empty, WQ, 0), list)
		b.addQuietMove(MOVE(from, to, Empty, WR, 0), list)
		b.addQuietMove(MOVE(from, to, Empty, WB, 0), list)
		b.addQuietMove(MOVE(from, to, Empty, WN, 0), list)
	} else {
		b.addQuietMove(MOVE(from, to, Empty, Empty, 0), list)
	}
}

func (b *Board) addBlackPawnCapMove(from, to, cap int, list *MoveList) {

	if RankBrd[from] == Rank2 {
		b.addQuietMove(MOVE(from, to, cap, BQ, 0), list)
		b.addQuietMove(MOVE(from, to, cap, BR, 0), list)
		b.addQuietMove(MOVE(from, to, cap, BB, 0), list)
		b.addQuietMove(MOVE(from, to, cap, BN, 0), list)
	} else {
		b.addQuietMove(MOVE(from, to, cap, Empty, 0), list)
	}
}

func (b *Board) addBlackPawnMove(from, to int, list *MoveList) {

	if RankBrd[from] == Rank2 {
		b.addQuietMove(MOVE(from, to, Empty, BQ, 0), list)
		b.addQuietMove(MOVE(from, to, Empty, BR, 0), list)
		b.addQuietMove(MOVE(from, to, Empty, BB, 0), list)
		b.addQuietMove(MOVE(from, to, Empty, BN, 0), list)
	} else {
		b.addQuietMove(MOVE(from, to, Empty, Empty, 0), list)
	}
}

func (b *Board) GenerateAllMoves(list *MoveList) {

	list.Count = 0

	side := b.Side

	if side == White {

		for pceNum := 0; pceNum < b.PceNum[WP]; pceNum++ {
			sq := b.PceList[WP][pceNum]

			if b.Pieces[sq+10] == Empty {
				b.addWhitePawnMove(sq, sq+10, list)
				if RankBrd[sq] == Rank2 && b.Pieces[sq+20] == Empty {
					b.addQuietMove(MOVE(sq, (sq+20), Empty, Empty, MFlagPS), list)
				}
			}

			if !isSqOffboard(sq+9) && PieceCol[b.Pieces[sq+9]] == Black {
				b.addWhitePawnCapMove(sq, sq+9, b.Pieces[sq+9], list)
			}
			if !isSqOffboard(sq+11) && PieceCol[b.Pieces[sq+11]] == Black {
				b.addWhitePawnCapMove(sq, sq+11, b.Pieces[sq+11], list)
			}

			if b.EnPas != NoSq {
				if sq+9 == b.EnPas {
					b.addEnPassantMove(MOVE(sq, sq+9, Empty, Empty, MFlagEP), list)
				}
				if sq+11 == b.EnPas {
					b.addEnPassantMove(MOVE(sq, sq+11, Empty, Empty, MFlagEP), list)
				}
			}
		}
		if b.CastlePerm&WKCA != 0 {
			if b.Pieces[F1] == Empty && b.Pieces[G1] == Empty {
				if !b.IsAttacked(E1, Black) && !b.IsAttacked(G1, Black) {
					b.addQuietMove(MOVE(E1, G1, Empty, Empty, MFlagCa), list)
				}
			}
		}
		if b.CastlePerm&WQCA != 0 {
			if b.Pieces[B1] == Empty && b.Pieces[D1] == Empty && b.Pieces[C1] == Empty {
				if !b.IsAttacked(E1, Black) && !b.IsAttacked(D1, Black) && !b.IsAttacked(C1, Black) {
					b.addQuietMove(MOVE(E1, C1, Empty, Empty, MFlagCa), list)

				}
			}
		}

	} else {
		for pceNum := 0; pceNum < b.PceNum[BP]; pceNum++ {
			sq := b.PceList[BP][pceNum]

			if b.Pieces[sq-10] == Empty {
				b.addBlackPawnMove(sq, sq-10, list)
				if RankBrd[sq] == Rank7 && b.Pieces[sq-20] == Empty {
					b.addQuietMove(MOVE(sq, (sq-20), Empty, Empty, MFlagPS), list)
				}
			}

			if !isSqOffboard(sq-9) && PieceCol[b.Pieces[sq-9]] == White {
				b.addBlackPawnCapMove(sq, sq-9, b.Pieces[sq-9], list)
			}
			if !isSqOffboard(sq-11) && PieceCol[b.Pieces[sq-11]] == White {
				b.addBlackPawnCapMove(sq, sq-11, b.Pieces[sq-11], list)
			}

			if b.EnPas != NoSq {
				if sq-9 == b.EnPas {
					b.addEnPassantMove(MOVE(sq, sq-9, Empty, Empty, MFlagEP), list)
				}
				if sq-11 == b.EnPas {
					b.addEnPassantMove(MOVE(sq, sq-11, Empty, Empty, MFlagEP), list)
				}
			}
		}
		if b.CastlePerm&BKCA != 0 {
			if b.Pieces[F8] == Empty && b.Pieces[G8] == Empty {
				if !b.IsAttacked(E8, White) && !b.IsAttacked(G8, White) {
					b.addQuietMove(MOVE(E8, G8, Empty, Empty, MFlagCa), list)

				}
			}
		}
		if b.CastlePerm&BQCA != 0 {
			if b.Pieces[B8] == Empty && b.Pieces[D8] == Empty && b.Pieces[C8] == Empty {
				if !b.IsAttacked(E8, White) && !b.IsAttacked(D8, White) && !b.IsAttacked(C8, White) {
					b.addQuietMove(MOVE(E8, C8, Empty, Empty, MFlagCa), list)

				}
			}
		}
	}

	// Loop for Slider Pieces
	pceIndex := LoopSlideIndex[side]
	pce := LoopSlidePc[pceIndex]
	for pce != 0 {
		pce = LoopSlidePc[pceIndex]
		for pceNum := 0; pceNum < b.PceNum[pce]; pceNum++ {
			sq := b.PceList[pce][pceNum]
			for index := 0; index < NumDir[pce]; index++ {
				dir := PceDir[pce][index]
				t_sq := dir + sq

				for !isSqOffboard(t_sq) {

					if b.Pieces[t_sq] != Empty {
						if PieceCol[b.Pieces[t_sq]] == side^1 {
							b.addCaptureMove(MOVE(sq, t_sq, b.Pieces[t_sq], Empty, 0), list)
						}
						break
					}
					b.addQuietMove(MOVE(sq, t_sq, Empty, Empty, 0), list)
					t_sq += dir
				}
			}
		}
		pce = LoopSlidePc[pceIndex]
		pceIndex++
	}

	// Loop for Non Slider Pieces
	pceIndex = LoopNonSlideIndex[side]
	pce = LoopNonSlidePc[pceIndex]
	for pce != 0 {
		pce = LoopNonSlidePc[pceIndex]
		for pceNum := 0; pceNum < b.PceNum[pce]; pceNum++ {
			sq := b.PceList[pce][pceNum]
			for index := 0; index < NumDir[pce]; index++ {
				dir := PceDir[pce][index]
				t_sq := dir + sq

				if isSqOffboard(t_sq) {
					continue
				}
				if b.Pieces[t_sq] != Empty {
					if PieceCol[b.Pieces[t_sq]] == side^1 {
						b.addCaptureMove(MOVE(sq, t_sq, b.Pieces[t_sq], Empty, 0), list)
					}
					continue
				}
				b.addQuietMove(MOVE(sq, t_sq, Empty, Empty, 0), list)

			}
		}
		pce = LoopNonSlidePc[pceIndex]
		pceIndex++
	}

}
