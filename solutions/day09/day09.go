package day09

import (
	"fmt"
	"io"
)

type Point struct {
	x int
	y int
}

func readHeightMap() map[Point]int {
	heightmap := map[Point]int{}

	for y := 0; ; y++ {
		var line string
		_, err := fmt.Scanln(&line)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		for x, height_rune := range line {
			heightmap[Point{x: x, y: y}] = int(height_rune - '0')
		}
	}
	return heightmap
}

func findLowPoints(height_map map[Point]int) chan Point {
	out := make(chan Point)
	go func() {
		for p := range height_map {
			if is_lowpoint(height_map, p) {
				out <- p
			}
		}
		close(out)
	}()
	return out
}

func is_lowpoint(height_map map[Point]int, p Point) bool {
	for q := range neighbours(height_map, p) {
		if height_map[q] <= height_map[p] {
			return false
		}
	}

	return true
}

func neighbours(height_map map[Point]int, p Point) chan Point {
	out := make(chan Point)
	go func() {
		for nx := p.x - 1; nx <= p.x+1; nx += 2 {
			q := Point{x: nx, y: p.y}
			_, found := height_map[q]
			if found {
				out <- q
			}
		}

		for ny := p.y - 1; ny <= p.y+1; ny += 2 {
			q := Point{x: p.x, y: ny}
			_, found := height_map[q]
			if found {
				out <- q
			}
		}
		close(out)
	}()
	return out
}
