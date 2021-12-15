package day13

import (
	"fmt"

	"github.com/heindsight/aoc21/registry"
)

func solveDay13a() {
	page := readDots()

	fold_instructions := readFolds() 
	first_fold := <- fold_instructions
	doFold(page, first_fold)
	fmt.Println(page.Length())
}

func init() {
	if err := registry.RegisterSolution("day13a", solveDay13a); err != nil {
		panic(err)
	}
}
