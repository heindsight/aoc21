package day10

import (
	"fmt"

	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/input"
)

func solveDay10a() {
	scores := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	score := 0
	for line := range input.ReadLines() {
		_, symbol, err := parse(line)
		if err == Incomplete {
			continue
		} else if err == Corrupt {
			score += scores[*symbol]
		} else if err != nil {
			panic(err)

		}
	}
	fmt.Println(score)
}

func init() {
	if err := registry.RegisterSolution("day10a", solveDay10a); err != nil {
		panic(err)
	}
}
