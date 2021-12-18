package day15

import (
	"container/heap"
	"fmt"

	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/grid"
	"github.com/heindsight/aoc21/utils/pqueue"
	"github.com/heindsight/aoc21/utils/set"
)

type Day15 struct {
	scale_factor int
}

func (d *Day15) solve() {
	cm := grid.ReadDigitGrid()
	extendMap(cm, d.scale_factor)

	width, height := cm.Dimensions()
	start := grid.Point{X: 0, Y: 0}

	goal := grid.Point{X: width - 1, Y: height - 1}

	risk := findLowestRiskPath(cm, start, goal)

	fmt.Println(risk)
}

func extendMap(tile grid.Grid, scale_factor int) {
	width, height := tile.Dimensions()

	for y_tile := 0; y_tile < scale_factor; y_tile++ {
		for x_tile := 0; x_tile < scale_factor; x_tile++ {
			if x_tile == 0 && y_tile == 0 {
				continue
			}
			var p grid.Point
			for p.Y = 0; p.Y < height; p.Y++ {
				for p.X = 0; p.X < width; p.X++ {
					value, _ := tile.Get(p)
					target := p.Offset(x_tile * width, y_tile * height)
					tgt_value := (value.(int) + x_tile + y_tile - 1) % 9 + 1

					tile.Set(target, tgt_value)
				}
			}
		}
	}
}

func findLowestRiskPath(cm grid.Grid, start grid.Point, goal grid.Point) int {
	visited := set.NewSet()
	queue := make(pqueue.PriorityQueue, 1, 128)

	queue[0] = pqueue.MakeItem(start, 0)
	heap.Init(&queue)

	for len(queue) > 0 {
		node := heap.Pop(&queue).(*pqueue.Item)
		position := node.Value().(grid.Point)
		total_risk := -node.Priority()

		if position == goal {
			return total_risk
		}

		for nb_coord := range cm.Neighbours(position, false) {
			if visited.Contains(nb_coord) {
				continue
			}

			nb_risk, err := cm.Get(nb_coord)
			if err != nil {
				panic(err)
			}

			q_item := pqueue.MakeItem(nb_coord, - (total_risk + nb_risk.(int)))
			heap.Push(&queue, q_item)
			visited.Add(nb_coord)
		}
	}
	return -1
}

func init() {
	day15a := Day15{scale_factor: 1}
	if err := registry.RegisterSolution("day15a", day15a.solve); err != nil {
		panic(err)
	}
	day15b := Day15{scale_factor: 5}
	if err := registry.RegisterSolution("day15b", day15b.solve); err != nil {
		panic(err)
	}
}
