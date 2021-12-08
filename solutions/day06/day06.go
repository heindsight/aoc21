package day06

import (
	"fmt"

	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/input"
)

const (
	spawnAfter = 6
	initialSpawnDelay = 8
)

type Day06 struct {
	days int
}

func (d *Day06) solve() {
	fish_counts := readInitialState()

	for day := 0; day < d.days; day++ {
		spawn := fish_counts[0]
		new_counts := make(map[int]int, len(fish_counts))
		for days_to_spawn, count := range fish_counts {
			if days_to_spawn == 0 {
				continue
			}
			new_counts[days_to_spawn - 1] = count
		}
		new_counts[spawnAfter] += spawn
		new_counts[initialSpawnDelay] = spawn
		fish_counts = new_counts
	}
	total_fish := 0
	for _, count := range fish_counts {
		total_fish += count
	}
	fmt.Println(total_fish)
}

func readInitialState() map[int]int {
	fish := map[int]int{}

	for item := range input.ReadCommaSepLineInts() {
		fish[item.Value]++
	}
	return fish
}

func init() {
	day06a := Day06{days: 80}
	if err := registry.RegisterSolution("day06a", day06a.solve); err != nil {
		panic(err)
	}
	day06b := Day06{days: 256}
	if err := registry.RegisterSolution("day06b", day06b.solve); err != nil {
		panic(err)
	}
}
