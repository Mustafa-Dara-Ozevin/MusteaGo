package common

var PceChar = [...]byte{'.', 'P', 'N', 'B', 'R', 'Q', 'K', 'p', 'n', 'b', 'r', 'q', 'k'}
var SİdeChar = [...]byte{'w', 'b', '-'}
var RankChar = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8'}
var FileChar = [...]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}

var PieceBig = [...]bool{false, false, true, true, true, true, true, false, true, true, true, true, true}
var PieceMaj = [...]bool{false, false, false, false, true, true, true, false, false, false, true, true, true}
var PieceMin = [...]bool{false, false, true, true, false, false, false, false, true, true, false, false, false}
var PieceVal = [...]int{0, 100, 325, 325, 550, 1000, 50000, 100, 325, 325, 550, 1000, 50000}
var PieceCol = [...]int{Both, White, White, White, White, White, White,
	Black, Black, Black, Black, Black, Black}

var PiecePawn = [...]bool{false, true, false, false, false, false, false, true, false, false, false, false, false}
var PieceKnight = [...]bool{false, false, true, false, false, false, false, false, true, false, false, false, false}
var PieceKing = [...]bool{false, false, false, false, false, false, true, false, false, false, false, false, true}
var PieceRookQueen = [...]bool{false, false, false, false, true, true, false, false, false, false, true, true, false}
var PieceBishopQueen = [...]bool{false, false, false, true, false, true, false, false, false, true, false, true, false}
var PieceSlides = [...]bool{false, false, false, true, true, true, false, false, false, true, true, true, false}

var Mirror64 = [...]int{
	56, 57, 58, 59, 60, 61, 62, 63,
	48, 49, 50, 51, 52, 53, 54, 55,
	40, 41, 42, 43, 44, 45, 46, 47,
	32, 33, 34, 35, 36, 37, 38, 39,
	24, 25, 26, 27, 28, 29, 30, 31,
	16, 17, 18, 19, 20, 21, 22, 23,
	8, 9, 10, 11, 12, 13, 14, 15,
	0, 1, 2, 3, 4, 5, 6, 7,
}
