package day19

import (
	"fmt"

	"github.com/heindsight/aoc21/registry"
)


func solveDay19b() {
	scanners := readScanners()

	_, transforms := alignScanners(scanners)

	max := 0

	for i := 0; i < len(transforms); i++ {
		for j := i + 1; j < len(transforms); j++ {
			diff := transforms[i].Translation.Sub(transforms[j].Translation)
			manhattan := diff.Manhattan()

			if manhattan > max {
				max = manhattan
			}
		}
	}

	fmt.Println(max)
}

func init() {
	if err := registry.RegisterSolution("day19b", solveDay19b); err != nil {
		panic(err)
	}
}
