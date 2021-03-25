package common

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Board struct {
	Pieces [BoardSqrNum]int
	Pawns  [3]uint64

	KingSq [2]int

	Side      int
	EnPas     int
	FiftyMove int

	Ply    int
	HisPly int

	CastlePerm int

	PosKey uint64

	PceNum   [13]int
	BigPce   [2]int
	MajPce   [2]int
	MinPce   [2]int
	Material [2]int

	History [MaxGameMoves]undo

	PceList [13][10]int
}

func (b *Board) ParseFen(fen string) {
	//rank := RANK_8;
	//file := FILE_A;
	piece := 0
	//count := 0;
	//i := 0;
	//isq64 := 0;
	//sq120 := 0;

	b.ResetBoard()

	var tokens = strings.Split(fen, " ")
	if len(tokens) <= 3 {
		fmt.Println("parse fen failed", fen)
	}

	var rank = Rank8
	var file = FileA
	for _, ch := range tokens[0] {
		if (file+1)%9 == 0 {
			rank--
			file = FileA
		}
		if unicode.IsDigit(ch) {
			var n, _ = strconv.Atoi(string(ch))
			file += n
		} else if unicode.IsLetter(ch) {
			switch ch {
			case 'p':
				piece = BP
			case 'r':
				piece = BR
			case 'n':
				piece = BN
			case 'b':
				piece = BB
			case 'k':
				piece = BK
			case 'q':
				piece = BQ
			case 'P':
				piece = WP
			case 'R':
				piece = WR
			case 'N':
				piece = WN
			case 'B':
				piece = WB
			case 'K':
				piece = WK
			case 'Q':
				piece = WQ
			}
			sq64 := rank*8 + file
			sq120 := Sq120(sq64)
			if piece != Empty {
				b.Pieces[sq120] = piece
			}
			file++

		}

	}
	if tokens[1] == "w" {
		b.Side = White
	} else {
		b.Side = Black
	}

	var sCastleRights = tokens[2]

	if strings.Contains(sCastleRights, "K") {
		b.CastlePerm |= WKCA
	}
	if strings.Contains(sCastleRights, "Q") {
		b.CastlePerm |= WKCA
	}
	if strings.Contains(sCastleRights, "k") {
		b.CastlePerm |= BKCA
	}
	if strings.Contains(sCastleRights, "q") {
		b.CastlePerm |= BQCA
	}
	enpas := tokens[3]
	if enpas != "-" {
		var file = int(byte(enpas[0]) - 'a')
		var rank, _ = strconv.Atoi(string(enpas[1]))
		b.EnPas = Fr2Sq(file, rank)
	}

	if len(tokens) > 4 {
		b.FiftyMove, _ = strconv.Atoi(tokens[4])
	}

	b.GeneratePosKey()
	b.UpdateListMaterial()
}

func (b *Board) UpdateListMaterial() {

	for index := 0; index < BoardSqrNum; index++ {
		sq := index
		piece := b.Pieces[index]
		if piece != OffBoard && piece != Empty {
			colour := PieceCol[piece]

			if PieceBig[piece] {
				b.BigPce[colour]++
			}
			if PieceMin[piece] {
				b.MinPce[colour]++
			}
			if PieceMaj[piece] {
				b.MajPce[colour]++
			}

			b.Material[colour] += PieceVal[piece]

			b.PceList[piece][b.PceNum[piece]] = sq
			b.PceNum[piece]++

			if piece == WK {
				b.KingSq[White] = sq
			}
			if piece == BK {
				b.KingSq[Black] = sq
			}

			if piece == WP {
				SetBit(&b.Pawns[White], Sq64(sq))
				SetBit(&b.Pawns[Both], Sq64(sq))
			} else if piece == BP {
				SetBit(&b.Pawns[Black], Sq64(sq))
				SetBit(&b.Pawns[Both], Sq64(sq))
			}
		}
	}
}

func (b *Board) ResetBoard() {
	for idx := 0; idx < BoardSqrNum; idx++ {
		b.Pieces[idx] = OffBoard
	}
	for idx := 0; idx < 64; idx++ {
		b.Pieces[Sq120(idx)] = Empty
	}

	for idx := 0; idx < 2; idx++ {
		b.BigPce[idx] = 0
		b.MajPce[idx] = 0
		b.MinPce[idx] = 0
		b.Pawns[idx] = 0
	}
	b.Pawns[2] = 0
	for idx := 0; idx < 13; idx++ {
		b.PceNum[idx] = 0
	}
	b.KingSq[White] = NoSq
	b.KingSq[Black] = NoSq

	b.Side = Both
	b.EnPas = NoSq
	b.FiftyMove = 0

	b.Ply = 0
	b.HisPly = 0

	b.CastlePerm = 0

	b.PosKey = 0

}

func (b *Board) PrintBoard() {
	fmt.Printf("\nGame Board:\n\n")
	for rank := Rank8; rank >= Rank1; rank-- {
		fmt.Printf("%d  ", rank+1)
		for file := FileA; file <= FileH; file++ {
			sq := Fr2Sq(file, rank)
			piece := b.Pieces[sq]
			fmt.Printf("%3c", PceChar[piece])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n   ")
	for file := FileA; file <= FileH; file++ {
		fmt.Printf("%3c", FileChar[file])
	}
	fmt.Printf("\n")
}

func (b *Board) GeneratePosKey() {
	var finalKey uint64
	for sq := 0; sq < BoardSqrNum; sq++ {
		piece := b.Pieces[sq]
		if piece != OffBoard && piece != Empty {
			finalKey ^= PieceKeys[piece][sq]
		}
	}

	if b.Side == White {
		finalKey ^= SideKey
	}
	if b.EnPas != NoSq {
		finalKey ^= PieceKeys[Empty][b.EnPas]
	}

	finalKey ^= CastleKeys[b.CastlePerm]
	b.PosKey = finalKey
}
