package day13

import (
	"fmt"

	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/grid"
	"github.com/heindsight/aoc21/utils/set"
)


func solveDay13b() {
	page := readDots()

	for fold := range readFolds() {
		doFold(page, fold)
	}

	printPage(page)
}

func printPage(page set.Set) {
	rendered := grid.NewGrid(true)
	for dot := range page.Iter() {
		coord := dot.(grid.Point)
		rendered.Set(coord, '#')
	}

	width, height := rendered.Dimensions()
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			point, _ := rendered.Get(grid.Point{X:x, Y:y})
			switch point.(type) {
			case rune:
				fmt.Printf("%c", point.(rune))
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func init() {
	if err := registry.RegisterSolution("day13b", solveDay13b); err != nil {
		panic(err)
	}
}
