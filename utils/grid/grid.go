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
	BoundingBox() (Point, Point)
	Region(Point) chan Point
	Neighbours(Point, bool) chan Point
	Iter() chan Point
}

type grid struct {
	cells map[Point]interface{}
	x_min int
	y_min int
	x_max int
	y_max int
	bounded bool
	lock sync.Mutex
}

func NewGrid(bounded bool) *grid {
	g := &grid{}
	g.cells = make(map[Point]interface{})
	g.bounded = bounded
	return g
}

func (g *grid) Get(pos Point) (interface{}, error) {
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.bounded && (pos.X < 0 || pos.X > g.x_max || pos.Y < 0 || pos.Y > g.y_max) {
		return nil, OutOfBoundsError
	}
	return g.cells[pos], nil
}

func (g *grid) Set(pos Point, value interface{}) error {
	g.lock.Lock()
	defer g.lock.Unlock()

	if g.bounded && (pos.X < 0 || pos.Y < 0) {
		return OutOfBoundsError
	}

	g.x_min = numeric.Min(g.x_min, pos.X)
	g.y_min = numeric.Min(g.y_min, pos.Y)
	g.x_max = numeric.Max(g.x_max, pos.X)
	g.y_max = numeric.Max(g.y_max, pos.Y)

	g.cells[pos] = value
	return nil
}

func (g *grid) Dimensions() (int, int) {
	g.lock.Lock()
	defer g.lock.Unlock()

	return g.x_max - g.x_min + 1, g.y_max - g.y_min + 1
}

func (g *grid) Area() int {
	width, height := g.Dimensions()
	return width * height
}

func (g *grid) BoundingBox() (Point, Point) {
	topLeft := Point{X: g.x_min, Y: g.y_min}
	botRight := Point{X: g.x_max, Y: g.y_max}
	return topLeft, botRight
}


func (g *grid) Region(p Point) chan Point {
	out := make(chan Point, 9)
	go func() {
		var (
			q Point
			x_min, y_min, x_max, y_max int
		)

		if g.bounded {
			width, height := g.Dimensions()
			x_min = numeric.Max(0, p.X - 1)
			y_min = numeric.Max(0, p.Y - 1)
			x_max = numeric.Min(p.X + 2, width)
			y_max = numeric.Min(p.Y + 2, height)
		} else {
			x_min = p.X - 1
			y_min = p.Y - 1
			x_max = p.X + 2
			y_max = p.Y + 2
		}

		for q.Y = y_min; q.Y < y_max; q.Y++ {
			for q.X = x_min; q.X < x_max; q.X++ {
				out <- q
			}
		}
		close(out)
	}()
	return out
}

func (g *grid) Neighbours(p Point, include_diagonal bool) chan Point {
	out := make(chan Point, 8)
	go func() {
		for q := range g.Region(p) {
			if (q == p) || (!include_diagonal && q.X != p.X && q.Y != p.Y) {
				continue
			}
			out <- q
		}
		close(out)
	}()
	return out
}

func (g *grid) Iter() chan Point {
	out := make(chan Point, len(g.cells))
	go func() {
		g.lock.Lock()
		defer g.lock.Unlock()
		for p := range g.cells {
			out <- p
		}
		close(out)
	}()
	return out
}

func ReadDigitGrid() *grid {
	grid := NewGrid(true)
	y := 0

	for line := range input.ReadLines() {
		for x, char := range line {
			grid.Set(Point{X: x, Y: y}, int(char - '0'))
		}
		y++
	}
	return grid
}
