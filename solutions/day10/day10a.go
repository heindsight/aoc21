package day10

import (
	"fmt"
	"io"

	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/stack"
)

var (
	closing = map[rune]bool {
		')': true,
		']': true,
		'}': true,
		'>': true,
	}
	matches = map[rune]rune {
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	scores = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
)

func solveDay10a() {
	score := 0
	for {
		symbols := stack.MakeStack(512)

		var line string
		_, err := fmt.Scanln(&line)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		for _, symbol := range line {
			if !closing[symbol] {
				symbols.Push(symbol)
				continue
			}

			v, err := symbols.Peek()
			if err == stack.EmptyStackError {
				break
			} else if err != nil {
				panic(err)
			}
			head := v.(rune)

			match := matches[head]

			if match == symbol {
				symbols.Pop()
			} else {
				score += scores[symbol]
				break
			}
		}
	}
	fmt.Println(score)
}

func init() {
	if err := registry.RegisterSolution("day10a", solveDay10a); err != nil {
		panic(err)
	}
}
