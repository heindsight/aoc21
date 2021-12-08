package day07

import (
	"fmt"
	"math"

	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/input"
	"github.com/heindsight/aoc21/utils/numeric"
)

type Day07 struct {
	fuel_cost func(int) int
}

func (d *Day07) solve() {
	crab_positions := readPositions()

	min_pos, max_pos := numeric.MinMax(crab_positions)
	min_fuel := math.MaxInt

	for target_pos := min_pos; target_pos <= max_pos; target_pos++ {
		fuel := 0
		for _, pos := range crab_positions {
			fuel += d.fuel_cost(numeric.Abs(target_pos - pos))
		}
		if fuel < min_fuel {
			min_fuel = fuel
		}
	}

	fmt.Println(min_fuel)
}

func linear_cost(distance int) int {
	return distance
}

func quadratic_cost(distance int) int {
	return distance * (distance + 1) / 2
}

func readPositions() []int {
	positions := []int{}

	for item := range input.ReadCommaSepLineInts() {
		positions = append(positions, item.Value)
	}
	return positions
}

func init() {
	day07a := Day07{fuel_cost: linear_cost}
	if err := registry.RegisterSolution("day07a", day07a.solve); err != nil {
		panic(err)
	}
	day07b := Day07{fuel_cost: quadratic_cost}
	if err := registry.RegisterSolution("day07b", day07b.solve); err != nil {
		panic(err)
	}
}
