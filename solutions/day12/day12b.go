package day12

import (
	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/set"
)

type oneSmallTwiceValidator struct {
	visited set.Set
	twice string
}

func newOneSmallTwiceValidator() *oneSmallTwiceValidator {
	validator := &oneSmallTwiceValidator{}
	validator.visited = set.NewSet()
	validator.twice = ""
	return validator
}

func (v *oneSmallTwiceValidator) Visit(cave string) {
	if !IsBigCave(cave) {
		if v.visited.Contains(cave) {
			v.twice = cave
		}
		v.visited.Add(cave)
	}
}

func (v *oneSmallTwiceValidator) CanVisit(cave string) bool {
	return cave != "start" &&
		(IsBigCave(cave) || !v.visited.Contains(cave) || len(v.twice) == 0)
}

func (v *oneSmallTwiceValidator) Leave(cave string) {
	if v.twice == cave {
		v.twice = ""
	} else {
		v.visited.Delete(cave)
	}
}

func init() {
	day12b := Day12{validator: newOneSmallTwiceValidator()}
	if err := registry.RegisterSolution("day12b", day12b.Solve); err != nil {
		panic(err)
	}
}
