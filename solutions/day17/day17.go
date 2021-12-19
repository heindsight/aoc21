package day17

import (
	"fmt"
	"math"

	"github.com/heindsight/aoc21/utils/numeric"
)

type Area struct {
	MinX int
	MaxX int
	MinY int
	MaxY int
}

func yVelocities(target Area) (int, int) {
	return target.MinY, numeric.Abs(target.MinY) - 1
}

func xVelocities(target Area) (int, int) {
	_, Vmin := quadratic(1, 1, -2.0*float64(target.MinX))
	return int(math.Ceil(Vmin)), target.MaxX
}

func quadratic(a float64, b float64, c float64) (float64, float64) {
	d := b*b - 4.0*a*c
	soln_0 := (-b - math.Sqrt(d)) / (2.0 * a)
	soln_1 := (-b + math.Sqrt(d)) / (2.0 * a)
	return math.Min(soln_0, soln_1), math.Max(soln_0, soln_1)
}

func seriesSum(n int) int {
	return n * (n + 1) / 2
}

func readTargetArea() Area {
	var targetArea Area

	_, err := fmt.Scanf(
		"target area: x=%d..%d, y=%d..%d\n",
		&targetArea.MinX, &targetArea.MaxX,
		&targetArea.MinY, &targetArea.MaxY)
	if err != nil {
		panic(err)
	}

	return targetArea
}
