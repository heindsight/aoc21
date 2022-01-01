package day23

import (
	"container/heap"
	"fmt"

	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/grid"
	"github.com/heindsight/aoc21/utils/input"
	"github.com/heindsight/aoc21/utils/pqueue"
	"github.com/heindsight/aoc21/utils/set"
	"github.com/heindsight/aoc21/utils/stack"
)

type Day23 struct {
	reader func() *amphipodBurrow
}

func (d Day23) solve() {
	burrow := d.reader()
	finalPositions := moveToDests(burrow)
	if finalPositions != nil {
		fmt.Println(finalPositions.TotalCost)
	}
}

var (
	moveCosts = map[rune]int{'A': 1, 'B': 10, 'C': 100, 'D': 1000}
	destRooms = map[rune]int{'A': 3, 'B': 5, 'C': 7, 'D': 9}
	noStopping = map[int]bool{3: true, 5: true, 7: true, 9: true}
)

type amphipodBurrow struct {
	Map       map[grid.Point]rune
	Amphipods []grid.Point
	TotalCost int
	Signature string
}

type amphipodMove struct {
	Amphipod int
	To       grid.Point
	Cost     int
}

func moveToDests(burrow *amphipodBurrow) *amphipodBurrow {
	queue := make(pqueue.PriorityQueue, 1, 128)
	queue[0] = pqueue.MakeItem(burrow, 0)
	items := map[string]*pqueue.Item{
		burrow.Signature: queue[0],
	}
	heap.Init(&queue)
	done := set.NewSet()

	for len(queue) > 0 {
		state := heap.Pop(&queue).(*pqueue.Item)

		burrow := state.Value().(*amphipodBurrow)
		done.Add(burrow.Signature)

		if burrow.goalConfiguration() {
			return burrow
		}

		for _, move := range burrow.getMoves() {
			next := burrow.Move(move)
			if done.Contains(next) {
				fmt.Println("We've been here before")
				continue
			}

			q_item, seen := items[next.Signature]

			if seen && next.TotalCost < -q_item.Priority() {
				queue.Update(q_item, next, -next.TotalCost)
			} else if !seen {
				q_item = pqueue.MakeItem(next, -next.TotalCost)
				heap.Push(&queue, q_item)
				items[next.Signature] = q_item
			}
		}
	}
	return nil
}

func (burrow *amphipodBurrow) Move(move amphipodMove) *amphipodBurrow {
	dup := &amphipodBurrow{}
	dup.Map = make(map[grid.Point]rune)
	dup.Amphipods = make([]grid.Point, len(burrow.Amphipods))
	
	sig := []rune(burrow.Signature)

	for pos, chr := range burrow.Map {
		dup.Map[pos] = chr
	}

	for i, pos := range burrow.Amphipods {
		dup.Amphipods[i] = pos
	}

	start := burrow.Amphipods[move.Amphipod]
	amphipodType := burrow.Map[start]

	dup.Map[start] = '.'
	sig[sigPos(start.X, start.Y)] = '.'	

	dup.Map[move.To] = amphipodType
	sig[sigPos(move.To.X, move.To.Y)] = amphipodType

	dup.Amphipods[move.Amphipod] = move.To
	dup.TotalCost = burrow.TotalCost + move.Cost
	dup.Signature = string(sig)

	return dup
}

func (burrow *amphipodBurrow) getMoves() []amphipodMove {
	moves := make([]amphipodMove, 0, 40)

	for amphipod := range burrow.Amphipods {
		for _, move := range burrow.getAmphipodMoves(amphipod) {
			moves = append(moves, move)
		}
	}
	return moves
}

