package day12

import (
	"fmt"
	"strings"

	"github.com/heindsight/aoc21/utils/input"
)

type pathValidator interface {
	Visit(string)
	CanVisit(string) bool
	Leave(string)
}

type Day12 struct {
	validator pathValidator
}

func (d *Day12) Solve() {
	caveMap := readCaveMap()
	paths_found := countPaths(caveMap, "start", "end", d.validator)
	fmt.Println(paths_found)
}

type graph struct {
	adjacencies map[string][]string
}

func newGraph() *graph {
	g := &graph{}
	g.adjacencies = make(map[string][]string)
	return g
}

func (g *graph) AddEdge(node_a string, node_b string) {
	g.adjacencies[node_a] = append(g.adjacencies[node_a], node_b)
	g.adjacencies[node_b] = append(g.adjacencies[node_b], node_a)
}

func (g *graph) Neighbours(node string) []string {
	return g.adjacencies[node]
}

func readCaveMap() *graph {
	caveMap := newGraph()

	for line := range input.ReadLines() {
		nodes := strings.Split(line, "-")
		caveMap.AddEdge(nodes[0], nodes[1])
	}
	return caveMap
}

func countPaths(caveMap *graph, start string, end string, validator pathValidator) int {
	pathCount := 0

	validator.Visit(start)

	for _, neighbour := range caveMap.Neighbours(start) {
		if neighbour == end {
			pathCount += 1
			continue
		}
		if validator.CanVisit(neighbour) {
			pathCount += countPaths(caveMap, neighbour, end, validator)
		}
	}
	validator.Leave(start)
	return pathCount
}

func IsBigCave(str string) bool {
	return str == strings.ToUpper(str)
}
