package day19

import (
	"math"
	"strconv"
	"strings"

	"github.com/heindsight/aoc21/utils/set"
	"github.com/heindsight/aoc21/utils/input"
	"github.com/heindsight/aoc21/utils/numeric"
)

type Vector [3]int

func (p Vector) Sub(q Vector) Vector {
	d := Vector{}
	for i := range p {
		d[i] = p[i] - q[i]
	}
	return d
}

func (p Vector) Add(q Vector) Vector {
	d := Vector{}
	for i := range p {
		d[i] = p[i] + q[i]
	}
	return d
}

func (p Vector) Length() float64 {
	s := 0.0
	for _, x := range p {
		s += math.Pow(float64(x), 2)
	}

	return math.Sqrt(s)
}

func (p Vector) Manhattan() int {
	s := 0
	for _, x := range p {
		s += numeric.Abs(x)
	}

	return s
}

func (p Vector) Dot(q Vector) int {
	v := 0
	for i := range p {
		v += p[i] * q[i]
	}
	return v
}

type Matrix [3]Vector

func (m *Matrix) Mul(p Vector) Vector {
	q := Vector{}

	for i, row := range m {
		q[i] = row.Dot(p)
	}
	return q
}

var Rotations = []Matrix{
	{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
	{{1, 0, 0}, {0, 0, 1}, {0, -1, 0}},
	{{1, 0, 0}, {0, 0, -1}, {0, 1, 0}},
	{{1, 0, 0}, {0, -1, 0}, {0, 0, -1}},
	{{0, 1, 0}, {1, 0, 0}, {0, 0, -1}},
	{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}},
	{{0, 1, 0}, {0, 0, -1}, {-1, 0, 0}},
	{{0, 1, 0}, {-1, 0, 0}, {0, 0, 1}},
	{{0, 0, 1}, {1, 0, 0}, {0, 1, 0}},
	{{0, 0, 1}, {0, 1, 0}, {-1, 0, 0}},
	{{0, 0, 1}, {0, -1, 0}, {1, 0, 0}},
	{{0, 0, 1}, {-1, 0, 0}, {0, -1, 0}},
	{{0, 0, -1}, {1, 0, 0}, {0, -1, 0}},
	{{0, 0, -1}, {0, 1, 0}, {1, 0, 0}},
	{{0, 0, -1}, {0, -1, 0}, {-1, 0, 0}},
	{{0, 0, -1}, {-1, 0, 0}, {0, 1, 0}},
	{{0, -1, 0}, {1, 0, 0}, {0, 0, 1}},
	{{0, -1, 0}, {0, 0, 1}, {-1, 0, 0}},
	{{0, -1, 0}, {0, 0, -1}, {1, 0, 0}},
	{{0, -1, 0}, {-1, 0, 0}, {0, 0, -1}},
	{{-1, 0, 0}, {0, 1, 0}, {0, 0, -1}},
	{{-1, 0, 0}, {0, 0, 1}, {0, 1, 0}},
	{{-1, 0, 0}, {0, 0, -1}, {0, -1, 0}},
	{{-1, 0, 0}, {0, -1, 0}, {0, 0, 1}},
}
var Identity = Rotations[0]

type Scanner struct {
	Beacons   []Vector
	Distances map[Vector]set.Set
}

func makeScanner() Scanner {
	s := Scanner{}
	s.Distances = make(map[Vector]set.Set)
	return s
}

func (s *Scanner) AddBeacon(p Vector) {
	s.Distances[p] = set.NewSet()

	for _, q := range s.Beacons {
		d := p.Sub(q).Length()
		s.Distances[p].Add(d)
		s.Distances[q].Add(d)
	}

	s.Beacons = append(s.Beacons, p)
}

func (s *Scanner) MatchBeacons(other *Scanner) map[Vector]Vector {
	matches := make(map[Vector]Vector)

	for p, dists_p := range s.Distances {
		for q, dists_q := range other.Distances {
			common := dists_p.Intersection(dists_q)
			if common.Length() >= 11 {
				matches[p] = q
			}
		}
	}
	return matches
}

func (s *Scanner) Align(transform Transformation) Scanner {
	transformed := Scanner{}
	transformed.Beacons = make([]Vector, len(s.Beacons))
	transformed.Distances = make(map[Vector]set.Set, len(s.Distances))

	for i, b := range s.Beacons {
		q := transform.Transform(b)
		transformed.Beacons[i] = q
		transformed.Distances[q] = s.Distances[b].Copy()
	}
	return transformed
}

type Transformation struct {
	Rotation    Matrix
	Translation Vector
}

func (t *Transformation) Transform(p Vector) Vector {
	return t.Rotation.Mul(p).Add(t.Translation)
}

func findXForm(matches map[Vector]Vector) (Transformation, bool) {
	var xform Transformation

	for _, xform.Rotation = range Rotations {
		first := true
		found := false
		for p, q := range matches {
			if first {
				rotated := xform.Rotation.Mul(p)
				xform.Translation = q.Sub(rotated)
				first = false
				found = true
			} else if xform.Transform(p) != q {
				found = false
				break
			}
		}
		if found {
			return xform, true
		}
	}
	return Transformation{}, false
}

func findAlignment(scanner *Scanner, aligned []Scanner) (Transformation, bool) {
	for j := len(aligned) - 1; j >= 0; j-- {
		matches := scanner.MatchBeacons(&aligned[j])
		if len(matches) >= 12 {
			transform, found_transform := findXForm(matches)
			if found_transform {
				return transform, true
			}
		}
	}
	return Transformation{}, false
}

func alignScanners(scanners []Scanner) ([]Scanner, []Transformation) {
	aligned := make([]Scanner, 1, len(scanners))
	transformations := make([]Transformation, 1, len(scanners))

	aligned[0] = scanners[0]
	transformations[0] = Transformation{Rotation: Identity}

	found_alignment := set.NewSet()
	found_alignment.Add(0)

	for len(aligned) < len(scanners) {
		start_size := len(aligned)
		for i := 1; i < len(scanners); i++ {
			if found_alignment.Contains(i) {
				continue
			}
			transform, found := findAlignment(&scanners[i], aligned)
			if found {
				aligned = append(aligned, scanners[i].Align(transform))
				transformations = append(transformations, transform)
				found_alignment.Add(i)
			}
		}
		if len(aligned) == start_size {
			panic("No new alignments found! Would loop forever")
		}
	}
	return aligned, transformations
}

func readScanners() []Scanner {
	scanners := []Scanner{}

	var scanner *Scanner

	for line := range input.ReadLines() {
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "---") {
			scanners = append(scanners, makeScanner())
			scanner = &scanners[len(scanners)-1]
			continue
		}

		beacon := Vector{}
		for i, val := range strings.Split(line, ",") {
			v, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			beacon[i] = v
		}
		scanner.AddBeacon(beacon)
	}
	return scanners
}
