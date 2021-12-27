package day21

import (
	"fmt"

	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/input"
)


func solveDay21b() {
	reader := input.ReadLines()
	player1 := readStartPosition(reader)
	player2 := readStartPosition(reader)

	winning_universes := playMultiverseDice(player1, player2)
	fmt.Println(winning_universes)
}

func playMultiverseDice(player1, player2 int) int64 {
	m := make(multiverse)
	initial_universe := universe{
		positions: [2]int{player1, player2},
	}
	wins_player1, wins_player2 := m.Play(initial_universe)
	if wins_player1 > wins_player2 {
		return wins_player1
	} else {
		return wins_player2
	}
}


type universe struct {
	positions [2]int
	scores [2]int
	turn int
}

func (u *universe) Roll() []universe {
	splits := make([]universe, 27)

	for i := 0; i < 27; i++ {
		splits[i] = *u
		d := 3 + i%3 + (i/3)%3 + (i/9)%3
		player := u.turn

		splits[i].turn = (player + 1) % 2
		splits[i].positions[player] = (splits[i].positions[player] + d - 1) % 10 + 1
		splits[i].scores[player] += splits[i].positions[player]
	}
	return splits
}

func (u *universe) Winner() (int, bool) {
	for player := 0; player < 2; player++ {
		if u.scores[player] >= 21 {
			return player, true
		}
	}
	return -1, false
}

type multiverse map[universe][2]int64

func (m *multiverse) Play(u universe) (int64, int64) {
	wins, found := (*m)[u]
	if !found {
		for _, child := range u.Roll() {
			winner, has_winner := child.Winner()
			if has_winner {
				wins[winner] += 1
			} else {
				w0, w1 := m.Play(child)
				wins[0] += w0
				wins[1] += w1
			}
		}
		(*m)[u] = wins
	}

	return wins[0], wins[1]
}

func init() {
	if err := registry.RegisterSolution("day21b", solveDay21b); err != nil {
		panic(err)
	}
}
