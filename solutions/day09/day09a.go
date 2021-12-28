package day09

import (
	"fmt"

	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/grid"
)

func solveDay09a() {
	total_risk := 0
	height_map := grid.ReadDigitGrid()

	for p := range findLowPoints(height_map) {
		height, _ := height_map.Get(p)
		total_risk += height.(int) + 1
	}

	fmt.Println(total_risk)
}

func init() {
	if err := registry.RegisterSolution("day09a", solveDay09a); err != nil {
		panic(err)
	}
}
