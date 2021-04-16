package common

var FileBBMasks [8]uint64
var RankBBMasks [8]uint64
var BlackPassedMask [64]uint64
var WhitePassedMask [64]uint64
var IsolatedMask [64]uint64

func initEvalMasks() {
	for sq := 0; sq < 8; sq++ {
		FileBBMasks[sq] = 0
		RankBBMasks[sq] = 0
	}
	for r := Rank8; r >= Rank1; r-- {
		for f := FileA; f <= FileH; f++ {
			sq := r*8 + f
			FileBBMasks[f] |= uint64(1) << sq
			RankBBMasks[r] |= uint64(1) << sq
		}
	}

	for sq := 0; sq < 64; sq++ {
		BlackPassedMask[sq] = 0
		WhitePassedMask[sq] = 0
		IsolatedMask[sq] = 0
	}
	for sq := 0; sq < 64; sq++ {
		tsq := sq + 8
		for tsq < 64 {
			WhitePassedMask[sq] |= uint64(1) << tsq
			tsq += 8
		}
		tsq = sq - 8
		for tsq >= 0 {
			BlackPassedMask[sq] |= uint64(1) << tsq
			tsq -= 8
		}
		if FileBrd[Sq120(sq)] > FileA {
			IsolatedMask[sq] |= FileBBMasks[FileBrd[Sq120(sq)]-1]
			tsq = sq + 7
			for tsq < 64 {
				WhitePassedMask[sq] |= uint64(1) << tsq
				tsq += 8
			}
			tsq = sq - 9
			for tsq >= 0 {
				BlackPassedMask[sq] |= uint64(1) << tsq
				tsq -= 8
			}
		}
		if FileBrd[Sq120(sq)] < FileH {
			IsolatedMask[sq] |= FileBBMasks[FileBrd[Sq120(sq)]+1]
			tsq = sq + 9
			for tsq < 64 {
				WhitePassedMask[sq] |= uint64(1) << tsq
				tsq += 8
			}
			tsq = sq - 7
			for tsq >= 0 {
				BlackPassedMask[sq] |= uint64(1) << tsq
				tsq -= 8
			}
		}

	}

}
func initMvvLva() {
	defer wg.Done()

	for attacker := WP; attacker < BK; attacker++ {
		for victim := WP; victim < BK; victim++ {
			MvvLvaScores[victim][attacker] = VictimScore[victim] + 6 - (VictimScore[attacker] / 100)
		}
	}
}
func initHashKeys() {
	defer wg.Done()

	for idx := 0; idx < 13; idx++ {
		for idx2 := 0; idx2 < 120; idx2++ {
			PieceKeys[idx][idx2] = Rand64()
		}
	}
	SideKey = Rand64()
	for idx := range CastleKeys {
		CastleKeys[idx] = Rand64()
	}
}

func initBitmasks() {
	defer wg.Done()

	for idx := range SetMask {
		SetMask[idx] = 0
		ClearMask[idx] = 0
	}
	for idx := range SetMask {
		SetMask[idx] |= 1 << idx
		ClearMask[idx] = ^SetMask[idx]
	}
}

func initFilesRanksBrd() {
	defer wg.Done()

	for index := 0; index < BoardSqrNum; index++ {
		FileBrd[index] = OffBoard
		RankBrd[index] = OffBoard
	}

	for rank := Rank1; rank <= Rank8; rank++ {
		for file := FileA; file <= FileH; file++ {
			sq := Fr2Sq(file, rank)
			FileBrd[sq] = file
			RankBrd[sq] = rank
		}
	}

}

func initSq120To64() {
	defer wg.Done()
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
	wg.Add(5)
	go initSq120To64()

	go initBitmasks()

	go initHashKeys()

	go initFilesRanksBrd()

	go initMvvLva()
	wg.Wait()

	initEvalMasks()
}
