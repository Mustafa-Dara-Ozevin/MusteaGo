package common

const PawnIsolated = -10

var PawnPassed = [8]int{0, 5, 10, 20, 35, 60, 100, 200}

const RookOpenFile = 10
const RookSemiOpenFile = 5
const QueenOpenFile = 5
const QueenSemiOpenFile = 3
const BishopPair = 30

var EndgameMat = 1*PieceVal[WR] + 2*PieceVal[WN] + 2*PieceVal[WP]

var PawnTable = [64]int{
	0, 0, 0, 0, 0, 0, 0, 0,
	10, 10, 0, -10, -10, 0, 10, 10,
	5, 0, 0, 5, 5, 0, 0, 5,
	0, 0, 10, 20, 20, 10, 0, 0,
	5, 5, 5, 10, 10, 5, 5, 5,
	10, 10, 10, 20, 20, 10, 10, 10,
	20, 20, 20, 30, 30, 20, 20, 20,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var KnightTable = [64]int{
	0, -10, 0, 0, 0, 0, -10, 0,
	0, 0, 0, 5, 5, 0, 0, 0,
	0, 0, 10, 10, 10, 10, 0, 0,
	0, 0, 10, 20, 20, 10, 5, 0,
	5, 10, 15, 20, 20, 15, 10, 5,
	5, 10, 10, 20, 20, 10, 10, 5,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var BishopTable = [64]int{
	0, 0, -10, 0, 0, -10, 0, 0,
	0, 0, 0, 10, 10, 0, 0, 0,
	0, 0, 10, 15, 15, 10, 0, 0,
	0, 10, 15, 20, 20, 15, 10, 0,
	0, 10, 15, 20, 20, 15, 10, 0,
	0, 0, 10, 15, 15, 10, 0, 0,
	0, 0, 0, 10, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var RookTable = [64]int{
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	25, 25, 25, 25, 25, 25, 25, 25,
	0, 0, 5, 10, 10, 5, 0, 0,
}

var KingE = [64]int{
	-50, -10, 0, 0, 0, 0, -10, -50,
	-10, 0, 10, 10, 10, 10, 0, -10,
	0, 10, 20, 20, 20, 20, 10, 0,
	0, 10, 20, 40, 40, 20, 10, 0,
	0, 10, 20, 40, 40, 20, 10, 0,
	0, 10, 20, 20, 20, 20, 10, 0,
	-10, 0, 10, 10, 10, 10, 0, -10,
	-50, -10, 0, 0, 0, 0, -10, -50,
}

var KingO = [64]int{
	0, 5, 5, -10, -10, 0, 10, 5,
	-30, -30, -30, -30, -30, -30, -30, -30,
	-50, -50, -50, -50, -50, -50, -50, -50,
	-70, -70, -70, -70, -70, -70, -70, -70,
	-70, -70, -70, -70, -70, -70, -70, -70,
	-70, -70, -70, -70, -70, -70, -70, -70,
	-70, -70, -70, -70, -70, -70, -70, -70,
	-70, -70, -70, -70, -70, -70, -70, -70,
}
var Mirror = [64]int{
	56, 57, 58, 59, 60, 61, 62, 63,
	48, 49, 50, 51, 52, 53, 54, 55,
	40, 41, 42, 43, 44, 45, 46, 47,
	32, 33, 34, 35, 36, 37, 38, 39,
	24, 25, 26, 27, 28, 29, 30, 31,
	16, 17, 18, 19, 20, 21, 22, 23,
	8, 9, 10, 11, 12, 13, 14, 15,
	0, 1, 2, 3, 4, 5, 6, 7,
}

func mirror64(sq int) int {
	return Mirror[sq]
}

func (b *Board) EvalPosition() int {
	score := b.Material[White] - b.Material[Black]
	var pce int
	var sq int
	pce = WP
	for pceNum := 0; pceNum < b.PceNum[pce]; pceNum++ {
		sq = b.PceList[pce][pceNum]
		score += PawnTable[Sq64(sq)]
		if (IsolatedMask[Sq64(sq)] & b.Pawns[White]) == 0 {
			score += PawnIsolated
		}
		if (WhitePassedMask[Sq64(sq)] & b.Pawns[Black]) == 0 {
			score += PawnPassed[RankBrd[sq]]
		}
	}

	pce = BP
	for pceNum := 0; pceNum < b.PceNum[pce]; pceNum++ {
		sq = b.PceList[pce][pceNum]
		score -= PawnTable[mirror64(Sq64(sq))]
		if (IsolatedMask[Sq64(sq)] & b.Pawns[Black]) == 0 {
			score -= PawnIsolated
		}
		if (BlackPassedMask[Sq64(sq)] & b.Pawns[White]) == 0 {
			score -= PawnPassed[7-RankBrd[sq]]
		}
	}

	pce = WN
	for pceNum := 0; pceNum < b.PceNum[pce]; pceNum++ {
		sq = b.PceList[pce][pceNum]
		score += KnightTable[Sq64(sq)]
	}

	pce = BN
	for pceNum := 0; pceNum < b.PceNum[pce]; pceNum++ {
		sq = b.PceList[pce][pceNum]
		score -= KnightTable[mirror64(Sq64(sq))]
	}

	pce = WB
	for pceNum := 0; pceNum < b.PceNum[pce]; pceNum++ {
		sq = b.PceList[pce][pceNum]
		score += BishopTable[Sq64(sq)]
	}

	pce = BB
	for pceNum := 0; pceNum < b.PceNum[pce]; pceNum++ {
		sq = b.PceList[pce][pceNum]
		score -= BishopTable[mirror64(Sq64(sq))]
	}
	pce = WR
	for pceNum := 0; pceNum < b.PceNum[pce]; pceNum++ {
		sq = b.PceList[pce][pceNum]
		score += RookTable[Sq64(sq)]
		if (b.Pawns[Both] & FileBBMasks[FileBrd[sq]]) == 0 {
			score += RookOpenFile
		} else if (b.Pawns[White] & FileBBMasks[FileBrd[sq]]) == 0 {
			score += RookSemiOpenFile
		}
	}

	pce = BR
	for pceNum := 0; pceNum < b.PceNum[pce]; pceNum++ {
		sq = b.PceList[pce][pceNum]
		score -= RookTable[mirror64(Sq64(sq))]
		if (b.Pawns[Both] & FileBBMasks[FileBrd[sq]]) == 0 {
			score -= RookOpenFile
		} else if (b.Pawns[Black] & FileBBMasks[FileBrd[sq]]) == 0 {
			score -= RookSemiOpenFile
		}
	}

	pce = WQ
	for pceNum := 0; pceNum < b.PceNum[pce]; pceNum++ {
		sq = b.PceList[pce][pceNum]
		if (b.Pawns[Both] & FileBBMasks[FileBrd[sq]]) == 0 {
			score += QueenOpenFile
		} else if (b.Pawns[White] & FileBBMasks[FileBrd[sq]]) == 0 {
			score += QueenSemiOpenFile
		}
	}

	pce = BQ
	for pceNum := 0; pceNum < b.PceNum[pce]; pceNum++ {
		sq = b.PceList[pce][pceNum]
		if (b.Pawns[Both] & FileBBMasks[FileBrd[sq]]) == 0 {
			score -= QueenOpenFile
		} else if (b.Pawns[Black] & FileBBMasks[FileBrd[sq]]) == 0 {
			score -= QueenSemiOpenFile
		}
	}
	sq = b.KingSq[White]
	if b.Material[Black] <= EndgameMat {
		score += KingE[Sq64(sq)]
	} else {
		score += KingO[Sq64(sq)]
	}
	sq = b.KingSq[Black]
	if b.Material[White] <= EndgameMat {
		score -= KingE[mirror64(Sq64(sq))]
	} else {
		score -= KingO[mirror64(Sq64(sq))]
	}
	if b.PceNum[WB] >= 2 {
		score += BishopPair
	}
	if b.PceNum[BB] >= 2 {
		score -= BishopPair
	}

	if b.Side == White {
		return score
	} else {
		return -score
	}
}
