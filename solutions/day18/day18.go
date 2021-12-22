package day18

import (
	"fmt"

	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/input"

)

func solveDay18a() {
	var sum *snailNumber
	for line := range input.ReadLines() {
		sn := Parse(line)
		if sum == nil {
			sum = sn
		} else {
			sum = sum.Add(*sn)
		}
	}
	fmt.Println(sum.Magnitude())
}

func solveDay18b() {
	numbers := []snailNumber{}
	for line := range input.ReadLines() {
		numbers = append(numbers, *Parse(line))
	}
	max_sum := getMaxSum(numbers)
	fmt.Println(max_sum)
}

func getMaxSum(numbers []snailNumber) int {
	max_sum := 0
	for _, num_a := range numbers {
		for _, num_b := range numbers {
			sum := num_a.Add(num_b)
			magnitude := sum.Magnitude()
			if magnitude > max_sum {
				max_sum = magnitude
			}
		}
	}
	return max_sum
}

func init() {
	if err := registry.RegisterSolution("day18a", solveDay18a); err != nil {
		panic(err)
	}
	if err := registry.RegisterSolution("day18b", solveDay18b); err != nil {
		panic(err)
	}
}
