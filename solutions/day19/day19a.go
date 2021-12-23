package day19

import (
	"fmt"

	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/set"
)

func solveDay19a() {
	scanners := readScanners()

	aligned, _ := alignScanners(scanners)

	beacons := set.NewSet()
	for _, scanner := range aligned {
		for _, b := range scanner.Beacons {
			beacons.Add(b)
		}
	}

	fmt.Println(beacons.Length())
}

func init() {
	if err := registry.RegisterSolution("day19a", solveDay19a); err != nil {
		panic(err)
	}
}
