package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/heindsight/aoc21/registry"
	_ "github.com/heindsight/aoc21/solutions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "<problem>")
		os.Exit(1)
	}

	problem_name := strings.ToLower(os.Args[1])

	soln, err := registry.GetSolution(problem_name)
	if err != nil {
		panic(err)
	}

	soln.Solve()
}
