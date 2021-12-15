package day10

import (
	"errors"

	"github.com/heindsight/aoc21/utils/stack"
)

var (
	Incomplete = errors.New("Incomplete")
	Corrupt = errors.New("Corrupt")

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
)

func parse(line string) (stack.Stack, *rune, error) {
	symbols := stack.NewStack(512)
	for _, symbol := range line {
		if !closing[symbol] {
			symbols.Push(symbol)
			continue
		}

		v, err := symbols.Peek()
		if err == stack.EmptyStackError {
			return symbols, &symbol, Corrupt
		} else if err != nil {
			panic(err)
		}
		head := v.(rune)

		match := matches[head]

		if match == symbol {
			symbols.Pop()
		} else {
			return symbols, &symbol, Corrupt
		}
	}
	return symbols, nil, Incomplete
}
