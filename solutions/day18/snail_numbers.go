package day18

import "math"

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
		_, _, exploded := sn.explode(0, nil, -1)
		if exploded {
			continue
		}
		split := sn.split()
		if !split {
			break
		}
	}
}

func (sn *snailNumber) explode(depth int, left *snailNumber, propagate int) (int, *snailNumber, bool) {
	if sn.isLeaf() {
		if propagate != -1 {
			sn.value += propagate
			return -1, nil, true
		}
		return -1, sn, false
	}

	if propagate == -1 && depth == 4 && sn.elements[0].isLeaf() && sn.elements[1].isLeaf() {
		if left != nil {
			left.value += sn.elements[0].value
		}
		propagate := sn.elements[1].value
		sn.elements = nil
		sn.value = 0
		return propagate, nil, false
	}

	propagate, new_left, done := sn.elements[0].explode(depth+1, left, propagate)
	if done {
		return -1, nil, done
	}

	return sn.elements[1].explode(depth+1, new_left, propagate)
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
