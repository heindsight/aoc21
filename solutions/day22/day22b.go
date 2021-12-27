package day22

import (
	"fmt"

	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/input"
	"github.com/heindsight/aoc21/utils/set"
)

func solveDay22b() {
	cuboids := set.NewSet()

	for line := range input.ReadLines() {
		step := parseBootStep(line)
		updateCuboids(step, cuboids)
	}
	numOn := countOn(cuboids)
	fmt.Println(numOn)
}

func updateCuboids(step *BootStep, cuboids set.Set) {
	intersect := []Cuboid{}

	for item := range cuboids.Iter() {
		cu := item.(Cuboid)

		if cu.Intersects(&step.Region) {
			intersect = append(intersect, cu)
		}
	}

	if step.On && len(intersect) == 0 {
		cuboids.Add(step.Region)
		return
	}

	for _, cu := range intersect {
		cuboids.Delete(cu)
	}

	intersect = append(intersect, step.Region)
	intersect = SplitCuboids(intersect)

	for _, cu := range intersect {
		if !step.On && cu.Intersects(&step.Region) {
			continue
		}
		cuboids.Add(cu)
	}
}

func countOn(cuboids set.Set) uint64 {
	var count uint64

	for item := range cuboids.Iter() {
		cu := item.(Cuboid)
		count += cu.Volume()
	}
	return count
}

func init() {
	if err := registry.RegisterSolution("day22b", solveDay22b); err != nil {
		panic(err)
	}
}
