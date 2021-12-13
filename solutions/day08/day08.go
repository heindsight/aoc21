package day08

import (
	"strings"

	"github.com/heindsight/aoc21/utils/input"
)

var uniqueDigits = map[int]int{
	2: 1,
	3: 7,
	4: 4,
	7: 8,
}

type displaySignal struct {
	uniques []string
	output  []string
}

func readDisplaySignals() chan displaySignal {
	out := make(chan displaySignal)

	go func() {
		for line := range input.ReadLines() {
			digits := strings.Split(line, " ")
			out <- displaySignal{uniques: digits[:10], output: digits[11:]}
		}
		close(out)
	}()
	return out
}
