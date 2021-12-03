package solutions

import (
	"errors"
	"fmt"
	"io"

	"github.com/heindsight/aoc21/registry"
)

type day02bSolution struct {
	aim        int
	horizontal int
	depth      int
}

func (soln day02bSolution) Solve() error {
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
			soln.horizontal += arg
			soln.depth += soln.aim * arg
		case "down":
			soln.aim += arg
		case "up":
			soln.aim -= arg
		default:
			return errors.New("Unknown command: " + command)
		}
	}

	fmt.Println(soln.horizontal * soln.depth)
	return nil
}

func init() {
	if err := registry.RegisterSolution("day02b", day02bSolution{}); err != nil {
		fmt.Println("Failed to register day02b solution", err)
	}
}
