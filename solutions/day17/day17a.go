package day17

import (
	"fmt"

	"github.com/heindsight/aoc21/registry"
)

func solveDay17a() {
	target := readTargetArea()
	_, velocityY := yVelocities(target)
	maxHeigt := seriesSum(velocityY)
	fmt.Println(maxHeigt)
}

func init() {
	if err := registry.RegisterSolution("day17a", solveDay17a); err != nil {
		panic(err)
	}
}
