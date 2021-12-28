package day09

import (
	"fmt"
	"sort"

	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/stack"
	"github.com/heindsight/aoc21/utils/grid"
)

func solveDay09b() {
	basin_sizes := make([]int, 0, 512)
	height_map := grid.ReadDigitGrid()

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

func getBasinSize(height_map grid.Grid, low_point grid.Point) int {
	size := 0
	basin_stack := stack.NewStack(512)
	basin_stack.Push(low_point)
	seen := map[grid.Point]bool{}

	for {
		v, err := basin_stack.Pop()
		if err == stack.EmptyStackError {
			break
		}
		p := v.(grid.Point)
		height, _ := height_map.Get(p)
		size += 1
		for _, q := range p.Neighbours(false) {
			nb, err := height_map.Get(q)
			if err != nil {
				continue
			}
			if height.(int) < nb.(int) && nb.(int) < 9 && !seen[q] {
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