func (burrow *amphipodBurrow) getAmphipodMoves(amphipod int) []amphipodMove {
	start := burrow.Amphipods[amphipod]
	amphipodType := burrow.Map[start]
	if amphipodType == '.' || amphipodType == '#' {
		return nil
	}

	if burrow.goalReached(start, amphipodType) {
		return nil
	}

	targets := make([]amphipodMove, 0, 32)
	s := stack.NewStack(32)
	seen := set.NewSet()

	s.Push(amphipodMove{Amphipod: amphipod, Cost: 0, To: start})

	for s.Length() > 0 {
		top, _ := s.Pop()
		m := top.(amphipodMove)

		for _, q := range m.To.Neighbours(false) {
			if burrow.Map[q] != '.' || seen.Contains(q) {
				continue
			}
			seen.Add(q)

			if m.To.Y == 1 && q.Y == 2 {
				if !burrow.canEnterRoom(q.X, amphipodType) {
					continue
				}
			}

			nextMove := amphipodMove{Amphipod: amphipod, Cost: m.Cost + moveCosts[amphipodType], To: q}
			s.Push(nextMove)

			if burrow.legalMove(start, q) {
				targets = append(targets, nextMove)
			}
		}
	}
	return targets
}

func (burrow *amphipodBurrow) goalConfiguration() bool {
	for _, amphipod := range burrow.Amphipods {
		if !burrow.goalReached(amphipod, burrow.Map[amphipod]) {
			return false
		}
	}
	return true
}

func (burrow *amphipodBurrow) canEnterRoom(x int, amphipodType rune) bool {
	if destRooms[amphipodType] != x {
		// Can't enter a room other than the amphipod's destination room
		return false
	}
	p := grid.Point{X: x}

	for p.Y = 2; burrow.Map[p] != '#'; p.Y++ {
		if burrow.Map[p] != '.' && burrow.Map[p] != amphipodType {
			return false
		}
	}
	return true
}

func (burrow *amphipodBurrow) goalReached(pos grid.Point, amphipodType rune) bool {
	if pos.Y < 2 || destRooms[amphipodType] != pos.X {
		return false
	}


	for pos.Y += 1; burrow.Map[pos] != '#'; pos.Y++ {
		if burrow.Map[pos] != amphipodType {
			return false
		}
	}
	return true
}

func (burrow *amphipodBurrow) legalMove(start, dest grid.Point) bool {
	if burrow.goalReached(dest, burrow.Map[start]) {
		return true
	}
	if start.Y != 1 && dest.Y == 1 && !noStopping[dest.X] {
		return true
	}
	return false
}

func readBurrowA() *amphipodBurrow {
	burrow := &amphipodBurrow{}
	burrow.Map = make(map[grid.Point]rune)
	burrow.Amphipods = make([]grid.Point, 8)
	sig := make([]rune, 19)

	y := 0
	i := 0
	for line := range input.ReadLines() {
		parseLine(burrow, line, y, sig, &i)
		y++
	}
	burrow.Signature = string(sig)
	return burrow
}

func readBurrowB() *amphipodBurrow {
	burrow := &amphipodBurrow{}
	burrow.Map = make(map[grid.Point]rune)
	burrow.Amphipods = make([]grid.Point, 16)
	sig := make([]rune, 27)

	y := 0
	i := 0
	for line := range input.ReadLines() {
		parseLine(burrow, line, y, sig, &i)
		if y == 2 {
			y += 3
		} else {
			y++
		}
	}

	folded := []string{
		"  #D#C#B#A#",
		"  #D#B#A#C#",
	}

	for y, line := range folded {
		parseLine(burrow, line, y + 3, sig, &i)
	}

	burrow.Signature = string(sig)
	return burrow
}

func parseLine(burrow *amphipodBurrow, line string, row int, sig []rune, amphipodIdx *int)  {
	for x, char := range line {
		if char == ' ' {
			continue
		}
		loc := grid.Point{X: x, Y: row}
		burrow.Map[loc] = char
		if char != '#' {
			sig[sigPos(x, row)] = char

			if char != '.' {
				burrow.Amphipods[*amphipodIdx] = loc
				*amphipodIdx++
			}
		}
	}
}

func sigPos(x, y int) int {
	if y == 1 {
		return x - 1
	} 
	return 10 + (y - 2) * 4 + x / 2
}

func init() {
	day23a := Day23{reader: readBurrowA}
	if err := registry.RegisterSolution("day23a", day23a.solve); err != nil {
		panic(err)
	}
	day23b := Day23{reader: readBurrowB}
	if err := registry.RegisterSolution("day23b", day23b.solve); err != nil {
		panic(err)
	}
}
