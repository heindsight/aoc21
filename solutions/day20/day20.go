package day20

import (
	"fmt"
	"unicode/utf8"

	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/grid"
	"github.com/heindsight/aoc21/utils/input"
)

type Day20 struct {
	iterations int
}

func (d Day20) solve() {
	inLines := input.ReadLines()
	rules := readRules(inLines)
	image := readImage(inLines)

	for i := 0; i < d.iterations; i++ {
		image = image.Process(rules)
	}
	lightCount := image.CountLight()
	fmt.Println(lightCount)
}

type Image struct {
	pixels  grid.Grid
	missing int
}

func (img *Image) Process(rules []int) Image {
	enhanced := Image{pixels: grid.NewGrid(false)}

	ruleLookups := make(map[grid.Point]uint)
	defaultLookups := map[int]uint {
		1: 0x1ff,
		0: 0,
	}

	for p := range img.pixels.Iter() {
		val := img.Get(p)
		i := 0
		for _, q := range p.Region() {
			lookup, found := ruleLookups[q]
			if !found {
				lookup = defaultLookups[img.missing]
			}
			if val == 1 {
				lookup |= uint(1) << i
			} else {
				lookup &= ^(uint(1)<<i)
			}
			ruleLookups[q] = lookup
			i++
		}
	}


	if img.missing == 0 {
		enhanced.missing = rules[0]
	} else {
		enhanced.missing = rules[511]
	}

	for p, lookup := range ruleLookups {
		val := rules[lookup]
		if val != enhanced.missing {
			enhanced.Set(p, val)
		}
	}

	return enhanced
}

func (img *Image) Set(p grid.Point, val int) {
	img.pixels.Set(p, val)
}

func (img *Image) Get(p grid.Point) int {
	v, _ := img.pixels.Get(p)
	if v != nil {
		return v.(int)
	} else {
		return img.missing
	}
}

func (img *Image) CountLight() int {
	count := 0
	topLeft, botRight := img.pixels.BoundingBox()

	for x := topLeft.X - 1; x < botRight.X+2; x++ {
		for y := topLeft.Y - 1; y < botRight.Y+2; y++ {
			count += img.Get(grid.Point{X: x, Y: y})
		}
	}
	return count
}

func readRules(inLines chan string) []int {
	ruleMap := map[rune]int {
		'#': 1,
		'.': 0,
	}

	line := <-inLines
	rules := make([]int, utf8.RuneCountInString(line))
	for i, c := range line {
		rules[i] = ruleMap[c]
	}
	return rules
}

func readImage(inLines chan string) Image {
	<-inLines
	img := Image{pixels: grid.NewGrid(false), missing: 0}
	row := 0

	for line := range inLines {
		for col, pixel := range line {
			if pixel == '#' {
				img.Set(grid.Point{X: col, Y: row}, 1)
			}
		}
		row++
	}
	return img
}

func init() {
	day20a := Day20{iterations: 2}
	if err := registry.RegisterSolution("day20a", day20a.solve); err != nil {
		panic(err)
	}
	day20b := Day20{iterations: 50}
	if err := registry.RegisterSolution("day20b", day20b.solve); err != nil {
		panic(err)
	}
}
