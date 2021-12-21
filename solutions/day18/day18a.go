package day18

import (
	"fmt"

	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/input"
	"github.com/heindsight/aoc21/utils/numeric"

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
			sum_a_b := num_a.Add(num_b)
			sum_b_a := num_b.Add(num_a)
			pair_max := numeric.Max(sum_a_b.Magnitude(), sum_b_a.Magnitude())
			if pair_max > max_sum {
				max_sum = pair_max
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
