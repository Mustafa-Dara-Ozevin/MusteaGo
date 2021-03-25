package common

import "math/rand"

const Name string = "Mustea 1.0"
const BoardSqrNum = 120
const MaxGameMoves = 2048
const MaxPosMoves = 256
const MFlagEP = 0x40000
const MFlagPS = 0x80000
const MFlagCa = 0x1000000
const MFlagCap = 0x7c000
const MFlagProm = 0xf00000
const StartingFen string = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

var Sq120ToSq64 [BoardSqrNum]int
var Sq64ToSq120 [64]int
var FileBrd [BoardSqrNum]int
var RankBrd [BoardSqrNum]int

var SetMask [64]uint64
var ClearMask [64]uint64

var PieceKeys [13][120]uint64
var SideKey uint64
var CastleKeys [16]uint64

const (
	Empty = iota
	WP
	WN
	WB
	WR
	WQ
	WK
	BP
	BN
	BB
	BR
	BQ
	BK
)

const (
	FileA = iota
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
	FileNone
)
const (
	Rank1 = iota
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
	RankNone
)

const (
	White = iota
	Black
	Both
)

const (
	A1 = iota + 21 //21
	B1
	C1
	D1
	E1
	F1
	G1
	H1
	A2 = iota + 23 //31
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A3 = iota + 25 //41
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A4 = iota + 27 //51
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A5 = iota + 29 //61
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A6 = iota + 31 //71
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A7 = iota + 33 //81
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A8 = iota + 35 //91
	B8
	C8
	D8
	E8
	F8
	G8
	H8
	NoSq
	OffBoard
)

const (
	WKCA = 1
	WQCA = 2
	BKCA = 4
	BQCA = 8
)

type undo struct {
	move       int
	CastlePerm int
	EnPas      int
	FiftyMove  int
	PosKey     uint64
}

type Move struct {
	move  int
	score int
}

type MoveList struct {
	moves [MaxPosMoves]Move
	count int
}

//Util Funcs

func Fr2Sq(f, r int) int {
	return (21 + f) + (r * 10)
}
func Sq64(sq120 int) int {
	return Sq120ToSq64[sq120]
}
func Sq120(sq64 int) int {
	return Sq64ToSq120[sq64]
}
func Pop(b *uint64) int {
	return PopBit(b)
}
func Cnt(b uint64) int {
	return CountBits(b)
}
func ClrBit(bb *uint64, sq int) {
	*bb &= ClearMask[sq]
}
func SetBit(bb *uint64, sq int) {
	*bb |= SetMask[sq]
}
func Rand64() uint64 {
	return uint64(rand.Uint32())<<32 + uint64(rand.Uint32())
}
func IsBq(sq int) bool {
	return PieceBishopQueen[sq]
}
func IsRq(sq int) bool {
	return PieceRookQueen[sq]
}
func IsKn(sq int) bool {
	return PieceKnight[sq]
}
func IsKi(sq int) bool {
	return PieceKing[sq]
}

func FromSq(move int) int {
	return move & 0x7f
}
func ToSq(move int) int {
	return (move >> 7) & 0x7f
}

func Captured(move int) int {
	return (move >> 14) & 0xf
}
func Promoted(move int) int {
	return (move >> 20) & 0xf
}
func IsEnPas(move int) bool {
	return (move & MFlagEP) != 0
}
func IsPaSt(move int) bool {
	return (move & MFlagPS) != 0
}
func IsCa(move int) bool {
	return (move & MFlagCa) != 0
}
func IsCp(move int) bool {
	return (move & MFlagCap) != 0
}
func IsPr(move int) bool {
	return (move & MFlagProm) != 0
}
