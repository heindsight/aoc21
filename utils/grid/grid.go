package grid

import (
	"errors"
	"sync"

	"github.com/heindsight/aoc21/utils/input"
	"github.com/heindsight/aoc21/utils/numeric"
)

type Point struct {
	X int
	Y int
}

func (p Point) Offset(x, y int) Point {
	return Point{X: p.X + x, Y: p.Y + y}
}

var OutOfBoundsError = errors.New("Point out of grid bounds")

type Grid interface {
	Get(Point) (interface{}, error)
	Set(Point, interface{}) error
	Dimensions() (int, int)
	Area() int
	Neighbours(Point, bool) chan Point
}

type grid struct {
	cells map[Point]interface{}
	x_max int
	y_max int
	lock sync.Mutex
}

func NewGrid() *grid {
	g := &grid{}
	g.cells = make(map[Point]interface{})
	return g
}

func (g *grid) Get(pos Point) (interface{}, error) {
	g.lock.Lock()
	defer g.lock.Unlock()
	if pos.X < 0 || pos.X > g.x_max || pos.Y < 0 || pos.Y > g.y_max {
		return nil, OutOfBoundsError
	}
	return g.cells[pos], nil
}

func (g *grid) Set(pos Point, value interface{}) error {
	g.lock.Lock()
	defer g.lock.Unlock()

	if pos.X < 0 || pos.Y < 0 {
		return OutOfBoundsError
	}

	g.x_max = numeric.Max(g.x_max, pos.X)
	g.y_max = numeric.Max(g.y_max, pos.Y)

	g.cells[pos] = value
	return nil
}

func (g *grid) Dimensions() (int, int) {
	g.lock.Lock()
	defer g.lock.Unlock()

	return g.x_max + 1, g.y_max + 1
}

func (g *grid) Area() int {
	width, height := g.Dimensions()
	return width * height
}

func (g *grid) Neighbours(p Point, include_diagonal bool) chan Point {
	out := make(chan Point, 8)
	go func() {
		var q Point
		width, height := g.Dimensions()

		for q.Y = numeric.Max(0, p.Y - 1); q.Y < numeric.Min(p.Y + 2, height); q.Y++ {
			for q.X = numeric.Max(0, p.X - 1); q.X < numeric.Min(p.X + 2, width); q.X++ {
				if (q == p) || (!include_diagonal && q.X != p.X && q.Y != p.Y) {
					continue
				}

				out <- q
			}
		}
		close(out)
	}()
	return out
}

func ReadDigitGrid() *grid {
	grid := NewGrid()
	y := 0

	for line := range input.ReadLines() {
		for x, char := range line {
			grid.Set(Point{X: x, Y: y}, int(char - '0'))
		}
		y++
	}
	return grid
}
