package day09

import (
	"fmt"
	"sort"

	"github.com/heindsight/aoc21/registry"
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

type Stack struct {
	stack []Point
}

func MakeStack() *Stack {
	return &Stack{stack: make([]Point, 0, 512)}
}

func (s *Stack) Push(p Point) {
	s.stack = append(s.stack, p)
}

func (s *Stack) Pop() Point {
	p := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack) - 1]
	return p
}

func (s *Stack) Length() int {
	return len(s.stack)
}

func getBasinSize(height_map map[Point]int, low_point Point) int {
	size := 0
	stack := MakeStack()
	stack.Push(low_point)
	seen := map[Point]bool{}

	for ; stack.Length() > 0; {
		p := stack.Pop()
		size += 1
		for q := range neighbours(height_map, p) {
			nb := height_map[q]
			if height_map[p] < nb && nb < 9 && !seen[q] {
				seen[q] = true
				stack.Push(q)
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
