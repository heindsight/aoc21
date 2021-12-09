package day05

import (
	"fmt"
	"io"

	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/numeric"
)

type Point struct {
	x int
	y int
}

func (p *Point) offset(x, y int) Point {
	return Point{p.x + x, p.y + y}
}

type Segment struct {
	points [2]Point
}

func (s *Segment) is_horiz() bool {
	return s.points[0].x == s.points[1].x
}

func (s *Segment) is_vert() bool {
	return s.points[0].y == s.points[1].y
}

func (s *Segment) walk() chan Point {
	out := make(chan Point)
	_walk := func() {
		dx := s.points[1].x - s.points[0].x
		dy := s.points[1].y - s.points[0].y

		length := numeric.Max(numeric.Abs(dx), numeric.Abs(dy))

		for i := 0; i <= length; i++ {
			out <- s.points[0].offset(
				int(i*dx/length), int(i*dy/length),
			)
		}
		close(out)
	}
	go _walk()
	return out
}

type Day05 struct {
	include_diagonal bool
}

func (d *Day05) solve() {
	points := map[Point]int{}
	vent_map := []Segment{}
	intersections := 0

	for {
		s, err := readSegment()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		if d.include_diagonal || s.is_horiz() || s.is_vert() {
			vent_map = append(vent_map, s)
		}
	}

	for _, segment := range vent_map {
		for p := range segment.walk() {
			points[p]++
			if points[p] == 2 {
				intersections++
			}
		}
	}
	fmt.Println(intersections)
}

func readSegment() (Segment, error) {
	var s Segment
	_, err := fmt.Scanf(
		"%d,%d -> %d,%d",
		&s.points[0].x,
		&s.points[0].y,
		&s.points[1].x,
		&s.points[1].y,
	)

	return s, err
}

func init() {
	day05a := Day05{include_diagonal: false};
	if err := registry.RegisterSolution("day05a", day05a.solve); err != nil {
		panic(err)
	}
	day05b := Day05{include_diagonal: true};
	if err := registry.RegisterSolution("day05b", day05b.solve); err != nil {
		panic(err)
	}
}
