package day21

import (
	"fmt"

	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/input"
	"github.com/heindsight/aoc21/utils/numeric"
)


func solveDay21a() {
	reader := input.ReadLines()
	player1 := readStartPosition(reader)
	player2 := readStartPosition(reader)

	rolls, loser_score := playDeterministicDice(player1, player2)
	fmt.Println(rolls * loser_score)
}

type DeterministicD100 struct {
	val int
	rolls int
}

func (d *DeterministicD100) Roll() int {
	v := d.val
	d.rolls++

	d.val = d.val % 100 + 1
	return v
}

func playDeterministicDice(player1, player2 int) (int, int) {
	positions := [2]int{player1, player2}
	scores := [2]int{}
	die := DeterministicD100{val: 1}

	for player := 0; numeric.Max(scores[0], scores[1]) < 1000; player = (player + 1) % 2 {
		rolls := [3]int{die.Roll(), die.Roll(), die.Roll()}
		roll := rolls[0] + rolls[1] + rolls[2]
		positions[player] = (positions[player] + roll - 1) % 10 + 1
		scores[player] += positions[player]
	}
	return die.rolls, numeric.Min(scores[0], scores[1])
}

func init() {
	if err := registry.RegisterSolution("day21a", solveDay21a); err != nil {
		panic(err)
	}
}
