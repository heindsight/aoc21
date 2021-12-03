package solutions

import (
	"errors"
	"fmt"
	"io"

	"github.com/heindsight/aoc21/registry"
)

type day02aSolution struct {
}

func (soln day02aSolution) Solve() error {
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
		case "down":
			depth += arg
		case "up":
			depth -= arg
		default:
			return errors.New("Unknown command: " + command)
		}
	}

	fmt.Println(horizontal * depth)
	return nil
}

func init() {
	if err := registry.RegisterSolution("day02a", day02aSolution{}); err != nil {
		fmt.Println("Failed to register day02a solution", err)
	}
}
