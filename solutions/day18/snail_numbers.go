package day18

import (
	"math"

	"github.com/heindsight/aoc21/utils/stack"
)

type snailNumber struct {
	value    int
	elements *[2]snailNumber
}

func (sn *snailNumber) Add(other snailNumber) *snailNumber {
	number := &snailNumber{elements: &[2]snailNumber{sn.Copy(), other.Copy()}}
	number.Reduce()
	return number
}

func (sn *snailNumber) Copy() snailNumber {
	dup := snailNumber{}
	if sn.isLeaf() {
		dup.value = sn.value
	} else {
		dup.elements = &[2]snailNumber{}
		dup.elements[0] = sn.elements[0].Copy()
		dup.elements[1] = sn.elements[1].Copy()
	}
	return dup
}

func (sn *snailNumber) Magnitude() int {
	if sn.isLeaf() {
		return sn.value
	}
	return 3 * sn.elements[0].Magnitude() + 2 * sn.elements[1].Magnitude()
}

func (sn *snailNumber) Reduce() {
	for {
		exploded := sn.explode()
		if exploded {
			continue
		}
		split := sn.split()
		if !split {
			break
		}
	}
}

type frame struct {
	depth int
	node *snailNumber
}

func (sn *snailNumber)explode() bool {
	s := stack.NewStack(64)
	s.Push(frame{depth: 0, node: sn})

	add_to_right := -1
	var previous *snailNumber

	for s.Length() > 0 {
		popped, _ := s.Pop()
		sn = popped.(frame).node
		depth := popped.(frame).depth

		if sn.isLeaf() {
			if add_to_right != -1 {
				sn.value += add_to_right
				return true
			}
			previous = sn
		} else if add_to_right == -1 && depth == 4 && sn.elements[0].isLeaf() && sn.elements[1].isLeaf() {
			if previous != nil {
				previous.value += sn.elements[0].value
			}
			add_to_right = sn.elements[1].value
			sn.elements = nil
			sn.value = 0
		} else {
			s.Push(frame{depth: depth + 1, node: &sn.elements[1]})
			s.Push(frame{depth: depth + 1, node: &sn.elements[0]})
		}
	}
	return false
}

func (sn *snailNumber) split() bool {
	if !sn.isLeaf() {
		for i := range sn.elements {
			if sn.elements[i].split() {
				return true
			}
		}
	} else if sn.value >= 10 {
		down := math.Floor(float64(sn.value) / 2.0)
		up := math.Ceil(float64(sn.value) / 2.0)
		sn.elements = &[2]snailNumber{{value: int(down)}, {value: int(up)}}
		sn.value = 0
		return true
	}

	return false
}

func (sn *snailNumber) isLeaf() bool {
	return sn.elements == nil
}
