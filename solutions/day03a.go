package solutions

import (
	"fmt"
	"io"
	"unicode/utf8"

	"github.com/heindsight/aoc21/registry"
)

type day03aSolution struct {
}

func (soln day03aSolution) Solve() {
	var bit_counts []int
	gamma := 0
	epsilon := 0
	num_lines := 0

	for ; ; num_lines++ {
		var bitstring string
		_, err := fmt.Scan(&bitstring)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		if num_lines == 0 {
			bit_counts = make([]int, utf8.RuneCountInString(bitstring))
		}
		for pos, bit := range bitstring {
			if bit == '1' {
				bit_counts[pos]++
			}
		}
	}

	for pos, count := range bit_counts {
		value := 1 << (len(bit_counts) - pos - 1)
		if count > num_lines/2 {
			gamma += value
		} else {
			epsilon += value
		}
	}

	fmt.Println(gamma * epsilon)
}

func init() {
	if err := registry.RegisterSolution("day03a", day03aSolution{}); err != nil {
		panic(err)
	}
}
