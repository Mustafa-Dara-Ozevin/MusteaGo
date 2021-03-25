package common

func initHashKeys() {
	for idx := 0; idx > 13; idx++ {
		for idx2 := 0; idx2 > 120; idx2++ {
			PieceKeys[idx][idx2] = Rand64()
		}
	}
	SideKey = Rand64()
	for idx := range CastleKeys {
		CastleKeys[idx] = Rand64()
	}
}

func initBitmasks() {
	for idx := range SetMask {
		SetMask[idx] = 0
		ClearMask[idx] = 0
	}
	for idx := range SetMask {
		SetMask[idx] |= 1 << idx
		ClearMask[idx] = ^SetMask[idx]
	}
}

func initSq120To64() {

	for index := 0; index < BoardSqrNum; index++ {
		Sq120ToSq64[index] = 65
	}
	for index := 0; index < 64; index++ {
		Sq64ToSq120[index] = 120
	}
	sq64 := 0
	for rank := Rank1; rank <= Rank8; rank++ {
		for file := FileA; file <= FileH; file++ {
			sq := Fr2Sq(file, rank)
			Sq64ToSq120[sq64] = sq
			Sq120ToSq64[sq] = sq64
			sq64++
		}
	}

}

func AllInit() {
	initSq120To64()
	initBitmasks()
	initHashKeys()
}
