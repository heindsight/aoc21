package solutions

import (
	"errors"
	"fmt"
	"io"

	"github.com/heindsight/aoc21/registry"
)

type day02bSolution struct {
}

func (soln day02bSolution) Solve() error {
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
			return err
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
			return errors.New("Unknown command: " + command)
		}
	}

	fmt.Println(horizontal * depth)
	return nil
}

func init() {
	if err := registry.RegisterSolution("day02b", day02bSolution{}); err != nil {
		fmt.Println("Failed to register day02b solution", err)
	}
}
