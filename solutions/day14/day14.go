package day14

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"

	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/input"
)

type Day14 struct {
	steps int
}

func (d Day14) solve() {
	inStream := input.ReadLines()

	polymer := <-inStream
	transformer := polymerTransformer{rules: readRules(inStream)}
	counts := transformer.Transform(polymer, d.steps)

	min, max := counts.getMinMax()
	fmt.Println(max - min)
}

type cacheKey struct {
	pair  string
	steps int
}

type polymerTransformer struct {
	rules polymerTransformRules
	cache map[cacheKey]runeCounts
}

func (xfm *polymerTransformer) Transform(template string, steps int) runeCounts {
	counts := make(runeCounts)

	xfm.cache = make(map[cacheKey]runeCounts)

	var prev rune

	for _, element := range template {
		counts[element]++
		if prev != 0 {
			pair := string([]rune{prev, element})
			pair_counts := xfm.transformPair(pair, steps)
			counts.combine(&pair_counts)
		}
		prev = element
	}
	return counts
}

func (xfm *polymerTransformer) transformPair(pair string, steps int) runeCounts {
	cached := xfm.getFromCache(pair, steps)
	if cached != nil {
		return cached
	}

	counts := make(runeCounts)

	insert, found := xfm.rules[pair]
	if steps > 0 && found {
		counts[insert]++

		for replace := range xfm.rules.Transform(pair) {
			pair_counts := xfm.transformPair(replace, steps-1)
			counts.combine(&pair_counts)
		}
	}
	xfm.addToCache(pair, steps, counts)
	return counts
}

func (xfm *polymerTransformer) getFromCache(pair string, steps int) runeCounts {
	key := cacheKey{pair, steps}
	return xfm.cache[key]
}

func (xfm *polymerTransformer) addToCache(pair string, steps int, counts runeCounts) {
	key := cacheKey{pair, steps}
	xfm.cache[key] = counts
}

type runeCounts map[rune]int

func (c *runeCounts) combine(other *runeCounts) {
	for chr, count := range *other {
		(*c)[chr] += count
	}
}

func (c *runeCounts) getMinMax() (int, int) {
	min := math.MaxInt
	max := 0

	for _, count := range *c {
		if count > max {
			max = count
		}
		if count < min {
			min = count
		}
	}
	return min, max
}

type polymerTransformRules map[string]rune

func readRules(stream chan string) polymerTransformRules {
	rules := make(polymerTransformRules)

	// Discard one line
	_ = <-stream

	for line := range stream {
		rule := strings.Fields(line)
		rules[rule[0]], _ = utf8.DecodeLastRuneInString(rule[2])
	}
	return rules
}

func (p *polymerTransformRules) Transform(pair string) chan string {
	out := make(chan string, 2)

	go func() {
		insert, found := (*p)[pair]
		if found {
			first, _ := utf8.DecodeRuneInString(pair)
			last, _ := utf8.DecodeLastRuneInString(pair)
			out <- string([]rune{first, insert})
			out <- string([]rune{insert, last})
		}
		close(out)
	}()

	return out
}

func init() {
	day14a := Day14{steps: 10}
	if err := registry.RegisterSolution("day14a", day14a.solve); err != nil {
		panic(err)
	}
	day14b := Day14{steps: 40}
	if err := registry.RegisterSolution("day14b", day14b.solve); err != nil {
		panic(err)
	}
}
