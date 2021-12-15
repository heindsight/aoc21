package day12

import (
	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/set"
)


type smallOnceValidator struct {
	visited set.Set
}

func newSmallOnceValidator() *smallOnceValidator {
	validator := &smallOnceValidator{}
	validator.visited = set.NewSet()
	return validator
}

func (v *smallOnceValidator) Visit(cave string) {
	if !IsBigCave(cave) {
		v.visited.Add(cave)
	}
}

func (v *smallOnceValidator) CanVisit(cave string) bool {
	return !v.visited.Contains(cave)
}

func (v *smallOnceValidator) Leave(cave string) {
	v.visited.Delete(cave)
}

func init() {
	day12a := Day12{validator: newSmallOnceValidator()}
	if err := registry.RegisterSolution("day12a", day12a.Solve); err != nil {
		panic(err)
	}
}
