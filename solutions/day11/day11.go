package day11

import (
	"fmt"

	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/grid"
)

const (
	iterations = 100
	flash_theshold = 9
)

func solveDay11a() {
	octopus_map := grid.ReadDigitGrid()
	flashes := 0

	for step := 1; step <= iterations; step++ {
		flashes += simulate_step(octopus_map)	
	}

	fmt.Println(flashes)
}

func solveDay11b() {
	octopus_map := grid.ReadDigitGrid()
	octopus_count := octopus_map.Area()
	var steps int

	for steps = 1; simulate_step(octopus_map) < octopus_count; steps++ {
	}

	fmt.Println(steps)
}

func simulate_step(octopi grid.Grid) int {
	flashes := 0
	width, height := octopi.Dimensions()
	var pos grid.Point

	for pos.Y = 0; pos.Y < height; pos.Y++ {
		for pos.X = 0; pos.X < width; pos.X++ {
			flashes += bump_energy(octopi, pos)
		}
	}
	for pos.Y = 0; pos.Y < height; pos.Y++ {
		for pos.X = 0; pos.X < width; pos.X++ {
			octopus, err := octopi.Get(pos)
			if err != nil {
				panic(err)
			}
			if octopus.(int) > flash_theshold {
				octopi.Set(pos, 0)
			}
		}
	}
	return flashes
}

func bump_energy(octopi grid.Grid, pos grid.Point) int {
	octopus, err := octopi.Get(pos)
	if err != nil {
		panic(err)
	}
	energy := octopus.(int)
	octopi.Set(pos, energy + 1)

	if energy == flash_theshold {
		return flash(octopi, pos)
	}
	return 0
}

func flash(octopi grid.Grid, pos grid.Point) int {
	flashes := 1
	for q := range octopi.Neighbours(pos, true) {
		flashes += bump_energy(octopi, q)
	}
	return flashes
}

func init() {
	if err := registry.RegisterSolution("day11a", solveDay11a); err != nil {
		panic(err)
	}
	if err := registry.RegisterSolution("day11b", solveDay11b); err != nil {
		panic(err)
	}
}
