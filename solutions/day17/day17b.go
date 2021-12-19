package day17

import (
	"fmt"
	"math"

	"github.com/heindsight/aoc21/registry"
)

func solveDay17b() {
	target := readTargetArea()
	minVx, maxVx := xVelocities(target)
	minVy, maxVy := yVelocities(target)

	num_solns := countCommon(
		Solns(target.MinX, target.MaxX, minVx, maxVx, true),
		Solns(target.MinY, target.MaxY, minVy, maxVy, false),
	)
	fmt.Println(num_solns)
}

type interval struct {
	lower float64
	upper float64
}

func Solns(minP, maxP, minV int, maxV int, unbound_upper bool) []interval  {
	solns := []interval{}
	for vX := minV; vX <= maxV; vX++ {
		min_1, max_1 := quadratic(1.0, -float64(2*vX+1), float64(2*minP))
		lower := math.Max(0.0, min_1)
		upper := max_1

		min_1, max_1 = quadratic(1.0, -float64(2*vX+1), float64(2*maxP))
		if !math.IsNaN(min_1) {
			if lower <= min_1 {
				upper = math.Min(upper, min_1)
			} else if upper >= max_1 {
				lower = math.Max(lower, max_1)
			} else {
				continue
			}
		}

		lower = math.Ceil(lower)
		upper = math.Floor(upper)
		if unbound_upper && upper >= float64(vX + 1) {
			upper = math.Inf(1)
		}

		if lower <= upper {
			solns = append(solns, interval{lower, upper})
		}
	}
	return solns
}


func countCommon(xSolns []interval, ySolns []interval) int {
	count := 0
	for _, x := range xSolns {
		for _, y := range ySolns {
			if !(x.lower > y.upper || y.lower > x.upper) {
				count++
			}
		}
	}
	return count
}

func init() {
	if err := registry.RegisterSolution("day17b", solveDay17b); err != nil {
		panic(err)
	}
}
