package day04

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/heindsight/aoc21/registry"
)

const boardSize = 5

type bingoResult struct {
	round int
	score int
}

type bingoBoard [boardSize][boardSize]int

func (board *bingoBoard) Play(draws map[int]int) bingoResult {
	win_value := board.calcWinValue(draws)
	unmarked := 0

	for _, row := range board {
		for _, value := range row {
			if draws[value] > draws[win_value] {
				unmarked += value
			}
		}
	}

	return bingoResult{round: draws[win_value], score: unmarked * win_value}
}

func (board *bingoBoard) calcWinValue(draws map[int]int) int {
	type play_info struct {
		hits int
		win_value int
	}

	win_round := math.MaxInt
	win_value := -1

	column_info := make([]play_info, len(board[0]))

	for col_idx := range column_info {
		column_info[col_idx].win_value = -1;
	}

	for _, row := range board {
		row_info := play_info{hits: 0, win_value: -1}

		for col_idx, value := range row {
			round, played := draws[value]
			if !played {
				continue
			}

			row_info.hits++
			if draws[row_info.win_value] < round {
				row_info.win_value = value
			}
			column_info[col_idx].hits++
			if draws[column_info[col_idx].win_value] < round {
				column_info[col_idx].win_value = value
			}
		}
		if row_info.hits == boardSize && draws[row_info.win_value] < win_round {
			win_round = draws[row_info.win_value]
			win_value = row_info.win_value
		}
	}
	for _, col_info := range column_info {
		if col_info.hits == boardSize && draws[col_info.win_value] < win_round {
			win_round = draws[col_info.win_value]
			win_value = col_info.win_value
		}
	}

	return win_value
}

func solveDay04(letTheSquidWin bool) registry.Solution {
	soln := func() {
		draws := readDraws()
		boards := readBoards()
		results := playBoards(boards, draws)

		var outcome bingoResult

		if letTheSquidWin {
			outcome = lastWin(results)
		} else {
			outcome = firstWin(results)
		}

		fmt.Println(outcome.score)
	}
	return soln
}

func readDraws() map[int]int {
	var line string
	_, err := fmt.Scanf("%s\n", &line)
	if err != nil {
		panic(err)
	}

	draws := map[int]int{}

	for pos, number := range strings.Split(line, ",") {
		value, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		draws[value] = pos
	}
	return draws
}

func readBoards() []bingoBoard {
	boards := []bingoBoard{}

	for {
		board := bingoBoard{}

		for row := 0; row < len(board); row++ {
			for col := 0; col < len(board[row]); col++ {
				_, err := fmt.Scan(&board[row][col])
				if err == io.EOF {
					return boards
				} else if err != nil {
					panic(err)
				}
			}
		}
		boards = append(boards, board)
	}
}

func playBoards(boards []bingoBoard, draws map[int]int) []bingoResult {
	results := make([]bingoResult, len(boards))

	for idx, board := range boards {
		results[idx] = board.Play(draws)
	}

	return results
}

func firstWin(results []bingoResult) bingoResult {
	win := bingoResult{round: math.MaxInt, score: 0}

	for _, result := range results {
		if result.round < win.round {
			win = result
		}
	}
	return win
}

func lastWin(results []bingoResult) bingoResult {
	win := bingoResult{round: -1, score: 0}

	for _, result := range results {
		if result.round > win.round {
			win = result
		}
	}
	return win
}

func init() {
	if err := registry.RegisterSolution("day04a", solveDay04(false)); err != nil {
		panic(err)
	}
	if err := registry.RegisterSolution("day04b", solveDay04(true)); err != nil {
		panic(err)
	}
}
