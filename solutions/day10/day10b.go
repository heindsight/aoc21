package day10

import (
	"fmt"
	"sort"

	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/input"
	"github.com/heindsight/aoc21/utils/stack"
)


func solveDay10b() {
	scores := make([]int, 0, 512)
	for line := range input.ReadLines() {
		symbols, _, err := parse(line)
		if err == Incomplete {
			scores = append(scores, get_auto_complete_score(symbols))
		} else if err == Corrupt {
			continue
		} else if err != nil {
			panic(err)

		}
	}
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}

func get_auto_complete_score(symbols stack.Stack) int {
	score := 0
	scores := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	for {
		v, err := symbols.Pop()
		if err == stack.EmptyStackError {
			break
		} else if err != nil {
			panic(err)
		}
		head := v.(rune)
		match := matches[head]
		score *= 5
		score += scores[match]
	}
	return score
}

func init() {
	if err := registry.RegisterSolution("day10b", solveDay10b); err != nil {
		panic(err)
	}
}
