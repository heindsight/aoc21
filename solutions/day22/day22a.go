package day22

import (
	"fmt"

	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/input"
	"github.com/heindsight/aoc21/utils/numeric"
	"github.com/heindsight/aoc21/utils/set"
)

func solveDay22a() {
	reactor := set.NewSet()
	for line := range input.ReadLines() {
		step := parseBootStep(line)
		executeBootStep(step, reactor)
	}
	fmt.Println(reactor.Length())
}

type cube struct {
	X int
	Y int
	Z int
}

func executeBootStep(step BootStep, reactor set.Set) {
	x_min := numeric.Max(step.Region.X.Lower, -50)
	x_max := numeric.Min(step.Region.X.Upper, 51)

	y_min := numeric.Max(step.Region.Y.Lower, -50)
	y_max := numeric.Min(step.Region.Y.Upper, 51)

	z_min := numeric.Max(step.Region.Z.Lower, -50)
	z_max := numeric.Min(step.Region.Z.Upper, 51)

	for x := x_min; x < x_max; x++ {
		for y := y_min; y < y_max; y++ {
			for z := z_min; z < z_max; z++ {
				if step.On {
					reactor.Add(cube{X: x, Y: y, Z: z})
				} else {
					reactor.Delete(cube{X: x, Y: y, Z: z})
				}
			}
		}
	}
}

func init() {
	if err := registry.RegisterSolution("day22a", solveDay22a); err != nil {
		panic(err)
	}
}
