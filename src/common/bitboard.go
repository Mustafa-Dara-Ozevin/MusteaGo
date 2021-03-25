package common

import "fmt"

var BitTable = [...]int{
	63, 30, 3, 32, 25, 41, 22, 33, 15, 50, 42, 13, 11, 53, 19, 34, 61, 29, 2,
	51, 21, 43, 45, 10, 18, 47, 1, 54, 9, 57, 0, 35, 62, 31, 40, 4, 49, 5, 52,
	26, 60, 6, 23, 44, 46, 27, 56, 16, 7, 39, 48, 24, 59, 14, 12, 55, 38, 28,
	58, 20, 37, 17, 36, 8}

func PopBit(bb *uint64) int {
	var b uint64 = *bb ^ (*bb - 1)
	var fold uint32 = uint32(((b & 0xffffffff) ^ (b >> 32)))
	*bb &= (*bb - 1)
	return BitTable[(fold*0x783a9b23)>>26]
}

func CountBits(b uint64) int {
	var r int
	for r = 0; b != 0; r++ {
		b &= b - 1
	}
	return r
}

func PrintBitBoard(bb uint64) {
	var shiftme uint64 = 1
	fmt.Printf("\n")
	for rank := Rank8; rank >= Rank1; rank-- {
		for file := FileA; file <= FileH; file++ {
			sq := Fr2Sq(file, rank)
			sq64 := Sq64(sq)
			if (shiftme<<sq64)&bb != 0 {
				fmt.Printf("X")
			} else {
				fmt.Printf("-")
			}
		}
		fmt.Printf("\n")
	}

}
