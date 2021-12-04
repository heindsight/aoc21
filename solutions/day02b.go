package solutions

import (
	"fmt"
	"io"

	"github.com/heindsight/aoc21/registry"
)

type day02bSolution struct {
}

func (soln day02bSolution) Solve() {
	aim := 0
	horizontal := 0
	depth := 0

	for {
		var command string
		var arg int
		_, err := fmt.Scanf("%s %d", &command, &arg)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		switch command {
		case "forward":
			horizontal += arg
			depth += aim * arg
		case "down":
			aim += arg
		case "up":
			aim -= arg
		default:
			panic("Unknown command: " + command)
		}
	}

	fmt.Println(horizontal * depth)
}

func init() {
	if err := registry.RegisterSolution("day02b", day02bSolution{}); err != nil {
		panic(err)
	}
}
