package day08

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/input"
)

func solveDay08a() {
	easy_count := 0
	uniqueDigits := map[int]int{
		2: 1,
		3: 7,
		4: 4,
		7: 8,
	}

	for entry := range readEntries() {
		for _, digit := range entry.output {
			_, unique := uniqueDigits[utf8.RuneCountInString(digit)]
			if unique {
				easy_count++
			}
		}
	}
	fmt.Println(easy_count)
}

type Entry struct {
	signals []string
	output  []string
}

func readEntries() chan Entry {
	out := make(chan Entry)

	go func() {
		for line := range input.ReadLines() {
			digits := strings.Split(line, " ")
			out <- Entry{signals: digits[:10], output: digits[11:]}
		}
		close(out)
	}()
	return out
}

func init() {
	if err := registry.RegisterSolution("day08a", solveDay08a); err != nil {
		panic(err)
	}
}
