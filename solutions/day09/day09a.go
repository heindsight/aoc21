package day09

import (
	"fmt"

	"github.com/heindsight/aoc21/registry"
)

func solveDay09a() {
	total_risk := 0
	height_map := readHeightMap()

	for p := range findLowPoints(height_map) {
		fmt.Printf("Low point %v has height %d\n\n", p, height_map[p])
		total_risk += height_map[p] + 1
	}

	fmt.Println(total_risk)
}

func init() {
	if err := registry.RegisterSolution("day09a", solveDay09a); err != nil {
		panic(err)
	}
}
