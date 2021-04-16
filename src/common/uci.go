package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (b *Board) ParsePosition(position string) {
	tokens := strings.Split(position, " ")
	if strings.Contains(position, "startpos") {
		b.ParseFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
		if strings.Contains(position, "moves") {
			moves := tokens[3:]
			for _, move := range moves {
				b.MakeMove(b.ParseMove(move))
				b.Ply--
			}
		}
	} else {
		b.ParseFen(tokens[1] + " " + tokens[2] + " " + tokens[3] + " " + tokens[4] + " " + tokens[5] + " " + tokens[6])
		if strings.Contains(position, "moves") {
			moves := tokens[8:]
			for _, move := range moves {
				b.MakeMove(b.ParseMove(move))
			}
		}
	}
}

func (b *Board) ParseGo(text string, info *SearchInfo) {
	defer wg.Done()
	info.Depth = 64
	info.MovesToGo = 30
	info.TimeSet = false
	moveTime := -1
	time := -1
	inc := -1
	tokens := strings.Split(text, " ")
	for idx := range tokens {
		if tokens[idx] == "infinite" {
			break
		} else if tokens[idx] == "depth" {
			info.Depth, _ = strconv.Atoi(tokens[idx+1])
		} else if tokens[idx] == "movestogo" {
			info.MovesToGo, _ = strconv.Atoi(tokens[idx+1])
		} else if tokens[idx] == "moveTime" {
			moveTime, _ = strconv.Atoi(tokens[idx+1])
		} else if b.Side == White && tokens[idx] == "wtime" {
			time, _ = strconv.Atoi(tokens[idx+1])
		} else if b.Side == Black && tokens[idx] == "btime" {
			time, _ = strconv.Atoi(tokens[idx+1])
		} else if b.Side == White && tokens[idx] == "winc" {
			inc, _ = strconv.Atoi(tokens[idx+1])
		} else if b.Side == Black && tokens[idx] == "binc" {
			inc, _ = strconv.Atoi(tokens[idx+1])
		}
	}
	if moveTime != -1 {
		time = moveTime
		info.MovesToGo = 1
	}

	info.StartTime = TimeMs()
	if time != -1 {
		info.TimeSet = true
		time /= info.MovesToGo
		time -= 50
		info.StopTime = info.StartTime + time + inc
	}
	b.SearchPosition(info)
}

func UciLoop() {
	fmt.Printf("id name %s\n", Name)
	fmt.Printf("id author Mustafa Dara Ozevin\n")
	fmt.Printf("uciok\n")

	var info SearchInfo
	var board Board
	var scanner = bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		info.Stopped = true
		text := string(scanner.Text())

		if text == "" {
			continue
		}
		if strings.Contains(text, "isready") {
			fmt.Printf("readyok\n")
		} else if text == "ucinewgame" {
			board.ParseFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
		} else if strings.Contains(text, "go") {
			wg.Wait()
			wg.Add(1)
			info.Stopped = false
			go board.ParseGo(text, &info)
		} else if strings.Contains(text, "quit") {
			info.Quit = true
			break
		} else if strings.Contains(text, "uci") {
			fmt.Printf("id name %s\n", Name)
			fmt.Printf("id author Mustafa Dara Ozevin\n")
			fmt.Printf("uciok\n")
		} else if strings.Contains(text, "position") {
			info.Stopped = true
			board.ParsePosition(text)
		} else if strings.Contains(text, "stop") {
		}
		if info.Quit {
			break
		}

	}
}
