package solutions

import (
	"fmt"
	"io"

	"github.com/heindsight/aoc21/registry"
)

type day01aSolution struct {
	depth_increases int
}

func (soln day01aSolution) Solve() error {
	first := true
	prev_depth := 0

	for {
		var depth int
		_, err := fmt.Scan(&depth)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if first {
			first = false
		} else if depth > prev_depth {
			soln.depth_increases++
		}
		prev_depth = depth
	}

	fmt.Println(soln.depth_increases)
	return nil
}

func init() {
	if err := registry.RegisterSolution("day01a", day01aSolution{}); err != nil {
		fmt.Println("Failed to register day01a solution", err)
	}
}
