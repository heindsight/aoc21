package day08

import (
	"fmt"
	"unicode/utf8"

	"github.com/heindsight/aoc21/registry"
)

func solveDay08a() {
	easy_count := 0

	for display := range readDisplaySignals() {
		for _, digit := range display.output {
			_, unique := uniqueDigits[utf8.RuneCountInString(digit)]
			if unique {
				easy_count++
			}
		}
	}
	fmt.Println(easy_count)
}

func init() {
	if err := registry.RegisterSolution("day08a", solveDay08a); err != nil {
		panic(err)
	}
}
