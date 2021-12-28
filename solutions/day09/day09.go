package day09

import (
	"github.com/heindsight/aoc21/utils/grid"
)

func findLowPoints(height_map grid.Grid) chan grid.Point {
	out := make(chan grid.Point)
	go func() {
		width, height := height_map.Dimensions()
		var p grid.Point

		for p.Y = 0; p.Y < height; p.Y++ {
			for p.X = 0; p.X < width; p.X++ {
				if is_lowpoint(height_map, p) {
					out <- p
				}
			}
		}
		close(out)
	}()
	return out
}

func is_lowpoint(height_map grid.Grid, p grid.Point) bool {
	p_height, err := height_map.Get(p)
	if err != nil {
		panic(err)
	}

	for _, q := range p.Neighbours(false) {
		q_height, err := height_map.Get(q)
		if err != nil {
			continue
		}
		if q_height.(int) <= p_height.(int) {
			return false
		}
	}

	return true
}
