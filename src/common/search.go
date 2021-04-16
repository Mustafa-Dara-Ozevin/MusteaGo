package common

import (
	"fmt"
)

func (s *SearchInfo) CheckUp() bool {
	if s.TimeSet && TimeMs() >= s.StopTime {
		s.Stopped = true
	}
	return false
}

func (m *MoveList) pickNextMove(moveNum int) {
	bestScore := 0
	bestNum := moveNum
	for index := moveNum; index < m.Count; index++ {
		if m.Moves[index].Score > bestScore {
			bestScore = m.Moves[index].Score
			bestNum = index
		}
	}
	temp := m.Moves[moveNum]
	m.Moves[moveNum] = m.Moves[bestNum]
	m.Moves[bestNum] = temp
}

func (b *Board) clearForSearch(info *SearchInfo) {
	for index := 0; index < 13; index++ {
		for index2 := 0; index2 < BoardSqrNum; index2++ {
			b.SearchHistory[index][index2] = 0
		}
	}

	for index := 0; index < 2; index++ {
		for index2 := 0; index2 < MaxDepth; index2++ {
			b.SearchKillers[index][index2] = 0
		}
	}
	b.PvTable.ClearPvTable()
	b.Ply = 0

	info.Nodes = 0
	info.fh = 0
	info.fhf = 0
}

func (b *Board) AlphaBeta(alpha, beta, depth int, info *SearchInfo, doNull bool) int {
	if info.Nodes%2047 == 0 {
		info.CheckUp()
	}
	if depth == 0 {
		info.Nodes++
		return b.Quiescence(alpha, beta, info)
	}

	info.Nodes++
	inCheck := b.IsAttacked(b.KingSq[b.Side], b.Side^1)
	if inCheck {
		depth++
	}

	if b.IsRepetition() || b.FiftyMove >= 100 && b.Ply >= 1 {
		return 0
	}
	if b.Ply > MaxDepth-1 {
		return b.EvalPosition()
	}

	var list MoveList
	b.GenerateAllMoves(&list)
	legal := 0
	oldAlpha := alpha
	bestMove := NoMove
	score := -Infinite
	pvMove := b.ProbePvMove()
	if pvMove != NoMove {
		for moveNum := range list.GetMoves() {
			if list.Moves[moveNum].Move == pvMove {
				list.Moves[moveNum].Score = 2000000
				break
			}
		}
	}
	for moveNum := range list.GetMoves() {
		list.pickNextMove(moveNum)
		if !b.MakeMove(list.Moves[moveNum].Move) {
			continue
		}
		legal++
		score = -b.AlphaBeta(-beta, -alpha, depth-1, info, true)
		b.TakeMove()
		if info.Stopped {
			return 0
		}
		if score > alpha {
			if score >= beta {
				if legal == 1 {
					info.fhf++
				}
				info.fh++

				if !IsCa(list.Moves[moveNum].Move) {
					b.SearchKillers[1][b.Ply] = b.SearchKillers[0][b.Ply]
					b.SearchKillers[0][b.Ply] = list.Moves[moveNum].Move
				}
				return beta
			}
			if bestMove != NoMove && !IsCa(list.Moves[moveNum].Move) {
				b.SearchHistory[b.Pieces[FromSq(bestMove)]][b.Pieces[ToSq(bestMove)]] += depth
			}
			alpha = score
			bestMove = list.Moves[moveNum].Move
		}
	}

	if legal == 0 {
		if inCheck {
			return -Mate + b.Ply
		} else {
			return 0
		}
	}
	if alpha != oldAlpha {
		b.StorePvMove(bestMove)
	}

	return alpha
}

func (b *Board) Quiescence(alpha, beta int, info *SearchInfo) int {

	if info.Nodes%2047 == 0 {
		info.CheckUp()
	}
	info.Nodes++

	if b.IsRepetition() || b.FiftyMove >= 100 && b.Ply >= 1 {
		return 0
	}
	if b.Ply > MaxDepth-1 {
		return b.EvalPosition()
	}
	score := b.EvalPosition()

	if score >= beta {
		return beta
	}
	if score >= alpha {
		alpha = score
	}

	var list MoveList
	b.GenerateAllCaps(&list)
	legal := 0
	oldAlpha := alpha
	bestMove := NoMove
	score = -Infinite

	for moveNum := range list.GetMoves() {
		list.pickNextMove(moveNum)
		if !b.MakeMove(list.Moves[moveNum].Move) {
			continue
		}
		legal++
		score = -b.Quiescence(-beta, -alpha, info)
		b.TakeMove()
		if info.Stopped {
			return 0
		}
		if score > alpha {
			if score >= beta {
				if legal == 1 {
					info.fhf++
				}
				info.fh++

				return beta
			}
			alpha = score
			bestMove = list.Moves[moveNum].Move
		}
	}

	if alpha != oldAlpha {
		b.StorePvMove(bestMove)
	}

	return alpha
}

func (b *Board) SearchPosition(info *SearchInfo) {
	bestMove := NoMove
	bestScore := -Infinite
	currentDepth := 0

	b.clearForSearch(info)
	for currentDepth = 1; currentDepth <= info.Depth; currentDepth++ {
		bestScore = b.AlphaBeta(-Infinite, Infinite, currentDepth, info, false)
		if info.Stopped {
			break
		}
		pvMoves := b.GetPvLine(currentDepth)
		bestMove = b.PvArray[0]
		fmt.Printf("%d %d %d %1d",
			currentDepth, bestScore, TimeMs()-info.StartTime, info.Nodes)
		for pvNum := 0; pvNum < pvMoves; pvNum++ {
			fmt.Printf(" %s", PrMove(b.PvArray[pvNum]))
		}
		fmt.Printf("\n")
		//fmt.Printf("Ordering:%.2f\n", info.fhf/info.fh)
		if bestScore > Mate-MaxDepth || bestScore < -Mate+MaxDepth {
			break
		}
	}
	fmt.Printf("bestmove %s\n", PrMove(bestMove))

}
