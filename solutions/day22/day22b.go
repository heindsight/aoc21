package day22

import (
	"fmt"

	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/input"
)

func solveDay22b() {
	reactor := []BootStep{}

	for line := range input.ReadLines() {
		step := parseBootStep(line)
		updateCuboids(step, &reactor)
	}
	numOn := countOn(reactor)
	fmt.Println(numOn)
}

func updateCuboids(step BootStep, reactor *[]BootStep) {
	for _, bs := range *reactor {
		intersection, intersects := bs.Region.Intersection(&step.Region)
		if intersects {
			*reactor = append(*reactor, BootStep{Region: intersection, On: !bs.On})
		}
	}

	if step.On {
		*reactor = append(*reactor, step)
	}
}

func countOn(cuboids []BootStep) uint64 {
	var count uint64

	for _, item := range cuboids {
		if item.On {
			count += item.Region.Volume()
		} else {
			count -= item.Region.Volume()
		}
	}
	return count
}

func init() {
	if err := registry.RegisterSolution("day22b", solveDay22b); err != nil {
		panic(err)
	}
}
