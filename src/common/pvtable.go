package common

func (b *Board) GetPvLine(depth int) int {
	move := b.ProbePvMove()
	count := 0
	for move != NoMove && count < depth {
		if b.MoveExist(move) {
			b.MakeMove(move)
			b.PvArray[count] = move
			count++
		} else {
			break
		}
		move = b.ProbePvMove()
	}
	for b.Ply > 0 {
		b.TakeMove()
	}
	return count
}

func (p *PvTable) ClearPvTable() {
	p.NumEntries = len(p.PTable)
	for i := range p.PTable {
		p.PTable[i] = PvEntry{PosKey: 0, Move: NoMove}
	}
}

func (b *Board) StorePvMove(move int) {
	index := b.PosKey % uint64(b.PvTable.NumEntries)
	b.PvTable.PTable[index].Move = move
	b.PvTable.PTable[index].PosKey = b.PosKey

}
func (b *Board) ProbePvMove() int {
	index := b.PosKey % uint64(b.PvTable.NumEntries)
	key := uint64(b.PvTable.PTable[index].PosKey)
	if key == b.PosKey {
		return b.PvTable.PTable[index].Move
	}
	return NoMove
}
