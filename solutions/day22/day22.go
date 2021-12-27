package day22

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Interval struct {
	Upper int
	Lower int
}

func (iv *Interval) Intersects(other *Interval) bool {
	return iv.Lower < other.Upper && other.Lower < iv.Upper
}

func (iv *Interval) Length() uint64 {
	return uint64(iv.Upper - iv.Lower)
}

func SplitIntervals(ivals []Interval) []Interval {
	endpoints := unique_endpoints(ivals)
	result := make([]Interval, len(endpoints)-1)

	for i := 0; i < len(endpoints)-1; i++ {
		result[i].Lower = endpoints[i]
		result[i].Upper = endpoints[i+1]
	}
	return result
}

func unique_endpoints(ivals []Interval) []int {
	endpoints := make([]int, 2*len(ivals))
	j := 0
	for _, iv := range ivals {
		endpoints[j] = iv.Lower
		endpoints[j+1] = iv.Upper
		j += 2
	}

	sort.Ints(endpoints)
	result := make([]int, 0, len(endpoints))

	for _, v := range endpoints {
		if len(result) > 0 && result[len(result)-1] == v {
			continue
		}
		result = append(result, v)
	}
	return result
}

func (iv Interval) GoString() string {
	return fmt.Sprintf("[%d, %d]", iv.Lower, iv.Upper)
}

type Cuboid struct {
	X, Y, Z Interval
}

func (cu *Cuboid) Intersects(other *Cuboid) bool {
	return cu.X.Intersects(&other.X) &&
		cu.Y.Intersects(&other.Y) &&
		cu.Z.Intersects(&other.Z)
}

func (cu *Cuboid) Volume() uint64 {
	return cu.X.Length() * cu.Y.Length() * cu.Z.Length()
}

func  SplitCuboids(cb []Cuboid) []Cuboid {
	X := make([]Interval, len(cb))
	Y := make([]Interval, len(cb))
	Z := make([]Interval, len(cb))

	for i, cu := range cb {
		X[i] = cu.X
		Y[i] = cu.Y
		Z[i] = cu.Z
	}
	X = SplitIntervals(X)
	Y = SplitIntervals(Y)
	Z = SplitIntervals(Z)

	result := make([]Cuboid, 0, len(X)*len(Y)*len(Z))

	for _, x := range X {
		for _, y := range Y {
			for _, z := range Z {
				b := Cuboid{X: x, Y: y, Z: z}
				for _, cu := range cb {
					if b.Intersects(&cu) {
						result = append(result, b)
						break
					}
				}
			}
		}
	}
	return result
}

func (cu Cuboid) GoString() string {
	return fmt.Sprintf("{X: %#v, Y: %#v, Z: %#v}", cu.X, cu.Y, cu.Z)
}

type BootStep struct {
	On bool
	Region Cuboid
}

func parseBootStep(instruction string) *BootStep {
	bs := &BootStep{}

	fields := strings.Fields(instruction)
	bs.On = fields[0] == "on"

	for _, interval := range strings.Split(fields[1], ",") {
		axis, ival := parseInterval(interval)
		switch axis {
		case 'x':
			bs.Region.X = ival
		case 'y':
			bs.Region.Y = ival
		case 'z':
			bs.Region.Z = ival
		}
	}
	return bs
}

func parseInterval(desc string) (rune, Interval) {
	parts := strings.Split(desc, "=")
	axis, _ := utf8.DecodeRuneInString(parts[0])
	endpoints := strings.Split(parts[1], "..")
	ival := Interval{}
	ival.Lower, _ = strconv.Atoi(endpoints[0])
	ival.Upper, _ = strconv.Atoi(endpoints[1])
	ival.Upper++
	return axis, ival
}
