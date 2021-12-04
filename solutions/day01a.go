package solutions

import (
	"fmt"
	"io"

	"github.com/heindsight/aoc21/registry"
)

type day01aSolution struct {
}

func (soln day01aSolution) Solve() {
	depth_increases := 0
	first := true
	prev_depth := 0

	for {
		var depth int
		_, err := fmt.Scan(&depth)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		if first {
			first = false
		} else if depth > prev_depth {
			depth_increases++
		}
		prev_depth = depth
	}

	fmt.Println(depth_increases)
}

func init() {
	if err := registry.RegisterSolution("day01a", day01aSolution{}); err != nil {
		panic(err)
	}
}
