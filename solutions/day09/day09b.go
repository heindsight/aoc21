package day09

import (
	"fmt"
	"sort"

	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/stack"
)

func solveDay09b() {
	basin_sizes := make([]int, 0, 512)
	height_map := readHeightMap()

	for p := range findLowPoints(height_map) {
		basin_size := getBasinSize(height_map, p)
		basin_sizes = append(basin_sizes, basin_size)
	}


	sort.Ints(basin_sizes)
	basin_prod := 1
	for _, size := range basin_sizes[len(basin_sizes) - 3:] {
		basin_prod *= size
	}
	fmt.Println(basin_prod)
}

func getBasinSize(height_map map[Point]int, low_point Point) int {
	size := 0
	basin_stack := stack.NewStack(512)
	basin_stack.Push(low_point)
	seen := map[Point]bool{}

	for {
		v, err := basin_stack.Pop()
		if err == stack.EmptyStackError {
			break
		} else if err != nil {
			panic(err)
		}
		p := v.(Point)
		size += 1
		for q := range neighbours(height_map, p) {
			nb := height_map[q]
			if height_map[p] < nb && nb < 9 && !seen[q] {
				seen[q] = true
				basin_stack.Push(q)
			}
		}
	}
	return size
}

func init() {
	if err := registry.RegisterSolution("day09b", solveDay09b); err != nil {
		panic(err)
	}
}
