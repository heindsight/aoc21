package day02

import (
	"fmt"
	"io"

	"github.com/heindsight/aoc21/registry"
)

func solveDay02a() {
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
		case "down":
			depth += arg
		case "up":
			depth -= arg
		default:
			panic("Unknown command: " + command)
		}
	}

	fmt.Println(horizontal * depth)
}

func init() {
	if err := registry.RegisterSolution("day02a", solveDay02a); err != nil {
		panic(err)
	}
}
