package common

var KnDir = [8]int{-8, -19, -21, -12, 8, 19, 21, 12}
var RkDir = [4]int{-1, -10, 1, 10}
var BiDir = [4]int{-9, -11, 11, 9}
var KiDir = [8]int{-1, -10, 1, 10, -9, -11, 11, 9}

func (b *Board) IsAttacked(sq, side int) bool {
	if side == White {
		if b.Pieces[sq-11] == WP || b.Pieces[sq-9] == WP {
			return true
		}
	} else {
		if b.Pieces[sq+11] == BP || b.Pieces[sq+9] == BP {
			return true
		}
	}

	// knights
	for index := 0; index < 8; index++ {
		pce := b.Pieces[sq+KnDir[index]]
		if pce != OffBoard && IsKn(pce) && PieceCol[pce] == side {
			return true
		}
	}

	// rooks, queens
	for index := 0; index < 4; index++ {
		dir := RkDir[index]
		t_sq := sq + dir
		pce := b.Pieces[t_sq]
		for pce != OffBoard {
			if pce != Empty {
				if IsRq(pce) && PieceCol[pce] == side {
					return true
				}
				break
			}
			t_sq += dir
			pce = b.Pieces[t_sq]
		}
	}

	// bishops, queens
	for index := 0; index < 4; index++ {
		dir := BiDir[index]
		t_sq := sq + dir
		pce := b.Pieces[t_sq]
		for pce != OffBoard {
			if pce != Empty {
				if IsBq(pce) && PieceCol[pce] == side {
					return true
				}
				break
			}
			t_sq += dir
			pce = b.Pieces[t_sq]
		}
	}

	// kings
	for index := 0; index < 8; index++ {
		pce := b.Pieces[sq+KiDir[index]]
		if pce != OffBoard && IsKi(pce) && PieceCol[pce] == side {
			return true
		}
	}

	return false
}
