package common

var CastlePerm = [120]int{
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 13, 15, 15, 15, 12, 15, 15, 14, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 7, 15, 15, 15, 3, 15, 15, 11, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
}

func (b *Board) hashPce(pce, sq int) {
	b.PosKey ^= PieceKeys[pce][sq]
}
func (b *Board) hashCas() {
	b.PosKey ^= CastleKeys[b.CastlePerm]
}
func (b *Board) hashSide() {
	b.PosKey ^= SideKey
}
func (b *Board) hashEP() {
	b.PosKey ^= PieceKeys[Empty][b.EnPas]
}

func (b *Board) clearPiece(sq int) {
	pce := b.Pieces[sq]

	col := PieceCol[pce]
	t_pceNum := -1

	b.hashPce(pce, sq)

	b.Pieces[sq] = Empty
	b.Material[col] -= PieceVal[pce]

	if PieceBig[pce] {
		b.BigPce[col]--
		if PieceMaj[pce] {
			b.MajPce[col]--
		} else {
			b.MinPce[col]--
		}
	} else {
		ClrBit(&b.Pawns[col], Sq64(sq))
		ClrBit(&b.Pawns[Both], Sq64(sq))
	}

	for index := 0; index < b.PceNum[pce]; index++ {
		if b.PceList[pce][index] == sq {
			t_pceNum = index
			break
		}
	}
	b.PceNum[pce]--

	b.PceList[pce][t_pceNum] = b.PceList[pce][b.PceNum[pce]]

}

func (b *Board) addPiece(sq, pce int) {

	var col = PieceCol[pce]

	b.hashPce(pce, sq)

	b.Pieces[sq] = pce

	if PieceBig[pce] {
		b.BigPce[col]++
		if PieceMaj[pce] {
			b.MajPce[col]++
		} else {
			b.MinPce[col]++
		}
	} else {
		SetBit(&b.Pawns[col], Sq64(sq))
		SetBit(&b.Pawns[Both], Sq64(sq))
	}

	b.Material[col] += PieceVal[pce]
	b.PceList[pce][b.PceNum[pce]] = sq
	b.PceNum[pce]++

}

func (b *Board) movePiece(from, to int) {

	pce := b.Pieces[from]
	col := PieceCol[pce]

	b.hashPce(pce, from)
	b.Pieces[from] = Empty

	b.hashPce(pce, to)
	b.Pieces[to] = pce

	if !PieceBig[pce] {
		ClrBit(&b.Pawns[col], Sq64(from))
		ClrBit(&b.Pawns[Both], Sq64(from))
		SetBit(&b.Pawns[col], Sq64(to))
		SetBit(&b.Pawns[Both], Sq64(to))
	}

	for index := 0; index < b.PceNum[pce]; index++ {
		if b.PceList[pce][index] == from {
			b.PceList[pce][index] = to
			break
		}
	}
}

func (b *Board) MakeMove(move int) bool {

	from := FromSq(move)
	to := ToSq(move)
	side := b.Side

	b.History[b.HisPly].PosKey = b.PosKey

	if IsEnPas(move) {
		if side == White {
			b.clearPiece(to - 10)
		} else {
			b.clearPiece(to + 10)
		}
	} else if IsCa(move) {
		switch to {
		case C1:
			b.movePiece(A1, D1)
		case C8:
			b.movePiece(A8, D8)
		case G1:
			b.movePiece(H1, F1)
		case G8:
			b.movePiece(H8, F8)
		default:
			break
		}
	}

	if b.EnPas != NoSq {
		b.hashEP()
	}
	b.hashCas()

	b.History[b.HisPly].move = move
	b.History[b.HisPly].FiftyMove = b.FiftyMove
	b.History[b.HisPly].EnPas = b.EnPas
	b.History[b.HisPly].CastlePerm = b.CastlePerm

	b.CastlePerm &= CastlePerm[from]
	b.CastlePerm &= CastlePerm[to]
	b.EnPas = NoSq

	b.hashCas()

	captured := Captured(move)
	b.FiftyMove++

	if captured != Empty {
		b.clearPiece(to)
		b.FiftyMove = 0
	}
	b.Ply++
	b.HisPly++

	if PiecePawn[b.Pieces[from]] {
		b.FiftyMove = 0
		if IsPaSt(move) {
			if side == White {
				b.EnPas = from + 10
			} else {
				b.EnPas = from - 10
			}
			{
				b.hashEP()
			}
		}
	}

	b.movePiece(from, to)

	prPce := Promoted(move)
	if IsPr(move) && prPce != Empty && !PiecePawn[prPce] {
		b.clearPiece(to)
		b.addPiece(to, prPce)
	}

	if PieceKing[b.Pieces[to]] {
		b.KingSq[b.Side] = to
	}

	b.Side ^= 1
	b.hashSide()

	if b.IsAttacked(b.KingSq[side], b.Side) {
		b.TakeMove()
		return false
	}

	return true

}

func (b *Board) TakeMove() {
	b.Ply--
	b.HisPly--
	move := b.History[b.HisPly].move
	from := FromSq(move)
	to := ToSq(move)

	if b.EnPas != NoSq {
		b.hashEP()
	}
	b.hashCas()

	b.CastlePerm = b.History[b.HisPly].CastlePerm
	b.FiftyMove = b.History[b.HisPly].FiftyMove
	b.EnPas = b.History[b.HisPly].EnPas

	if b.EnPas != NoSq {
		b.hashEP()
	}
	b.hashCas()

	b.Side ^= 1
	b.hashSide()

	if IsEnPas(move) {
		if b.Side == White {
			b.addPiece(to-10, BP)
		} else {
			b.addPiece(to+10, WP)
		}
	} else if IsCa(move) {
		switch to {
		case C1:
			b.movePiece(D1, A1)
		case C8:
			b.movePiece(D8, A8)
		case G1:
			b.movePiece(F1, H1)
		case G8:
			b.movePiece(F8, H8)
		default:
			break
		}
	}

	b.movePiece(to, from)

	if PieceKing[b.Pieces[from]] {
		b.KingSq[b.Side] = from
	}

	captured := Captured(move)
	if captured != Empty {
		b.addPiece(to, captured)
	}

	if IsPr(move) {
		b.clearPiece(from)
		var x int
		if PieceCol[Promoted(move)] == White {
			x = WP
		} else {
			x = BP
		}
		b.addPiece(from, x)
	}

}

func (b *Board) MakeNullMove() {

	b.History[b.HisPly].PosKey = b.PosKey

	if b.EnPas != NoSq {
		b.hashEP()
	}

	b.History[b.HisPly].move = NoMove
	b.History[b.HisPly].FiftyMove = b.FiftyMove
	b.History[b.HisPly].EnPas = b.EnPas
	b.History[b.HisPly].CastlePerm = b.CastlePerm
	b.EnPas = NoSq

	b.Side ^= 1
	b.Ply++
	b.HisPly++
	b.hashSide()

} // MakeNullMove

func (b *Board) TakeNullMove() {

	b.HisPly--
	b.Ply--

	if b.EnPas != NoSq {
		b.hashEP()
	}

	b.CastlePerm = b.History[b.HisPly].CastlePerm
	b.FiftyMove = b.History[b.HisPly].FiftyMove
	b.EnPas = b.History[b.HisPly].EnPas

	if b.EnPas != NoSq {
		b.hashEP()
	}
	b.Side ^= 1
	b.hashSide()

}
