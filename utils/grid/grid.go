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

func (p *Point) Region() []Point {
	out := make([]Point, 9)
	var q Point

	i := 0

	for q.Y = p.Y - 1; q.Y <= p.Y+1; q.Y++ {
		for q.X = p.X - 1; q.X <= p.X+1; q.X++ {
			out[i] = q
			i++
		}
	}
	return out
}

func (p *Point) Neighbours(include_diagonal bool) []Point {
	var out []Point

	if include_diagonal {
		out = make([]Point, 8)
	} else {
		out = make([]Point, 4)
	}

	i := 0
	for _, q := range p.Region() {
		if (q == *p) || (!include_diagonal && q.X != p.X && q.Y != p.Y) {
			continue
		}
		out[i] = q
		i++
	}
	return out
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
	Iter() chan Point
}

type grid struct {
	cells   map[Point]interface{}
	x_min   int
	y_min   int
	x_max   int
	y_max   int
	bounded bool
	lock    sync.Mutex
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
			grid.Set(Point{X: x, Y: y}, int(char-'0'))
		}
		y++
	}
	return grid
}
